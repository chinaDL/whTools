package crypto

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	goHash "hash"
)

func (c *Cipher) RSAGeneratePrivateKey(bits int) *Cipher {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		c.Err = err
		return c
	}

	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)

	privateBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: X509PrivateKey}
	out := bytes.NewBuffer(nil)
	c.Err = pem.Encode(out, &privateBlock)
	c.PrivateKey = out.Bytes()
	return c
}

func (c *Cipher) RSAGeneratePublicKey(priKeyOption ...[]byte) *Cipher {
	var priKey []byte
	if len(priKeyOption) > 0 {
		priKey = priKeyOption[0]
	} else {
		priKey = c.PrivateKey
	}
	block, _ := pem.Decode(priKey)
	if block == nil {
		c.Err = errors.New("key is invalid format")
		return c
	}

	// x509 parse
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		c.Err = err
		return c
	}
	publicKey := privateKey.PublicKey
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		c.Err = err
		return c
	}

	publicBlock := pem.Block{Type: "RSA PUBLIC KEY", Bytes: X509PublicKey}
	out := bytes.NewBuffer(nil)
	c.Err = pem.Encode(out, &publicBlock)
	c.PublicKey = out.Bytes()
	return c
}

func (e Encrypt) ByRSA(c *Cipher) Encrypt {
	if c.mode == "" {
		e.Err = errors.New("mode not set")
		return e
	}
	block, _ := pem.Decode(c.PublicKey)
	if block == nil {
		e.Err = errors.New("key is invalid format")
		return e
	}

	// x509 parse
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		e.Err = err
		return e
	}

	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		e.Err = errors.New("the kind of key is not a rsa.PublicKey")
		return e
	}
	// encrypt
	var dst []byte

	if c.mode == RSA_PKCS1_V1_5 {
		dst, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey, e.src)
	} else {
		//if len(c.key) == 0 {
		//	e.Err = errors.New("rsa key is not a null")
		//	return e
		//}
		var tHash goHash.Hash
		switch c.mode {
		case RSA_OAEP_SHA1:
			tHash = sha1.New()
		case RSA_OAEP_SHA256:
			tHash = sha256.New()
		case RSA_OAEP_SHA384:
			tHash = sha512.New384()
		case RSA_OAEP_SHA512:
			tHash = sha512.New()
		case RSA_OAEP_MD5:
			tHash = md5.New()
		}
		dst, err = rsa.EncryptOAEP(tHash, rand.Reader, publicKey, e.src, nil)
	}

	if err != nil {
		e.Err = err
		return e
	}
	e.dst = dst
	return e
}

func (e Decrypt) ByRSA(c *Cipher) Decrypt {
	block, _ := pem.Decode(c.PrivateKey)
	if block == nil {
		e.Err = errors.New("key is invalid format")
		return e
	}

	// x509 parse
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		e.Err = err
		return e
	}

	var dst []byte

	if c.mode == RSA_PKCS1_V1_5 {
		dst, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, e.src)
	} else {
		//if len(c.key) == 0 {
		//	e.Err = errors.New("rsa key is not a null")
		//	return e
		//}
		var tHash goHash.Hash
		switch c.mode {
		case RSA_OAEP_SHA1:
			tHash = sha1.New()
		case RSA_OAEP_SHA256:
			tHash = sha256.New()
		case RSA_OAEP_SHA384:
			tHash = sha512.New384()
		case RSA_OAEP_SHA512:
			tHash = sha512.New()
		case RSA_OAEP_MD5:
			tHash = md5.New()
		}
		dst, err = rsa.DecryptOAEP(tHash, rand.Reader, privateKey, e.src, nil)
	}

	if err != nil {
		e.Err = err
		return e
	}
	e.dst = dst
	return e
}
