package hash

import (
	"crypto/md5"
)

// ByMd5 encrypts by md5.
// 通过 md5 加密
func (e Hash) ByMd5() Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	h := md5.New()
	e.sum(h)
	bytes := h.Sum(nil)
	e.dst = bytes[:]
	return e
}
