package main

import "fmt"

func main() {
	c := make(chan int) //开一个无缓冲的channel
	//内部匿名函数
	go func() {
		for i := 0; i < 5; i++ {
			//把循环因子i->c
			c <- i
		}

		//close可以关闭一个channel
		close(c)
	}()

	for {
		//ok为true imply c is open, ok 为false imply c is close
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Finished....")
}
