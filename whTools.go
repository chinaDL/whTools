package whTools

import (
	"github.com/chinaDL/whTools/encoding"
	"github.com/chinaDL/whTools/hash"
)

var (
	Encode = encoding.NewEncode()
	Decode = encoding.NewDecode()
	Hash   = hash.NewHash()
)
