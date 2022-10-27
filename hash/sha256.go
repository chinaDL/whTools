package hash

import (
	"crypto/sha256"
)

// BySha256 encrypts by sha256.
// 通过 sha256 加密
func (e Hash) BySha256() Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	h := sha256.New()
	e.sum(h)
	bytes := h.Sum(nil)
	e.dst = bytes[:]
	return e
}
