package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"github.com/chinaDL/whTools/utils"
	"hash"
)

// ByHmacSha224 encrypts by hmac_sha224.
// 通过 hmac_sha224 加密
func (e Hash) ByHmacSha224(key interface{}) Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha256.New224, utils.Str2bytes(v))
	case []byte:
		mac = hmac.New(sha256.New224, v)
	}
	e.sum(mac)
	e.dst = mac.Sum(nil)
	return e
}
