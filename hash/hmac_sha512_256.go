package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"github.com/chinaDL/whTools/utils"
	"hash"
)

// ByHmacSha512256 encrypts by hmac_sha512_256.
// 通过 hmac_sha512_256 加密
func (e Hash) ByHmacSha512256(key interface{}) Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha512.New512_256, utils.Str2bytes(v))
	case []byte:
		mac = hmac.New(sha512.New512_256, v)
	}
	e.sum(mac)
	e.dst = mac.Sum(nil)
	return e
}
