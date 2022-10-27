package hash

import (
	"crypto/sha512"
)

// BySha512 encrypts by sha512.
// 通过 sha512 加密
func (e Hash) BySha512() Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	h := sha512.New()
	e.sum(h)
	bytes := h.Sum(nil)
	e.dst = bytes[:]
	return e
}
