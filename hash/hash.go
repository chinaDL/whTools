package hash

import (
	"errors"
	"github.com/chinaDL/whTools/errs"
	"github.com/chinaDL/whTools/utils"
	"io/ioutil"
	"os"
)

const (
	MaxFileBlock = 10 * 1024 * 1024
)

type Hash struct {
	BaseStruct
}

func NewHash() Hash {
	return Hash{}
}

// FromString hash from string.
// 对字符串进行编码
func (e Hash) FromString(s string) Hash {
	e.src = utils.Str2bytes(s)
	return e
}

// FromBytes hash from byte slice.
// 对字节切片进行编码
func (e Hash) FromBytes(b []byte) Hash {
	e.src = b
	return e
}

// FromFile hash from file.
// 对文件进行编码
func (e Hash) FromFile(f interface{}) Hash {
	filename := ""
	switch v := f.(type) {
	case string:
		filename = v
	case []byte:
		filename = utils.Bytes2String(v)
	}
	if len(filename) == 0 {
		return e
	}
	tFile, err := os.Open(filename)
	if err != nil {
		e.Err = errors.New(err.Error())
		return e
	}
	defer tFile.Close()
	fInfo, _ := tFile.Stat()
	if fInfo.Size() >= MaxFileBlock {
		e.src = nil
		e.isBigFile = true
		e.filePath = filename
		return e
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		e.Err = errs.InvalidFileError(filename)
		return e
	}
	e.src = bytes
	e.isBigFile = false
	return e
}
