package utils

import (
	"time"
)

// ExecTimeDiff 返回函数执行用时(s)
func ExecTimeDiff(fn func()) float64 {
	startTime := time.Now()

	fn()
	return (time.Now().Sub(startTime)).Seconds()

}
