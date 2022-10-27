package encoding

import (
	"github.com/chinaDL/whTools/utils"
)

// 该部分代码参考于 https://github.com/golang-module/dongle

type Decode struct {
	BaseStruct
}

func NewDecode() Decode {
	return Decode{}
}

// FromString decodes from string.
// 对字符串进行解密
func (d Decode) FromString(s string) Decode {
	d.src = utils.Str2bytes(s)
	return d
}

// FromBytes decodes from byte slice.
// 对字节切片进行解密
func (d Decode) FromBytes(b []byte) Decode {
	d.src = b
	return d
}

// String implements the interface Stringer for decode struct.
// 实现 Stringer 接口
func (d Decode) String() string {
	return d.ToString()
}

// ToString outputs as string.
// 输出字符串
func (d Decode) ToString() string {
	return utils.Bytes2String(d.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (d Decode) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}
