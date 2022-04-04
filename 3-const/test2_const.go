package main

import "fmt"

//const 来定义枚举类型
const (
	//可以在const()里面添加关键iota,每行的iota都会累加1,第一行的iota默认为0
	BEIJING = 10
	SHANGHAI
	HANGZHOU = iota
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f = iota * 2, iota * 3
	g, h
)

func main() {
	//常量（只读属性）
	const length int = 10
	fmt.Println("length = ", length)
	//改变一下常量的值
	// length = 1直接报错
	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("HANGZHOU = ", HANGZHOU)
	fmt.Println("SHENZHEN = ", SHENZHEN)
	//iota公式
	fmt.Println("a = ", a, " b = ", b)
	fmt.Println("c = ", c, " d = ", d)
	fmt.Println("e = ", e, " f = ", f)
	fmt.Println("g = ", g, " h = ", h)
	//iota只能在const中使用

}
