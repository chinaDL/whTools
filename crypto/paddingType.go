package crypto

// PaddingType padding constants
// 填充常量
type PaddingType string

const (
	No    PaddingType = "no"
	Zero  PaddingType = "zero"
	PKCS5 PaddingType = "pkcs5"
	PKCS7 PaddingType = "pkcs7"
)

func (p *PaddingType) No() PaddingType {
	return No
}

func (p *PaddingType) Zero() PaddingType {
	return Zero
}

func (p *PaddingType) PKCS5() PaddingType {
	return PKCS5
}

func (p *PaddingType) PKCS7() PaddingType {
	return PKCS7
}
