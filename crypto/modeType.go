package crypto

// ModeType mode constants
// 模式常量
type ModeType string

const (
	CBC             ModeType = "CBC"
	ECB             ModeType = "ECB"
	CFB             ModeType = "CFB"
	OFB             ModeType = "OFB"
	CTR             ModeType = "CTR"
	RSA_OAEP_SHA1   ModeType = "RSA_OAEP_SHA1"
	RSA_OAEP_MD5    ModeType = "RSA_OAEP_MD5"
	RSA_OAEP_SHA256 ModeType = "RSA_OAEP_SHA256"
	RSA_OAEP_SHA384 ModeType = "RSA_OAEP_SHA384"
	RSA_OAEP_SHA512 ModeType = "RSA_OAEP_SHA512"
	RSA_PKCS1_V1_5  ModeType = "RSA_PKCS1_V1_5"
)

func (m *ModeType) CBC() ModeType {
	return CBC
}
func (m *ModeType) ECB() ModeType {
	return ECB
}
func (m *ModeType) CFB() ModeType {
	return CFB
}
func (m *ModeType) OFB() ModeType {
	return OFB
}
func (m *ModeType) CTR() ModeType {
	return CTR
}

func (m *ModeType) RsaOaepSHA1() ModeType {
	return RSA_OAEP_SHA1
}
func (m *ModeType) RsaOaepMD5() ModeType {
	return RSA_OAEP_MD5
}
func (m *ModeType) RsaOaepSHA256() ModeType {
	return RSA_OAEP_SHA256
}
func (m *ModeType) RsaOaepSHA384() ModeType {
	return RSA_OAEP_SHA384
}
func (m *ModeType) RsaOaepSHA512() ModeType {
	return RSA_OAEP_SHA512
}
