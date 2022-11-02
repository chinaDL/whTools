package terminal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/chinaDL/whTools/utils"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/shirou/gopsutil/v3/process"
	"io"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	DefaultTerminalConfig = &TerminalConfig{
		BufferSize: 1024 * 4,
	}
	terminalManagerOnce sync.Once
	terminalManager     *TerminalManagerStruct
)

type TerminalManagerStruct struct {
	Terminals map[string]ITerminal
}

func (b *TerminalManagerStruct) GetAll() map[string]ITerminal {
	return b.Terminals
}

func (b *TerminalManagerStruct) Get(id string) ITerminal {
	if v, ok := b.Terminals[id]; ok {
		return v
	} else {
		return nil
	}
}

func (b *TerminalManagerStruct) CloseAll() bool {
	deleteKeys := make([]string, 0)
	for k, v := range b.Terminals {
		err := v.Close()
		if err == nil {
			deleteKeys = append(deleteKeys, k)
		}
	}
	for _, v := range deleteKeys {
		delete(b.Terminals, v)
	}
	return len(b.Terminals) == 0
}
func (b *TerminalManagerStruct) Exec(id string, command string) error {
	terminal := b.Get(id)
	if terminal == nil {
		return errors.New("终端ID不存在")
	}
	return terminal.Exec(command)
}

const (
	SplitCommandStr = "InBugW_H_A_Z"
	CommandEndStr   = "========InBugCommandRunEnd========"
)

type ITerminal interface {
	Exec(command string, isWaitInput ...bool) error
	Terminate() error
	Close() error
	GetID() string
	SetID(id string)
}

type TerminalConfig struct {
	BufferSize int32
}

type Terminal struct {
	ID           string
	PID          int
	inPipe       io.WriteCloser
	outPipe      io.ReadCloser
	errPipe      io.ReadCloser
	out          *bytes.Buffer
	Shell        *exec.Cmd
	buf          []byte
	TerminateFun func() error
	isInitOK     bool
	IsRun        bool
	isClose      bool
	commandId    int32
	wg           sync.WaitGroup
	//ReadBufferCh    chan []byte
}

func (b *Terminal) GetID() string {
	return b.ID
}

func (b *Terminal) SetID(id string) {
	delete(GetTerminalManager().Terminals, b.ID)
	b.ID = id
	GetTerminalManager().Terminals[b.ID] = b
}

func (b *Terminal) Exec(command string, isWaitInput ...bool) error {
	pd, err := process.PidExists(int32(b.PID))
	if b.isClose || !pd {
		return errors.New("控制台已经关闭")
	}
	isWait := false
	if len(isWaitInput) > 0 {
		isWait = isWaitInput[0]
	}
	if b.IsRun && isWait {
		//b.wg.Wait()
	}
	b.IsRun = true
	defer func() {
		b.IsRun = false
	}()
	commandStr := fmt.Sprintf("%s\n", command)
	if isWait {
		b.wg.Add(1)
		commandStr += "echo " + CommandEndStr + "\n"
	}
	inPipe := bufio.NewWriter(b.inPipe)
	_, err = inPipe.WriteString(commandStr)
	if err != nil {
		return err
	}
	err = inPipe.Flush()
	if err != nil {
		return err
	}
	if isWait {
		b.wg.Wait()
	}
	return nil
}

func (b *Terminal) Terminate() error {
	return b.TerminateFun()
}

func (b *Terminal) Close() error {
	b.isClose = true
	err := b.Terminate()
	err = b.inPipe.Close()
	err = b.outPipe.Close()
	err = b.errPipe.Close()
	err = b.Shell.Process.Kill()
	return err
}

