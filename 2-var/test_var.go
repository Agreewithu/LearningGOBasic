/*
	四种声明变量的方式
*/
package main

import "fmt"

//声明全局变量
var Ga int = 10
var (
	vv int  = 100
	jj bool = false
)
var xx, yy int = 100, 200

// := 的方法不能用于声明全局变量

func main() {
	//方法1：声明一个变量
	var a int
	fmt.Println("a = ", a)
	//方法2:声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	//方法3：自动匹配类型
	var c = 100
	fmt.Println("c = ", c)
	//printf代表格式化输出
	//println代表非格式化
	fmt.Printf("type of c = %T\n", c)
	//测试一下
	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb is %T\n", bb, bb)
	//方法4，最常用的方法，省去var
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("e = %d , type of e is %T\n", e, e)
	f := 3.14
	fmt.Println("f = ", f)
	fmt.Printf("f = %f , type of f is %T\n", f, f)
	//测试全局变量
	fmt.Println("Ga = ", Ga)
	//同时声明多个变量
	// var xx, yy int = 100, 200
	fmt.Printf("xx = %d yy = %d\n", xx, yy)
	var kk, ll = 15, "aaa"
	fmt.Printf("xx = %d yy = %s\n", kk, ll)
	//多行多变量声明
	// var (
	// 	vv int  = 100
	// 	jj bool = false
	// )

	fmt.Println("vv = ", vv)
	fmt.Println("jj = ", jj)

}
