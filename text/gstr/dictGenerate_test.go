package gstr

import (
	"context"
	"fmt"
	"testing"
)

func TestGenerateDictDo(t *testing.T) {
	i := 1
	GenerateDictDo("123a", 2, func(s string, cancelFunc context.CancelFunc, count int) {
		if i == 1 {
			fmt.Println("总生成数:", count)
		}
		fmt.Println(s)
		i++
		if i > 5 {
			cancelFunc()
		}
	})
}
