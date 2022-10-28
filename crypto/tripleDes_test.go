package crypto

import (
	"fmt"
	"testing"
)

func TestTripleDes(t *testing.T) {
	content := "测试加密123!"
	key := "1234abcd1234abcd1234abcd"
	iv := "1234abcd"
	encrypt := NewEncrypt()
	decrypt := NewDecrypt()
	p := new(PaddingType)
	m := new(ModeType)
	ci := NewCipher()
	ci.SetPadding(p.Zero())
	ci.SetKey(key)
	ci.SetIV(iv)

	ci.SetMode(m.ECB())
	enB64 := encrypt.FromString(content).ByTripleDes(ci).ToBase64String()
	fmt.Println("3des ecb en:", enB64)
	deStr := decrypt.FromBase64String(enB64).ByTripleDes(ci).ToString()
	fmt.Println("3des ecb de:", deStr)
	if deStr != content {
		t.Error("ECB 错误")
	}

	ci.SetMode(m.CBC())
	enB64 = encrypt.FromString(content).ByTripleDes(ci).ToBase64String()
	fmt.Println("3des cbc en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByTripleDes(ci).ToString()
	fmt.Println("3des cbc de:", deStr)
	if deStr != content {
		t.Error("cbc 错误")
	}

	ci.SetMode(m.OFB())
	enB64 = encrypt.FromString(content).ByTripleDes(ci).ToBase64String()
	fmt.Println("3des ofb en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByTripleDes(ci).ToString()
	fmt.Println("3des ofb de:", deStr)
	if deStr != content {
		t.Error("ofb 错误")
	}

	ci.SetMode(m.CTR())
	enB64 = encrypt.FromString(content).ByTripleDes(ci).ToBase64String()
	fmt.Println("3des ctr en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByTripleDes(ci).ToString()
	fmt.Println("3des ctr de:", deStr)
	if deStr != content {
		t.Error("ECB 错误")
	}

	ci.SetMode(m.CFB())
	enB64 = encrypt.FromString(content).ByTripleDes(ci).ToBase64String()
	fmt.Println("3des cfb en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByTripleDes(ci).ToString()
	fmt.Println("3des cfb de:", deStr)
	if deStr != content {
		t.Error("CFB 错误")
	}
}
