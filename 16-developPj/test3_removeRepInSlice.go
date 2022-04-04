package main

import "fmt"

func removeRepByMap(slc []int) []int {
	result := []int{}        //存放返回的不重复切片
	tempMap := map[int]int{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

func main() {
	s := []int{1, 2, 3, 3, 5, 3, 9}
	fmt.Printf("before remove repcite element ---> %v\n", s)
	s = removeRepByMap(s)
	fmt.Printf("after remove repcite element ---> %v\n", s)
}
