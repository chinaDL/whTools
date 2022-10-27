package encoding

import "encoding/base64"

// 该部分代码参考于 https://github.com/golang-module/dongle

// ByBase64 encodes by base64.
// 通过 base64 编码
func (e Encode) ByBase64() Encode {
	if len(e.src) == 0 {
		return e
	}
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(e.src)))
	base64.StdEncoding.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByBase64 decodes by base64.
// 通过 base64 解码
func (d Decode) ByBase64() Decode {
	if len(d.src) == 0 {
		return d
	}
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(d.src)))
	n, err := base64.StdEncoding.Decode(buf, d.src)
	d.dst, d.Err = buf[:n], err
	return d
}
