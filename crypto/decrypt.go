package crypto

import (
	"errors"
	"github.com/chinaDL/whTools/encoding"
	"github.com/chinaDL/whTools/errs"
	"github.com/chinaDL/whTools/utils"
	"io/ioutil"
	"os"
)

type Decrypt struct {
	BaseStruct
}

func NewDecrypt() Decrypt {
	return Decrypt{}
}

// FromString hash from string.
// 对字符串进行编码
func (e Decrypt) FromString(s string) Decrypt {
	e.src = utils.Str2bytes(s)
	return e
}

// FromBytes hash from byte slice.
// 对字节切片进行编码
func (e Decrypt) FromBytes(b []byte) Decrypt {
	e.src = b
	return e
}
func (e Decrypt) FromHexString(s string) Decrypt {
	e.src = encoding.NewDecode().FromString(s).ByHex().ToBytes()
	return e
}

func (e Decrypt) FromBase64String(s string) Decrypt {
	e.src = encoding.NewDecode().FromString(s).ByBase64().ToBytes()
	return e
}

func (e Decrypt) FromBase32String(s string) Decrypt {
	e.src = encoding.NewDecode().FromString(s).ByBase32().ToBytes()
	return e
}

func (e Decrypt) FromHexByte(b []byte) Decrypt {
	e.src = encoding.NewDecode().FromBytes(b).ByHex().ToBytes()
	return e
}

func (e Decrypt) FromBase64Byte(b []byte) Decrypt {
	e.src = encoding.NewDecode().FromBytes(b).ByBase64().ToBytes()
	return e
}

func (e Decrypt) FromBase32Byte(b []byte) Decrypt {
	e.src = encoding.NewDecode().FromBytes(b).ByBase32().ToBytes()
	return e
}

// FromFile hash from file.
// 对文件进行编码
func (e Decrypt) FromFile(f interface{}) Decrypt {
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
