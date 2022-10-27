package crypto

import (
	"fmt"
	"testing"
)

func TestAesMode(t *testing.T) {
	content := "测试加密123!"
	key := "1234567890abcdef"
	encrypt := NewEncrypt()
	decrypt := NewDecrypt()
	p := new(PaddingType)
	m := new(ModeType)
	ci := NewCipher()
	ci.SetPadding(p.Zero())
	ci.SetKey(key)
	ci.SetIV(key)

	ci.SetMode(m.ECB())
	enB64 := encrypt.FromString(content).ByAes(ci).ToBase64String()
	fmt.Println("aes ecb en:", enB64)
	deStr := decrypt.FromBase64String(enB64).ByAes(ci).ToString()
	fmt.Println("aes ecb de:", deStr)
	if deStr != content {
		t.Error("ECB 错误")
	}

	ci.SetMode(m.CBC())
	enB64 = encrypt.FromString(content).ByAes(ci).ToBase64String()
	fmt.Println("aes cbc en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByAes(ci).ToString()
	fmt.Println("aes cbc de:", deStr)
	if deStr != content {
		t.Error("cbc 错误")
	}

	ci.SetMode(m.OFB())
	enB64 = encrypt.FromString(content).ByAes(ci).ToBase64String()
	fmt.Println("aes ofb en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByAes(ci).ToString()
	fmt.Println("aes ofb de:", deStr)
	if deStr != content {
		t.Error("ofb 错误")
	}

	ci.SetMode(m.CTR())
	enB64 = encrypt.FromString(content).ByAes(ci).ToBase64String()
	fmt.Println("aes ctr en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByAes(ci).ToString()
	fmt.Println("aes ctr de:", deStr)
	if deStr != content {
		t.Error("ECB 错误")
	}

	ci.SetMode(m.CFB())
	enB64 = encrypt.FromString(content).ByAes(ci).ToBase64String()
	fmt.Println("aes cfb en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByAes(ci).ToString()
	fmt.Println("aes cfb de:", deStr)
	if deStr != content {
		t.Error("CFB 错误")
	}

}
