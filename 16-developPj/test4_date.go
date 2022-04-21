package main

import (
	"fmt"
	"time"
)

func main() {
	// currentTime := time.Now()
	// fmt.Println(currentTime.String())
	timeModelStr := "2006-01-02"
	s1 := "2022-04-06"
	t1, _ := time.ParseInLocation(timeModelStr, s1, time.Local)
	// fmt.Println(t1)
	s2 := "2022-04-30"
	t2, _ := time.ParseInLocation(timeModelStr, s2, time.Local)

	sub := int(t2.Sub(t1).Hours())
	fmt.Printf("%T\n", sub)
	daySub := (sub / 24) % 21
	fmt.Printf("%T\n", daySub)
	fmt.Println("daySub = ", daySub)
}
