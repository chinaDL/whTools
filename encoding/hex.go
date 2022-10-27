package encoding

import "encoding/hex"

// 该部分代码参考于 https://github.com/golang-module/dongle

// ByHex encodes by hex.
// 通过 hex 编码
func (e Encode) ByHex() Encode {
	buf := make([]byte, hex.EncodedLen(len(e.src)))
	hex.Encode(buf, e.src)
	e.dst = buf
	return e
}

// ByHex decodes by hex.
// 通过 hex 解码
func (d Decode) ByHex() Decode {
	buf := make([]byte, hex.DecodedLen(len(d.src)))
	n, err := hex.Decode(buf, d.src)
	if n > 0 {
		d.dst = buf
	}
	d.Err = err
	return d
}
