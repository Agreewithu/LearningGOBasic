package main

import "fmt"

//本质是一个指针
type AnimialIF interface { //animial infomation
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

//具体的类
type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("cat is sleeping...")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

//go是通过这种方式去实现接口的~
//这样AnimialIF的指针才能指向实现该接口的类了

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog is sleeping...")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

//通过接口去触发多态的现象
func main() {
	var animial AnimialIF //接口类型，父类指针
	animial = &Cat{color: "green"}
	animial.Sleep() //调用的是cat的sleep方法
	animial = &Dog{color: "yellow"}
	animial.Sleep()
}
