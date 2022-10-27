package hash

import (
	"crypto/sha512"
)

// BySha512256 encrypts by sha512_256.
// 通过 sha512_256 加密
func (e Hash) BySha512256() Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	h := sha512.New512_256()
	e.sum(h)
	bytes := h.Sum(nil)
	e.dst = bytes[:]
	return e
}
