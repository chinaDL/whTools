package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

func (e Encrypt) aesEncrypt(c *Cipher) Encrypt {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		e.Err = invalidAesKeyError(len(c.key))
		return e
	}
	newBy := c.Padding(e.src, block.BlockSize())
	blockSize := block.BlockSize()
	e.dst = make([]byte, len(newBy))

	switch c.mode {
	case ECB:
		newECBEncrypt(block).CryptBlocks(e.dst, newBy)
	case CBC:
		cipher.NewCBCEncrypter(block, c.iv).CryptBlocks(e.dst, newBy)
	case CFB:
		cipher.NewCFBEncrypter(block, c.iv[:blockSize]).XORKeyStream(e.dst, newBy)
	case OFB:
		cipher.NewOFB(block, c.iv[:blockSize]).XORKeyStream(e.dst, newBy)
	case CTR:
		cipher.NewCTR(block, c.iv[:blockSize]).XORKeyStream(e.dst, newBy)
	}
	return e
}

// ByAes encrypts by aes algorithm.
// 通过 aes 加密
func (e Encrypt) ByAes(c *Cipher) Encrypt {
	return e.aesEncrypt(c)
}

func (e Decrypt) aesDecrypt(c *Cipher) Decrypt {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		e.Err = invalidAesKeyError(len(c.key))
		return e
	}
	newBy := e.src
	//decryptData := make([]byte, len(newBy))
	blockSize := block.BlockSize()
	e.dst = make([]byte, len(newBy))

	switch c.mode {
	case ECB:
		NewECBDecrypt(block).CryptBlocks(e.dst, newBy)
	case CBC:
		cipher.NewCBCDecrypter(block, c.iv).CryptBlocks(e.dst, newBy)
	case CFB:
		cipher.NewCFBDecrypter(block, c.iv[:blockSize]).XORKeyStream(e.dst, newBy)
	case OFB:
		cipher.NewOFB(block, c.iv[:blockSize]).XORKeyStream(e.dst, newBy)
	case CTR:
		cipher.NewCTR(block, c.iv[:blockSize]).XORKeyStream(e.dst, newBy)
	}
	e.dst = c.UnPadding(e.dst)
	return e
}

// ByAes decrypts by aes algorithm.
// 通过 aes 解密
func (e Decrypt) ByAes(c *Cipher) Decrypt {

	return e.aesDecrypt(c)
}
