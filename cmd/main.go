package main

import (
	"fmt"
	"github.com/chinaDL/whTools"
	"github.com/chinaDL/whTools/utils"
)

func main() {
	en := "abc123"
	de := whTools.Encode.FromString(en).ByBase32().ToString()
	fmt.Println("base32 encode:", de)
	fmt.Println("base32 decode:", whTools.Decode.FromString(de).ByBase32().ToString())

	de = whTools.Encode.FromString(en).ByBase45().ToString()
	fmt.Println("base45 encode:", de)
	fmt.Println("base45 decode:", whTools.Decode.FromString(de).ByBase45().ToString())

	de = whTools.Encode.FromString(en).ByBase58().ToString()
	fmt.Println("base58 encode:", de)
	fmt.Println("base58 decode:", whTools.Decode.FromString(de).ByBase58().ToString())

	de = whTools.Encode.FromString(en).ByBase64().ToString()
	fmt.Println("base64 encode:", de)
	fmt.Println("base64 decode:", whTools.Decode.FromString(de).ByBase64().ToString())

	de = whTools.Encode.FromString(en).ByBase64URL().ToString()
	fmt.Println("base64Url encode:", de)
	fmt.Println("base64Url decode:", whTools.Decode.FromString(de).ByBase64URL().ToString())

	fPath := "/Users/whaz/Downloads/rustdesk-1.1.9.dmg"
	sDiff := utils.ExecTimeDiff(func() {
		hValue := whTools.Hash.FromFile(fPath).BySha1().ToHexString()
		fmt.Println("文件", fPath, "Hash sha1:", hValue)
		hValue = whTools.Hash.FromString(fPath).BySha1().ToHexString()
		fmt.Println("字符", fPath, "Hash sha1:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := whTools.Hash.FromFile(fPath).BySha224().ToHexString()
		fmt.Println("文件", fPath, "Hash sha224:", hValue)
		hValue = whTools.Hash.FromString(fPath).BySha224().ToHexString()
		fmt.Println("字符", fPath, "Hash sha224:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := whTools.Hash.FromFile(fPath).BySha256().ToHexString()
		fmt.Println("文件", fPath, "Hash sha256:", hValue)
		hValue = whTools.Hash.FromString(fPath).BySha256().ToHexString()
		fmt.Println("字符", fPath, "Hash sha256:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := whTools.Hash.FromFile(fPath).BySha384().ToHexString()
		fmt.Println("文件", fPath, "Hash sha384:", hValue)
		hValue = whTools.Hash.FromString(fPath).BySha384().ToHexString()
		fmt.Println("字符", fPath, "Hash sha384:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := whTools.Hash.FromFile(fPath).BySha512().ToHexString()
		fmt.Println("文件", fPath, "Hash sha512:", hValue)
		hValue = whTools.Hash.FromString(fPath).BySha512().ToHexString()
		fmt.Println("字符", fPath, "Hash sha512:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := whTools.Hash.FromFile(fPath).BySha512256().ToHexString()
		fmt.Println("文件", fPath, "Hash sha512/256:", hValue)
		hValue = whTools.Hash.FromString(fPath).BySha512256().ToHexString()
		fmt.Println("字符", fPath, "Hash sha512/256:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := whTools.Hash.FromFile(fPath).BySha512224().ToHexString()
		fmt.Println("文件", fPath, "Hash sha512/224:", hValue)
		hValue = whTools.Hash.FromString(fPath).BySha512224().ToHexString()
		fmt.Println("字符", fPath, "Hash sha512/224:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {
		hValue := whTools.Hash.FromFile(fPath).ByMd5().ToHexString()
		fmt.Println("文件", fPath, "Hash md5:", hValue)
		hValue = whTools.Hash.FromString(fPath).ByMd5().ToHexString()
		fmt.Println("字符", fPath, "Hash md5:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	hmacKey := "password"
	sDiff = utils.ExecTimeDiff(func() {

		hValue := whTools.Hash.FromFile(fPath).ByHmacMd5(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac md5:", hValue)
		hValue = whTools.Hash.FromString(fPath).ByHmacMd5(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac md5:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := whTools.Hash.FromFile(fPath).ByHmacSha1(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha1:", hValue)
		hValue = whTools.Hash.FromString(fPath).ByHmacSha1(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha1:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := whTools.Hash.FromFile(fPath).ByHmacSha224(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha224:", hValue)
		hValue = whTools.Hash.FromString(fPath).ByHmacSha224(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha224:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := whTools.Hash.FromFile(fPath).ByHmacSha256(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha256:", hValue)
		hValue = whTools.Hash.FromString(fPath).ByHmacSha256(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha256:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := whTools.Hash.FromFile(fPath).ByHmacSha384(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha384:", hValue)
		hValue = whTools.Hash.FromString(fPath).ByHmacSha384(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha384:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := whTools.Hash.FromFile(fPath).ByHmacSha512(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha512:", hValue)
		hValue = whTools.Hash.FromString(fPath).ByHmacSha512(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha512:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := whTools.Hash.FromFile(fPath).ByHmacSha512224(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha512_224:", hValue)
		hValue = whTools.Hash.FromString(fPath).ByHmacSha512224(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha512_224:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)

	sDiff = utils.ExecTimeDiff(func() {

		hValue := whTools.Hash.FromFile(fPath).ByHmacSha512256(hmacKey).ToHexString()
		fmt.Println("文件", fPath, "Hash hmac sha512_256:", hValue)
		hValue = whTools.Hash.FromString(fPath).ByHmacSha512256(hmacKey).ToHexString()
		fmt.Println("字符", fPath, "Hash hmac sha512_256:", hValue)
	})
	fmt.Printf("耗时: %f\n", sDiff)
}
