package hash

import (
	"crypto/hmac"
	"crypto/sha1"
	"github.com/chinaDL/whTools/utils"
	"hash"
)

// ByHmacSha1 encrypts by hmac_sha1.
// 通过 hmac_sha1 加密
func (e Hash) ByHmacSha1(key interface{}) Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha1.New, utils.Str2bytes(v))
	case []byte:
		mac = hmac.New(sha1.New, v)
	}
	e.sum(mac)
	e.dst = mac.Sum(nil)
	return e
}
