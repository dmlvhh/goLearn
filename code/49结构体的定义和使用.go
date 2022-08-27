package main

import "fmt"

/**
在函数外部定义结构体 作用域是全局的
type 结构体名 struct{
	结构体成员
}
*/
type Student struct {
	id      int
	name    string
	sex     string
	age     int
	address string
}

func main() {
	//通过结构体名 定义结构体变量
	//var s Student
	//s.id=101
	//s.name="张飞"
	//s.sex="男"
	//s.age=28
	//s.address="北京燕郊"
	//fmt.Println(s)

	//var s Student=Student{id:101,name: "关羽",sex: "男",age: 29,address:"山西运城"}
	s := Student{id: 101, name: "关羽", sex: "男", age: 29, address: "山西运城"}
	fmt.Println(s.id)
	fmt.Println(s.name)
	fmt.Println(s.sex)
	fmt.Println(s.age)
	fmt.Println(s.address)

}
