package gfile

import (
	"bytes"
	"io"
	"os"
)

const (
	bufferCount = 1 * 1024 * 1024
)

func ReadDo(srcPath string, callback func(cBuf []byte) error) error {
	f, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := make([]byte, bufferCount)

	for {
		readSize, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if callback(buf[:readSize]) != nil {
			break
		}
	}
	return nil
}

func ReadLinesDo(srcPath string, callback func(text string) error) error {
	return ReadCharDo(srcPath, callback, '\n')
}
func ReadCharDo(srcPath string, callback func(text string) error, char byte) error {
	f, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer f.Close()
	upBuf := make([]byte, 0)
	line := ""
	buf := make([]byte, bufferCount)

exitFor:
	for {
		readSize, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		buffPosition := 0
		for {
			i := bytes.IndexByte(buf[0:], char)
			if i == -1 {
				if readSize == buffPosition {
					upBuf = upBuf[:0]
				} else {
					upBuf = buf[0:]
				}
				break
			} else {
				if len(upBuf) > 0 {
					lBy := append(upBuf, buf[:i]...)
					line = string(lBy)
					upBuf = upBuf[:0]
				} else {
					line = string(buf[:i])
				}
			}
			buffPosition += i + 1
			buf = buf[i+1:]
			if callback(line) != nil {
				break exitFor
			}
		}

	}
	return nil
}
