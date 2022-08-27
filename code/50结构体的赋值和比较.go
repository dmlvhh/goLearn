package main

import "fmt"

type Students struct {
	id      int
	name    string
	sex     string
	age     int
	address string
}

func main() {
	s := Students{id: 101, name: "貂蝉", sex: "女", age: 18, address: "徐州"}
	s1 := s
	s1.age = 22
	fmt.Println(s1)
	fmt.Println(s)

	//在结构体中使用==  ！=可以对结构体成员进行比较操作
	if s1 == s {
		fmt.Println("相同")
	} else {
		fmt.Println("不相同")
	}
	//大于 小于可以比较结构体成员
	if s.age > s1.age {
		fmt.Println("真")
	} else {
		fmt.Println("假")
	}

}
