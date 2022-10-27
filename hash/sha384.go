package hash

import (
	"crypto/sha512"
)

// BySha384 encrypts by sha384.
// 通过 sha384 加密
func (e Hash) BySha384() Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	h := sha512.New384()
	e.sum(h)
	bytes := h.Sum(nil)
	e.dst = bytes[:]
	return e
}
