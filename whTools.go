package whTools

import (
	"github.com/chinaDL/whTools/crypto"
	"github.com/chinaDL/whTools/encoding"
	"github.com/chinaDL/whTools/hash"
)

var (
	Encode      = encoding.NewEncode()
	Decode      = encoding.NewDecode()
	Hash        = hash.NewHash()
	Encrypt     = crypto.NewEncrypt()
	Decrypt     = crypto.NewDecrypt()
	PaddingType crypto.PaddingType
	ModeType    crypto.ModeType
)

func NewCipher() *crypto.Cipher {

	return crypto.NewCipher()
}
