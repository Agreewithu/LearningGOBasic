package main

import (
	"fmt"
	"time"
)

func main() {
	//用go创建承载一个形参为空，返回值为空的一个函数
	go func() {
		//还记得defer...是java中final
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
