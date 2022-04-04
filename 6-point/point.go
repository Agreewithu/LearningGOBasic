package main

import "fmt"

func main() {
	a := 10
	b := 20
	//swap
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d", a, b)
}

func swap(a *int, b *int) {
	temp := *a // 把a的内容赋给temp
	*a = *b
	*b = temp
}
