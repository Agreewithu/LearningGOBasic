package main

import "fmt"

func main() {
	var myArray1 [10]int
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}
	//另一种方式定义和打印
	myArray2 := [10]int{1, 2, 3, 4}
	for index, value := range myArray2 {
		fmt.Printf("index = %d, vaule = %d\n", index, value)
	}
	//查看数组的数据类型
	fmt.Printf("type of myArray1 is %T\n", myArray1)
	fmt.Printf("type of myArray2 is %T\n", myArray2)
}
