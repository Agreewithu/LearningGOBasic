package lib1

import "fmt"

func init() {
	fmt.Println("lib1. init() ...")
}

//定义两个方法，尝试使用
//lib1包当前提供的api

func Lib1Test() {
	fmt.Println("Lib1Test()...")
}
