package hash

import (
	"crypto/sha512"
)

// BySha512224 encrypts by sha512_224.
// 通过 sha512_224 加密
func (e Hash) BySha512224() Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	h := sha512.New512_224()
	e.sum(h)
	bytes := h.Sum(nil)
	e.dst = bytes[:]
	return e
}