func newTerminal(cmd *Terminal, handleMsg func(id string, msgBytes []byte), terminalConfig *TerminalConfig, terminalId ...string) {
	cmd.ID = guid.S()
	if len(terminalId) > 0 {
		cmd.ID = terminalId[0]
	}
	var command string
	var changeCodePageCommand string
	if utils.IsWindows() {
		command = "cmd"
		changeCodePageCommand = "chcp 65001"
	} else {
		command = "/bin/bash"
		changeCodePageCommand = "export LANG=en_US.UTF-8"
	}
	cmd.Shell = exec.Command(command)

	cmd.inPipe, _ = cmd.Shell.StdinPipe()
	cmd.outPipe, _ = cmd.Shell.StdoutPipe()
	cmd.errPipe, _ = cmd.Shell.StderrPipe()

	cmd.isInitOK = false
	cmd.isClose = false
	cmd.Shell.SysProcAttr = &syscall.SysProcAttr{
		//HideWindow: true,
	}
	cmd.TerminateFun = func() error {
		ps, err := GetProcessChildren(cmd.PID)
		if err == nil {
			for _, v := range ps {
				err := v.Terminate()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		return err
	}
	cmd.buf = make([]byte, terminalConfig.BufferSize)

	_ = cmd.Shell.Start()
	cmd.PID = cmd.Shell.Process.Pid
	_ = cmd.Exec(changeCodePageCommand, false)
	_ = cmd.Exec("echo "+SplitCommandStr, false)
	outReader := bufio.NewReader(cmd.outPipe)
	errReader := bufio.NewReader(cmd.errPipe)
	for {
		tmpLine, err := outReader.ReadString('\n')
		if err == nil {
			//fmt.Println(tmpLine)
			if gstr.HasPrefix(tmpLine, SplitCommandStr) {
				cmd.isInitOK = true
				break
			}
		}
		time.Sleep(time.Millisecond * 10)
	}

	go func() {
		count := 0

		//buffer := bytes.NewBuffer(nil)
		for {
			count, _ = outReader.Read(cmd.buf)

			if count > 0 && cmd.isInitOK {
				if strings.Index(string(cmd.buf[:count]), "\n"+CommandEndStr) != -1 {
					//cmd.wg.Done()
				}
				handleMsg(cmd.ID, cmd.buf[:count])
			} else {
				count, _ = errReader.Read(cmd.buf)
				if count > 0 && cmd.isInitOK {
					if strings.Index(string(cmd.buf[:count]), "\n"+CommandEndStr) != -1 {
						//cmd.wg.Done()
					}
					handleMsg(cmd.ID, cmd.buf[:count])
				}
			}

			if cmd.isClose {
				runtime.Goexit()
				break
			}
		}
	}()

}

type WindowsTerminal struct {
	Terminal
}

func NewWindowsTerminal(handleMsg func(id string, msgBytes []byte), terminalConfig *TerminalConfig, terminalId ...string) *WindowsTerminal {
	cmd := &WindowsTerminal{}
	newTerminal(&cmd.Terminal, handleMsg, terminalConfig, terminalId...)
	return cmd
}

type LinuxTerminal struct {
	Terminal
}

func NewLinuxTerminal(handleMsg func(id string, msgBytes []byte), terminalConfig *TerminalConfig, terminalId ...string) *LinuxTerminal {
	cmd := &LinuxTerminal{}
	newTerminal(&cmd.Terminal, handleMsg, terminalConfig, terminalId...)
	return cmd
}

func GetTerminalManager() *TerminalManagerStruct {
	terminalManagerOnce.Do(func() {
		terminalManager = &TerminalManagerStruct{
			Terminals: map[string]ITerminal{},
		}
	})
	return terminalManager
}

func NewTerminal(handleMsg func(id string, msgBytes []byte), terminalConfig ...*TerminalConfig) ITerminal {
	var cmd ITerminal
	GetTerminalManager()
	tmpTerminalConfig := DefaultTerminalConfig
	if len(terminalConfig) > 0 {
		tmpTerminalConfig = terminalConfig[0]
	}

	if utils.IsWindows() {
		cmd = NewWindowsTerminal(handleMsg, tmpTerminalConfig)
	} else if utils.IsLinux() {
		cmd = NewLinuxTerminal(handleMsg, tmpTerminalConfig)
	} else {
		cmd = NewLinuxTerminal(handleMsg, tmpTerminalConfig)
	}
	GetTerminalManager().Terminals[cmd.GetID()] = cmd
	return cmd
}

func GetProcessChildren(ppid int) ([]*process.Process, error) {
	p, err := process.NewProcess(int32(ppid))
	if err != nil {
		return nil, err
	}
	return p.Children()

}
