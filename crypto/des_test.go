package crypto

import (
	"fmt"
	"testing"
)

func TestDes(t *testing.T) {
	content := "测试加密123!"
	key := "1234abcd"
	encrypt := NewEncrypt()
	decrypt := NewDecrypt()
	p := new(PaddingType)
	m := new(ModeType)
	ci := NewCipher()
	ci.SetPadding(p.Zero())
	ci.SetKey(key)
	ci.SetIV(key)

	ci.SetMode(m.ECB())
	enB64 := encrypt.FromString(content).ByDes(ci).ToBase64String()
	fmt.Println("des ecb en:", enB64)
	deStr := decrypt.FromBase64String(enB64).ByDes(ci).ToString()
	fmt.Println("des ecb de:", deStr)
	if deStr != content {
		t.Error("ECB 错误")
	}

	ci.SetMode(m.CBC())
	enB64 = encrypt.FromString(content).ByDes(ci).ToBase64String()
	fmt.Println("des cbc en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByDes(ci).ToString()
	fmt.Println("des cbc de:", deStr)
	if deStr != content {
		t.Error("cbc 错误")
	}

	ci.SetMode(m.OFB())
	enB64 = encrypt.FromString(content).ByDes(ci).ToBase64String()
	fmt.Println("des ofb en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByDes(ci).ToString()
	fmt.Println("des ofb de:", deStr)
	if deStr != content {
		t.Error("ofb 错误")
	}

	ci.SetMode(m.CTR())
	enB64 = encrypt.FromString(content).ByDes(ci).ToBase64String()
	fmt.Println("des ctr en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByDes(ci).ToString()
	fmt.Println("des ctr de:", deStr)
	if deStr != content {
		t.Error("ECB 错误")
	}

	ci.SetMode(m.CFB())
	enB64 = encrypt.FromString(content).ByDes(ci).ToBase64String()
	fmt.Println("des cfb en:", enB64)
	deStr = decrypt.FromBase64String(enB64).ByDes(ci).ToString()
	fmt.Println("des cfb de:", deStr)
	if deStr != content {
		t.Error("CFB 错误")
	}

}
