package encoding

// 该部分代码参考于 https://github.com/golang-module/dongle

type BaseStruct struct {
	src []byte
	dst []byte
	Err error
}
