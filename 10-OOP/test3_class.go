package main

import "fmt"

type Human struct {
	name string
	sex  string
}

//写Human的方法
func (this *Human) Eat() {
	fmt.Println("someone is eating....")
}

func (this *Human) Walk() {
	fmt.Println("someone is walking....")
}

//superMan类继承Human
type SuperMan struct {
	Human //SuperMan 继承了Human类的方法
	level int
}

//重新定义父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("superMan is eating ....")
}

//写一个新的方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan is flying...")
}

func main() {
	h := Human{name: "rx", sex: "male"}
	h.Eat()
	h.Walk()
	//定义一个子类对象
	s := SuperMan{Human{"rx", "male"}, 5}
	s.Walk()
	s.Eat()
	s.Fly()
}
