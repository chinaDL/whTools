package crypto

// ModeType mode constants
// 模式常量
type ModeType string

const (
	CBC ModeType = "CBC"
	ECB ModeType = "ECB"
	CFB ModeType = "CFB"
	OFB ModeType = "OFB"
	CTR ModeType = "CTR"
	//GCM ModeType = "GCM"
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
