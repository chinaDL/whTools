package hash

import (
	"crypto/sha256"
)

// BySha224 encrypts by sha224.
// 通过 sha224 加密
func (e Hash) BySha224() Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	h := sha256.New224()
	e.sum(h)
	bytes := h.Sum(nil)
	e.dst = bytes[:]
	return e
}
