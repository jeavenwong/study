package main

import (
	"fmt"
	"main/modules"
	"main/test"
)

func main() {
	str := "hello, world!"
	fmt.Println(str)
	test.TestStr(str)
	modules.TestA(str)
}
