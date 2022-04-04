package main

import "fmt"

//go语言怎么定义函数呢

func fool(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

//多返回值的函数
func fool2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 666, 777
}

//返回多个返回值，有形参名的
func fool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("======fool3======")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

//写一个函数用于打印
func printResult(ret1 int, ret2 int) {
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)
}

func main() {
	c := fool("abc", 555)
	fmt.Println("c = ", c)
	//测试多返回值的函数
	ret1, ret2 := fool2("hahah", 999)
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)
	ret1, ret2 = fool3("fool3", 333)
	printResult(ret1, ret2)
}
