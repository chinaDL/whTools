package hash

import (
	"crypto/sha1"
)

// BySha1 encrypts by sha1.
// 通过 sha1 加密
func (e Hash) BySha1() Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	h := sha1.New()
	e.sum(h)
	bytes := h.Sum(nil)
	e.dst = bytes[:]
	return e
}
