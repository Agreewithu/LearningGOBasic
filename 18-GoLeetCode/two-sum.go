package main

import "fmt"

func two_sum(nums []int, target int) interface{} {
	result := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				return result
			}
		}
	}
	return nil
}

//onters' better solution
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, _ := range nums {
		diff := target - nums[i]
		if j, ok := m[diff]; ok {
			return []int{i, j}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	res := two_sum(nums, 9)
	fmt.Printf("res is %T\n", res)
	fmt.Printf("res = %v\n", res)
	fmt.Printf("=================\n")
	for _, value := range nums {
		// fmt.Printf("index = %d\t", index)
		fmt.Printf("value = %d\n", value)
	}
	fmt.Printf("=================\n")
	m := make(map[int]int)
	j, ok := m[7]
	fmt.Printf("j = %v\n", j)
	fmt.Printf("j is %T\n", j)
	fmt.Printf("ok = %v\n", ok)
	fmt.Printf("ok is %T\n", ok)
	fmt.Printf("m = %v\n", m)
	fmt.Printf("\nOther show his better solution: another_try = %v\n", twoSum(nums, 9))
}
