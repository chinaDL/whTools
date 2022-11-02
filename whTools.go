package whTools

import (
	"github.com/chinaDL/whTools/crypto"
	"github.com/chinaDL/whTools/encoding"
	"github.com/chinaDL/whTools/hash"
	"github.com/chinaDL/whTools/tools/terminal"
)

var (
	Encode      = encoding.NewEncode()
	Decode      = encoding.NewDecode()
	Hash        = hash.NewHash()
	Encrypt     = crypto.NewEncrypt()
	Decrypt     = crypto.NewDecrypt()
	PaddingType crypto.PaddingType
	ModeType    crypto.ModeType
	NewTerminal = terminal.NewTerminal
)

func NewCipher() *crypto.Cipher {

	return crypto.NewCipher()
}
