package main

import (
	"fmt"
	"github.com/chinaDL/whTools"
)

func main() {

	fmt.Println(whTools.Decode.FromString("AAIGc3RyaW5nDAgABnVzZXJJZAR1aW50BgIAHgZzdHJpbmcMCgAIdXNlcm5hbWUGc3RyaW5nDA0ACzEzODg4ODg4ODg4").ByBase64().ToBytes())

}
