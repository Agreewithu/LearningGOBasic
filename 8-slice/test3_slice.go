package main

import "fmt"

func main() {
	//声明slice是一个切片，并且初始化
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slcie = %v\n", len(slice1), slice1) //%v表示打印出任何类型的详细信息
	//声明slice 是一个切片，但是没有给slice分配空间
	var slice2 []int
	//给slice分配空间
	slice2 = make([]int, 3) //通过make来分配空间
	slice2[0] = 100
	fmt.Printf("len = %d, slcie = %v\n", len(slice2), slice2) //%v表示打印出任何类型的详细信息
	//声明slice是一个切片，同时给slice分配空间
	var slice3 []int = make([]int, 3) //slice := make([]int, 3)
	fmt.Printf("len = %d, slcie = %v\n", len(slice3), slice3)

}
