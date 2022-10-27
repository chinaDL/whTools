package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"github.com/chinaDL/whTools/utils"
	"hash"
)

// ByHmacSha512 encrypts by hmac_sha512.
// 通过 hmac_sha512 加密
func (e Hash) ByHmacSha512(key interface{}) Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha512.New, utils.Str2bytes(v))
	case []byte:
		mac = hmac.New(sha512.New, v)
	}
	e.sum(mac)
	e.dst = mac.Sum(nil)
	return e
}
