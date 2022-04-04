package main

import "fmt"

func main() {
	s := []int{1, 2, 3} //len = 3, cap =3

	s1 := s[0:2]
	fmt.Printf("%v\n", s1)

	//s与s1指向同一个地址
	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)
	//与python相像
	//copy 复制切片的内容
	s2 := make([]int, 2) //s2 = [0,0,0]
	//将s的值复制到s2
	copy(s2, s)
	s2[0] = 2
	fmt.Println(s)
	fmt.Println(s2)
}
