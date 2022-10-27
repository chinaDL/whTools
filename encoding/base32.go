package encoding

import (
	"encoding/base32"
)

// 该部分代码参考于 https://github.com/golang-module/dongle

// ByBase32 encodes by base32.
// 通过 base32 编码
func (e Encode) ByBase32() Encode {
	if len(e.src) == 0 {
		return e
	}
	buf := make([]byte, base32.StdEncoding.EncodedLen(len(e.src)))
	base32.StdEncoding.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByBase32 decodes by base32.
// 通过 base32 解码
func (d Decode) ByBase32() Decode {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, base32.StdEncoding.DecodedLen(len(d.src)))
	n, err := base32.StdEncoding.Decode(buf, d.src)
	d.dst, d.Err = buf[:n], err
	return d
}
