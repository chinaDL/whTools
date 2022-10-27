package utils

import (
	"strconv"
	"unsafe"
)

// IsNumeric 返回字符串是否是数字
func IsNumeric(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}

// Bytes2String 将字节切片转换为字符串
func Bytes2String(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// Str2bytes 将字符串转换为字节切片
func Str2bytes(s string) []byte {
	if len(s) == 0 {
		return nil
	}
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
