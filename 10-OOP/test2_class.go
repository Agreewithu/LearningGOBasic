package main

import "fmt"

//go语言中的class就是通过结构体绑定方法
type Hero struct {
	name  string
	ad    int
	Level int
}

//方法本身也是函数,类中的函数被称为方法
func (this *Hero) SetName(newName string) {
	(*this).name = newName
}

func (this *Hero) GetName() string {
	return (*this).name
}

func (this *Hero) Show() {
	fmt.Println("hero = ", this)
}

func main() {
	//创建一个对象
	hero := Hero{name: "rx", ad: 100, Level: 1}
	fmt.Println(hero.GetName())
	hero.Show()
	hero.SetName("lh") //this 不是什么其他的东西，不过是一个变量名而已
	hero.Show()
}

// 如果说类的首字母大写，其他包也能访问
// 方法首字母大写，其他包也能访问
