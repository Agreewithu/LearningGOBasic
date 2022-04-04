package main

import "fmt"

func fibonaci(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			//如果c可写
			temp := x
			x = y
			y = temp + x
		case <-quit:
			//如果quit可读的话，那么循环结束了
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	//同时监控两个chan的状态
	fibonaci(c, quit)

}
