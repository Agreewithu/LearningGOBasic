package main

import "fmt"

func main() {
	cityMap := make(map[string]string)
	//add
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	//遍历
	for key, value := range cityMap {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}
	//delete
	delete(cityMap, "China")
	//调用函数
	search(cityMap)
}

func search(objMap map[string]string) {
	for key, value := range objMap {
		fmt.Printf("key = %s, value = %s\n", key, value)
	}
}
