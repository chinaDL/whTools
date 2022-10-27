package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"github.com/chinaDL/whTools/utils"
	"hash"
)

// ByHmacSha256 encrypts by hmac_sha256.
// 通过 hmac_sha256 加密
func (e Hash) ByHmacSha256(key interface{}) Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha256.New, utils.Str2bytes(v))
	case []byte:
		mac = hmac.New(sha256.New, v)
	}
	e.sum(mac)
	e.dst = mac.Sum(nil)
	return e
}
