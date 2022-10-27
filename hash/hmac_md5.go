package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"github.com/chinaDL/whTools/utils"
	"hash"
)

// ByHmacMd5 encrypts by hmac_md5.
// 通过 hmac_md5 加密
func (e Hash) ByHmacMd5(key interface{}) Hash {
	if len(e.src) == 0 && !e.isBigFile {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(md5.New, utils.Str2bytes(v))
	case []byte:
		mac = hmac.New(md5.New, v)
	}
	e.sum(mac)
	e.dst = mac.Sum(nil)
	return e
}
