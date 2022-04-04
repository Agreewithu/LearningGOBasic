package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5) //len = 3,cap = 5
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	/*
		[][][][][]
		|    \ptr 后面的是不可访问的，但是已经创建好了
		numbers
	*/
	//向numbers切片追加一个元素1
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	//再添加一个元素2
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	//cap的真正含义,头部指针和尾部指针的位置差值
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	//相当于java中的ArrayList
}
