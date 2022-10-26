package gstr

const (
	Whitespace         = "\t\n\r\v\f"
	AsciiLowercase     = "abcdefghijklmnopqrstuvwxyz"
	AsciiUppercase     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AsciiLetters       = AsciiLowercase + AsciiUppercase
	Digits             = "0123456789"
	Hexdigits          = Digits + "abcdef" + "ABCDEF"
	HexdigitsLowercase = Digits + "abcdef"
	HexdigitsUppercase = Digits + "ABCDEF"
	Octdigits          = "01234567"
	Punctuation        = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	Printable          = Digits + AsciiLetters + Punctuation + Whitespace
)
