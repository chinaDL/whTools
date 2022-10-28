package crypto

import (
	"fmt"
	"testing"
)

func TestRSA(t *testing.T) {
	en := NewEncrypt()
	de := NewDecrypt()
	//encode := encoding.NewEncode()
	c := NewCipher()
	c.SetPrivateKey("priKey.pem").SetPublicKey("pubKey.pem")
	mode := new(ModeType)
	c.SetKey("abcd1234")
	//c.RSAGeneratePrivateKey(1024).RSAGeneratePublicKey()
	//ioutil.WriteFile("priKey.pem", c.PrivateKey, 0777)
	//ioutil.WriteFile("pubKey.pem", c.PublicKey, 0777)
	fmt.Println(string(c.PrivateKey))
	fmt.Println(string(c.PublicKey))

	enBs64 := en.FromString("测试123!@#").ByRSA(c.SetMode(mode.RsaOaepMD5())).ToBase64String()
	fmt.Println(enBs64)

	//deStr := de.FromBase64String(enBs64).ByRSA(c).ToString()
	//fmt.Println(deStr)

	fmt.Println(de.FromBase64String(enBs64).ByRSA(c).ToString())

}
