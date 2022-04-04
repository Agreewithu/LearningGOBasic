package main

import "fmt"

//声明一种新的数据类型
type Book struct {
	title string
	auth  string
}

func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "rx"

	fmt.Printf("%v", book1)

	//change
	changeBook(&book1)

	fmt.Printf("%v", book1)
}

//go 语言结构的传参
/*
对于go而言,引用数据类型的有:slice、map、channel、指针、func
对于java这种面向对象的语言而言
基本数据类型：byte, short, int, long, float, double, char, boolean
*/
func changeBook(objbook *Book) {
	objbook.title = "changed_title"
	objbook.auth = "changed_auth"
}
