package crypto

import (
	"crypto/cipher"
	"fmt"
	"github.com/chinaDL/whTools/encoding"
	"github.com/chinaDL/whTools/os/gfile"
	"github.com/chinaDL/whTools/utils"
	"io"
)

type BaseStruct struct {
	src       []byte
	dst       []byte
	Err       error
	isBigFile bool
	filePath  string
}

const (
	MaxFileBlock = 1 * 1024 * 1024
)

var (
	cryptFunMap = make(map[string]func())
)

func init() {
	fmt.Println("init")
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypt ecb

// NewECBEncrypt returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func newECBEncrypt(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypt)(newECB(b))
}

func (x *ecbEncrypt) BlockSize() int { return x.blockSize }

func (x *ecbEncrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypt ecb

// NewECBDecrypt returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypt(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypt)(newECB(b))
}

func (x *ecbDecrypt) BlockSize() int { return x.blockSize }

func (x *ecbDecrypt) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// String implements the interface Stringer for encode struct.
// 实现 Stringer 接口
func (e BaseStruct) String() string {
	return e.ToString()
}

// ToString outputs as string.
// 输出字符串
func (e BaseStruct) ToString() string {
	return utils.Bytes2String(e.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (e BaseStruct) ToBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}

func (e BaseStruct) sum(w io.Writer) {
	if e.isBigFile {
		_ = gfile.ReadDo(e.filePath, func(cBuf []byte) error {
			_, err := w.Write(cBuf)
			return err
		})
	} else {
		_, _ = w.Write(e.src)
	}

}

// ToHexString outputs as string with hex encoding.
// 输出经过 hex 编码的字符串
func (e BaseStruct) ToHexString() string {
	return encoding.NewEncode().FromBytes(e.dst).ByHex().ToString()
}

// ToBase32String outputs as string with base32 encoding.
// 输出经过 base32 编码的字符串
func (e BaseStruct) ToBase32String() string {
	return encoding.NewEncode().FromBytes(e.dst).ByBase32().ToString()
}

// ToBase64String outputs as string with base64 encoding.
// 输出经过 base64 编码的字符串
func (e BaseStruct) ToBase64String() string {
	return encoding.NewEncode().FromBytes(e.dst).ByBase64().ToString()
}

// ToHexBytes outputs as byte with hex encoding.
// 输出经过 hex 编码的字节切片
func (e BaseStruct) ToHexBytes() []byte {
	return encoding.NewEncode().FromBytes(e.dst).ByHex().ToBytes()
}

// ToBase32Bytes outputs as byte slice with base32 encoding.
// 输出经过 base32 编码的字节切片
func (e BaseStruct) ToBase32Bytes() []byte {
	return encoding.NewEncode().FromBytes(e.dst).ByBase32().ToBytes()
}

// ToBase64Bytes outputs as byte slice with base64 encoding.
// 输出经过 base64 编码的字节切片
func (e BaseStruct) ToBase64Bytes() []byte {
	return encoding.NewEncode().FromBytes(e.dst).ByBase64().ToBytes()
}

// returns an invalid aes key error
// 返回无效的 aes 密钥大小错误
func invalidAesKeyError(size int) error {
	return fmt.Errorf("invalid aes key size %d, the key must be 16, 24 or 32 bytes", size)
}
