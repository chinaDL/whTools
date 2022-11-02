package crypto

import "crypto/rc4"

func rc4Encrypt(key, src []byte) []byte {
	dst := make([]byte, len(src))
	cipher, _ := rc4.NewCipher(key)
	cipher.XORKeyStream(dst, src)
	return dst
}
func (e Encrypt) ByRC4(c *Cipher) Encrypt {
	e.dst = rc4Encrypt(c.key, e.src)
	return e
}

func (e Decrypt) ByRC4(c *Cipher) Decrypt {
	e.dst = rc4Encrypt(c.key, e.src)
	return e
}
