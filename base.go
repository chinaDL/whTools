package whTools

import (
	"github.com/chinaDL/whTools/encoding"
	"github.com/chinaDL/whTools/os/gfile"
	"github.com/chinaDL/whTools/utils"
	"io"
)

type BaseStruct struct {
	src       []byte
	dst       []byte
	Err       error
	isBigFile bool
	filePath  string
}

// String implements the interface Stringer for encode struct.
// 实现 Stringer 接口
func (e BaseStruct) String() string {
	return e.ToString()
}

// ToString outputs as string.
// 输出字符串
func (e BaseStruct) ToString() string {
	return utils.Bytes2String(e.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (e BaseStruct) ToBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}

func (e BaseStruct) sum(w io.Writer) {
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
func (e BaseStruct) ToHexString() string {
	return encoding.NewEncode().FromBytes(e.dst).ByHex().ToString()
}

// ToBase32String outputs as string with base32 encoding.
// 输出经过 base32 编码的字符串
func (e BaseStruct) ToBase32String() string {
	return encoding.NewEncode().FromBytes(e.dst).ByBase32().ToString()
}

// ToBase64String outputs as string with base64 encoding.
// 输出经过 base64 编码的字符串
func (e BaseStruct) ToBase64String() string {
	return encoding.NewEncode().FromBytes(e.dst).ByBase64().ToString()
}

// ToHexBytes outputs as byte with hex encoding.
// 输出经过 hex 编码的字节切片
func (e BaseStruct) ToHexBytes() []byte {
	return encoding.NewEncode().FromBytes(e.dst).ByHex().ToBytes()
}

// ToBase32Bytes outputs as byte slice with base32 encoding.
// 输出经过 base32 编码的字节切片
func (e BaseStruct) ToBase32Bytes() []byte {
	return encoding.NewEncode().FromBytes(e.dst).ByBase32().ToBytes()
}

// ToBase64Bytes outputs as byte slice with base64 encoding.
// 输出经过 base64 编码的字节切片
func (e BaseStruct) ToBase64Bytes() []byte {
	return encoding.NewEncode().FromBytes(e.dst).ByBase64().ToBytes()
}
