package hash

import (
	"fmt"
	"github.com/chinaDL/whTools/utils"
	"testing"
)

func TestHash(t *testing.T) {
	h := NewHash()
	fPath := "/Users/whaz/Downloads/rustdesk-1.1.9.dmg"
	sDiff := utils.ExecTimeDiff(func() {
		hValue := h.FromFile(fPath).BySha1().ToHexString()
		fmt.Println("文件", fPath, "Hash sha1:", hValue)
		hValue = h.FromString(fPath).BySha1().ToHexString()
		fmt.Println("字符", fPath, "Hash sha1:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := h.FromFile(fPath).BySha224().ToHexString()
		fmt.Println("文件", fPath, "Hash sha224:", hValue)
		hValue = h.FromString(fPath).BySha224().ToHexString()
		fmt.Println("字符", fPath, "Hash sha224:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := h.FromFile(fPath).BySha256().ToHexString()
		fmt.Println("文件", fPath, "Hash sha256:", hValue)
		hValue = h.FromString(fPath).BySha256().ToHexString()
		fmt.Println("字符", fPath, "Hash sha256:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := h.FromFile(fPath).BySha384().ToHexString()
		fmt.Println("文件", fPath, "Hash sha384:", hValue)
		hValue = h.FromString(fPath).BySha384().ToHexString()
		fmt.Println("字符", fPath, "Hash sha384:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := h.FromFile(fPath).BySha512().ToHexString()
		fmt.Println("文件", fPath, "Hash sha512:", hValue)
		hValue = h.FromString(fPath).BySha512().ToHexString()
		fmt.Println("字符", fPath, "Hash sha512:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := h.FromFile(fPath).BySha512256().ToHexString()
		fmt.Println("文件", fPath, "Hash sha512/256:", hValue)
		hValue = h.FromString(fPath).BySha512256().ToHexString()
		fmt.Println("字符", fPath, "Hash sha512/256:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := h.FromFile(fPath).BySha512224().ToHexString()
		fmt.Println("文件", fPath, "Hash sha512/224:", hValue)
		hValue = h.FromString(fPath).BySha512224().ToHexString()
		fmt.Println("字符", fPath, "Hash sha512/224:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := h.FromFile(fPath).ByMd5().ToHexString()
		fmt.Println("文件", fPath, "Hash md5:", hValue)
		hValue = h.FromString(fPath).ByMd5().ToHexString()
		fmt.Println("字符", fPath, "Hash md5:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	hmacKey := "password"
	sDiff = utils.ExecTimeDiff(func() {

		hValue := h.FromFile(fPath).ByHmacMd5(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac md5:", hValue)
		hValue = h.FromString(fPath).ByHmacMd5(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac md5:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := h.FromFile(fPath).ByHmacSha1(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha1:", hValue)
		hValue = h.FromString(fPath).ByHmacSha1(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha1:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := h.FromFile(fPath).ByHmacSha224(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha224:", hValue)
		hValue = h.FromString(fPath).ByHmacSha224(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha224:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := h.FromFile(fPath).ByHmacSha256(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha256:", hValue)
		hValue = h.FromString(fPath).ByHmacSha256(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha256:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := h.FromFile(fPath).ByHmacSha384(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha384:", hValue)
		hValue = h.FromString(fPath).ByHmacSha384(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha384:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := h.FromFile(fPath).ByHmacSha512(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha512:", hValue)
		hValue = h.FromString(fPath).ByHmacSha512(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha512:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := h.FromFile(fPath).ByHmacSha512224(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha512_224:", hValue)
		hValue = h.FromString(fPath).ByHmacSha512224(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha512_224:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := h.FromFile(fPath).ByHmacSha512256(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha512_256:", hValue)
		hValue = h.FromString(fPath).ByHmacSha512256(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha512_256:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)
}
