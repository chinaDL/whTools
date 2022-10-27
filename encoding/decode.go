package encoding

import (
	"errors"
	"github.com/chinaDL/whTools/errs"
	"github.com/chinaDL/whTools/utils"
	"io/ioutil"
	"os"
)

// 该部分代码参考于 https://github.com/golang-module/dongle

type Decode struct {
	BaseStruct
}

func NewDecode() Decode {
	return Decode{}
}

// FromString hash from string.
// 对字符串进行编码
func (d Decode) FromString(s string) Decode {
	d.src = utils.Str2bytes(s)
	return d
}

// FromBytes hash from byte slice.
// 对字节切片进行编码
func (d Decode) FromBytes(b []byte) Decode {
	d.src = b
	return d
}

// FromFile hash from file.
// 对文件进行编码
func (d Decode) FromFile(f interface{}) Decode {
	filename := ""
	switch v := f.(type) {
	case string:
		filename = v
	case []byte:
		filename = utils.Bytes2String(v)
	}
	if len(filename) == 0 {
		return d
	}
	tFile, err := os.Open(filename)
	if err != nil {
		d.Err = errors.New(err.Error())
		return d
	}
	defer tFile.Close()
	fInfo, _ := tFile.Stat()
	if fInfo.Size() >= MaxFileBlock {
		d.src = nil
		d.isBigFile = true
		d.filePath = filename
		return d
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		d.Err = errs.InvalidFileError(filename)
		return d
	}
	d.src = bytes
	d.isBigFile = false
	return d
}
