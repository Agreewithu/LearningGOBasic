package main

import "fmt"

type Knowledge struct {
	Id   int
	Name string
}

func main() {
	//测试一下指针类型的变量重新赋值
	knowledge1 := &Knowledge{
		Id:   1,
		Name: "goodtime",
	}
	knowledge2 := &Knowledge{
		Id:   2,
		Name: "badtime",
	}
	fmt.Printf("knowledge1 = %v, knowledge2 = %v\n", knowledge1, knowledge2)

	knowledge2 = knowledge1
	fmt.Printf("knowledge1 = %v, knowledge2 = %v\n", knowledge1, knowledge2)
	//指针类型的变量是可以重新赋值的

	//开始新的测试
	fmt.Println(knowledge1.getKnowledgeName())
	//所以是可以调用定义在外面的函数的

}

//绑定类的方法
func (s *Knowledge) getKnowledgeName() string {
	printWhatWant("It will return name")
	return s.Name
}

func (s *Knowledge) setKnowledgeId(resetId int) {
	s.Id = resetId
}

func printWhatWant(str string) {
	fmt.Println(str)
}
