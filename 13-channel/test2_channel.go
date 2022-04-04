package main

import "fmt"

func main() {
	c := make(chan int, 3) //带有缓冲的channel
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("子go程正在运行", ", 发送的元素i = ", i, ", len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	//主go程去把channel中的东西拿出来
	for i := 0; i < 3; i++ {
		num := <-c // 从c中接收数据
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束")
}
