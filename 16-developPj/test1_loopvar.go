package main

import "fmt"

func main() {
	//测试一下循环新建是否会报错
	for i := 0; i < 5; i++ {
		j := i
		fmt.Println("j = ", j)
	}
}

//测试完毕，循环新建不会报错
