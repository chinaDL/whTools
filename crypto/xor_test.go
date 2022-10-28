package crypto

import (
	"fmt"
	"testing"
)

func TestXor(t *testing.T) {
	content := "测试一下123!@#"
	key := []byte("abcd1234")
	en := NewEncrypt()
	de := NewDecrypt()
	e := en.FromString(content).ByXor(key).ToBase64String()
	d := de.FromBase64String(e).ByXor(key).ToString()
	fmt.Println(e)
	fmt.Println(d)
}
