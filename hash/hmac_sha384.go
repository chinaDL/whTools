package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"github.com/chinaDL/whTools/utils"
	"hash"
)

// ByHmacSha384 encrypts by hmac_sha384.
// 通过 hmac_sha384 加密
func (e Hash) ByHmacSha384(key interface{}) Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha512.New384, utils.Str2bytes(v))
	case []byte:
		mac = hmac.New(sha512.New384, v)
	}
	e.sum(mac)
	e.dst = mac.Sum(nil)
	return e
}
