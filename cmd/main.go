package main

import (
	"context"
	"fmt"
	"github.com/chinaDL/whTools"
	"github.com/chinaDL/whTools/text/gstr"
)

func main() {

	gstr.GenerateDictDo("01234", 10, func(s string, cancelFunc context.CancelFunc, count int) bool {
		fmt.Println(count)
		return false
	})

	fmt.Println(whTools.Decode.FromString("AAIGc3RyaW5nDAgABnVzZXJJZAR1aW50BgIAHgZzdHJpbmcMCgAIdXNlcm5hbWUGc3RyaW5nDA0ACzEzODg4ODg4ODg4").ByBase64().ToBytes())

}
