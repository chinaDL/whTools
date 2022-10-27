package encoding

import (
	"github.com/chinaDL/whTools/errs"
	"github.com/chinaDL/whTools/utils"
	"io/ioutil"
)

// 该部分代码参考于 https://github.com/golang-module/dongle

type Encode struct {
	BaseStruct
}

func NewEncode() Encode {
	return Encode{}
}

// FromString encodes from string.
// 对字符串进行编码
func (e Encode) FromString(s string) Encode {
	e.src = utils.Str2bytes(s)
	return e
}

// FromBytes encodes from byte slice.
// 对字节切片进行编码
func (e Encode) FromBytes(b []byte) Encode {
	e.src = b
	return e
}

// FromFile encodes from file.
// 对文件进行编码
func (e Encode) FromFile(f interface{}) Encode {
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

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		e.Err = errs.InvalidFileError(filename)
		return e
	}
	e.src = bytes
	return e
}

// String implements the interface Stringer for encode struct.
// 实现 Stringer 接口
func (e Encode) String() string {
	return e.ToString()
}

// ToString outputs as string.
// 输出字符串
func (e Encode) ToString() string {
	return utils.Bytes2String(e.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (e Encode) ToBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}

// ToHexString outputs as string with hex encoding.
// 输出经过 hex 编码的字符串
func (e Encode) ToHexString() string {
	return NewEncode().FromBytes(e.dst).ByHex().ToString()
}

// ToBase32String outputs as string with base32 encoding.
// 输出经过 base32 编码的字符串
func (e Encode) ToBase32String() string {
	return NewEncode().FromBytes(e.dst).ByBase32().ToString()
}

// ToBase64String outputs as string with base64 encoding.
// 输出经过 base64 编码的字符串
func (e Encode) ToBase64String() string {
	return NewEncode().FromBytes(e.dst).ByBase64().ToString()
}

// ToHexBytes outputs as byte with hex encoding.
// 输出经过 hex 编码的字节切片
func (e Encode) ToHexBytes() []byte {
	return NewEncode().FromBytes(e.dst).ByHex().ToBytes()
}

// ToBase32Bytes outputs as byte slice with base32 encoding.
// 输出经过 base32 编码的字节切片
func (e Encode) ToBase32Bytes() []byte {
	return NewEncode().FromBytes(e.dst).ByBase32().ToBytes()
}

// ToBase64Bytes outputs as byte slice with base64 encoding.
// 输出经过 base64 编码的字节切片
func (e Encode) ToBase64Bytes() []byte {
	return NewEncode().FromBytes(e.dst).ByBase64().ToBytes()
}
