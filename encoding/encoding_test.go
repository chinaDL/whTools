package encoding

import (
	"fmt"
	"testing"
)

func TestEncoding(t *testing.T) {
	encode := NewEncode()
	decode := NewDecode()
	en := "测试abc123"
	de := encode.FromString(en).ByBase32().ToString()
	fmt.Println("base32 encode:", de)
	fmt.Println("base32 decode:", decode.FromString(de).ByBase32().ToString())

	de = encode.FromString(en).ByBase45().ToString()
	fmt.Println("base45 encode:", de)
	fmt.Println("base45 decode:", decode.FromString(de).ByBase45().ToString())

	de = encode.FromString(en).ByBase58().ToString()
	fmt.Println("base58 encode:", de)
	fmt.Println("base58 decode:", decode.FromString(de).ByBase58().ToString())

	de = encode.FromString(en).ByBase64().ToString()
	fmt.Println("base64 encode:", de)
	fmt.Println("base64 decode:", decode.FromString(de).ByBase64().ToString())

	de = encode.FromString(en).ByBase64URL().ToString()
	fmt.Println("base64Url encode:", de)
	fmt.Println("base64Url decode:", decode.FromString(de).ByBase64URL().ToString())
}
