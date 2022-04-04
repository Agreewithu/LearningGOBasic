package main

import "fmt"

func main() {
	//定义一个channel
	c := make(chan int)

	// 其中一个线程是我们的主线程main
	// 新开一个线程
	go func() {
		defer fmt.Println("goroutine end ...")
		fmt.Println("goroutine is running...")
		c <- 666 ///将666发送给c
	}()

	num := <-c //从c中接收数据，并发送给num
	fmt.Println("num = ", num)
	fmt.Println("main goroutine end...")

}
