package crypto

func xor(src []byte, key []byte) []byte {
	dst := make([]byte, len(src))
	keyLen := len(key)
	for i := 0; i < len(src); i++ {
		dst[i] = src[i] ^ key[i%keyLen]
	}
	return dst
}

func (e Encrypt) ByXor(key []byte) Encrypt {
	e.dst = xor(e.src, key)
	return e
}

func (e Decrypt) ByXor(key []byte) Decrypt {
	e.dst = xor(e.src, key)
	return e
}
