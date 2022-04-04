package main

import "fmt"

func main() {
	myArr := []int{1, 2, 3, 4} //动态数组，切片，slice,引用类型
	fmt.Printf("type of myArr is %T\n", myArr)
	fmt.Println("=========")
	for _, value := range myArr {
		fmt.Printf("value = %d\n", value)
	}
	fmt.Println("========")
	printArr(myArr)
	fmt.Println("========")
	for _, value := range myArr {
		fmt.Printf("value = %d\n", value)
	}
}

//写一个函数遍历切片
func printArr(arr []int) {
	//arr = 0x0011
	for index, value := range arr {
		fmt.Printf("index = %d, value = %d\n", index, value)
	}
	arr[0] = 100
}
