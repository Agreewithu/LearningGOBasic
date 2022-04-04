package main

import (
	//要做成有关GOPATH的相对路径
	"../5-init/lib1"

	"../5-init/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
