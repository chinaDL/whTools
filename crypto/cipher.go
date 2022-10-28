package crypto

import (
	"bytes"
	"github.com/chinaDL/whTools/utils"
	"github.com/gogf/gf/v2/os/gfile"
	"io/ioutil"
)

type paddingGroup struct {
	PaddingFun   func([]byte, int) []byte
	UnPaddingFun func([]byte) []byte
}

// Cipher defines a Cipher struct.
// 定义 Cipher 结构体
type Cipher struct {
	mode       ModeType    // 分组模式
	padding    PaddingType // 填充模式
	key        []byte      // 密钥
	iv         []byte      // 偏移向量
	PublicKey  []byte      // 公钥
	PrivateKey []byte      // 私钥
	paddingMap map[PaddingType]paddingGroup
	Err        error
}

// NewCipher returns a new Cipher instance.
// 初始化 Cipher 结构体
func NewCipher() *Cipher {
	ret := &Cipher{
		mode:    CBC,
		padding: PKCS7,
	}
	ret.paddingMap = make(map[PaddingType]paddingGroup)
	ret.paddingMap[No] = paddingGroup{
		PaddingFun: func(src []byte, blockSize int) []byte {
			return src
		},
		UnPaddingFun: func(src []byte) []byte {
			return src
		},
	}
	ret.paddingMap[Zero] = paddingGroup{
		PaddingFun:   ret.ZeroPadding,
		UnPaddingFun: ret.ZeroUnPadding,
	}
	ret.paddingMap[PKCS5] = paddingGroup{
		PaddingFun: func(src []byte, _ int) []byte {
			return ret.PKCS5UnPadding(src)
		},
		UnPaddingFun: ret.PKCS5UnPadding,
	}
	ret.paddingMap[PKCS7] = paddingGroup{
		PaddingFun:   ret.PKCS7Padding,
		UnPaddingFun: ret.PKCS7UnPadding,
	}
	return ret
}

// SetMode sets mode.
// 设置分组模式
func (c *Cipher) SetMode(mode ModeType) *Cipher {
	c.mode = mode
	return c
}

// SetPadding sets padding.
// 设置填充模式
func (c *Cipher) SetPadding(padding PaddingType) *Cipher {
	c.padding = padding
	return c
}

// SetKey sets key.
// 设置密钥
func (c *Cipher) SetKey(key interface{}) *Cipher {
	switch v := key.(type) {
	case string:
		if gfile.Exists(v) {
			c.key, _ = ioutil.ReadFile(v)
		} else {
			c.key = utils.Str2bytes(v)
		}
	case []byte:
		c.key = v
	}
	return c
}

// SetIV sets iv.
// 设置偏移向量
func (c *Cipher) SetIV(iv interface{}) *Cipher {
	switch v := iv.(type) {
	case string:
		if gfile.Exists(v) {
			c.iv, _ = ioutil.ReadFile(v)
		} else {
			c.iv = utils.Str2bytes(v)
		}
	case []byte:
		c.iv = v
	}
	return c
}

// ZeroPadding padding with Zero mode.
// 进行零填充
func (c *Cipher) ZeroPadding(src []byte, size int) []byte {
	if len(src) == 0 {
		return nil
	}
	padding := bytes.Repeat([]byte{byte(0)}, size-len(src)%size)
	return append(src, padding...)
}

// ZeroUnPadding removes padding with Zero mode.
// 移除零填充
func (c *Cipher) ZeroUnPadding(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
}

// PKCS5Padding padding with PKCS5 mode.
// 进行 PKCS5 填充
func (c *Cipher) PKCS5Padding(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}
	return c.PKCS7Padding(src, 8)
}

// PKCS5UnPadding removes padding with PKCS5 mode.
// 移除 PKCS5 填充
func (c *Cipher) PKCS5UnPadding(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}
	return c.PKCS7UnPadding(src)
}

// PKCS7Padding padding with PKCS7 mode.
// 进行 PKCS7 填充
func (c *Cipher) PKCS7Padding(src []byte, size int) []byte {
	if len(src) == 0 {
		return nil
	}
	paddingCount := size - len(src)%size
	paddingText := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	return append(src, paddingText...)
}

// PKCS7UnPadding removes padding with PKCS7 mode.
// 移除 PKCS7 填充
func (c *Cipher) PKCS7UnPadding(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}
	trim := len(src) - int(src[len(src)-1])
	if trim < 0 {
		return nil
	}
	return src[:trim]
}

func (c *Cipher) Padding(src []byte, blockSize int) []byte {
	padFun, ok := c.paddingMap[c.padding]
	if !ok {
		return src
	}
	return padFun.PaddingFun(src, blockSize)
}

func (c *Cipher) UnPadding(src []byte) []byte {
	padFun, ok := c.paddingMap[c.padding]
	if !ok {
		return src
	}
	return padFun.UnPaddingFun(src)
}

func (c *Cipher) SetPrivateKey(key interface{}) *Cipher {
	switch v := key.(type) {
	case string:
		if gfile.Exists(v) {
			c.PrivateKey, _ = ioutil.ReadFile(v)
		} else {
			c.PrivateKey = utils.Str2bytes(v)
		}
	case []byte:
		c.PrivateKey = v
	}
	return c
}

func (c *Cipher) SetPublicKey(key interface{}) *Cipher {
	switch v := key.(type) {
	case string:
		if gfile.Exists(v) {
			c.PublicKey, _ = ioutil.ReadFile(v)
		} else {
			c.PublicKey = utils.Str2bytes(v)
		}
	case []byte:
		c.PublicKey = v
	}
	return c
}
