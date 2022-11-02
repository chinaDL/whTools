package crypto

import (
	"fmt"
	"testing"
)

func TestRC4(t *testing.T) {
	en := NewEncrypt()
	de := NewDecrypt()
	c := NewCipher().SetKey("abcd1234")
	content := "测试加密123!"
	e := en.FromString(content).ByRC4(c).ToBase64String()
	fmt.Println(e)
	fmt.Println(de.FromBase64String(e).ByRC4(c).ToString())
}
