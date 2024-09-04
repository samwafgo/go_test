package main

import (
	"fmt"
	"go_test/global"
)

func main() {
	var tip = "hello this is go build test"
	fmt.Println(tip)
	fmt.Println("GWAF_RELEASE:" + global.GWAF_RELEASE)
	fmt.Println("GWAF_RELEASE_VERSION_NAME:" + global.GWAF_RELEASE_VERSION_NAME)
	fmt.Println("GWAF_RELEASE_VERSION:" + global.GWAF_RELEASE_VERSION)
}
