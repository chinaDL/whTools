package hash

import (
	"errors"
	"github.com/chinaDL/whTools/encoding"
	"github.com/chinaDL/whTools/errs"
	"github.com/chinaDL/whTools/os/gfile"
	"github.com/chinaDL/whTools/utils"
	"io"
	"io/ioutil"
	"os"
)

const (
	MaxFileBlock = 10 * 1024 * 1024
)

type Hash struct {
	BaseStruct
	isBigFile bool
	filePath  string
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

// String implements the interface Stringer for encode struct.
// 实现 Stringer 接口
func (e Hash) String() string {
	return e.ToString()
}

// ToString outputs as string.
// 输出字符串
func (e Hash) ToString() string {
	return utils.Bytes2String(e.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (e Hash) ToBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}

func (e Hash) sum(w io.Writer) {
	if e.isBigFile {
		_ = gfile.ReadDo(e.filePath, func(cBuf []byte) error {
			_, err := w.Write(cBuf)
			return err
		})
	} else {
		_, _ = w.Write(e.src)
	}

}

// ToHexString outputs as string with hex encoding.
// 输出经过 hex 编码的字符串
func (e Hash) ToHexString() string {
	return encoding.NewEncode().FromBytes(e.dst).ByHex().ToString()
}

// ToBase32String outputs as string with base32 encoding.
// 输出经过 base32 编码的字符串
func (e Hash) ToBase32String() string {
	return encoding.NewEncode().FromBytes(e.dst).ByBase32().ToString()
}

// ToBase64String outputs as string with base64 encoding.
// 输出经过 base64 编码的字符串
func (e Hash) ToBase64String() string {
	return encoding.NewEncode().FromBytes(e.dst).ByBase64().ToString()
}

// ToHexBytes outputs as byte with hex encoding.
// 输出经过 hex 编码的字节切片
func (e Hash) ToHexBytes() []byte {
	return encoding.NewEncode().FromBytes(e.dst).ByHex().ToBytes()
}

// ToBase32Bytes outputs as byte slice with base32 encoding.
// 输出经过 base32 编码的字节切片
func (e Hash) ToBase32Bytes() []byte {
	return encoding.NewEncode().FromBytes(e.dst).ByBase32().ToBytes()
}

// ToBase64Bytes outputs as byte slice with base64 encoding.
// 输出经过 base64 编码的字节切片
func (e Hash) ToBase64Bytes() []byte {
	return encoding.NewEncode().FromBytes(e.dst).ByBase64().ToBytes()
}
