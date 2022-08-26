package main

import "fmt"

func main1() {
	a:=10 //int
	b:=20 //int
	c:=a/b
	//整数数据相除结果为整型
	//除数不能为0
	fmt.Println(c)
}

//取余运算符
func main2(){
	a:=10
	b:=20
	//取模运算符 取余运算符只能用于整型数据 取余运算符除数不能为0
	c:=a%b
	fmt.Println(c)
}
//自增后自减运算符
func main()  {
	a:=10
	//++ --只能写在变量的后面 叫做后自增后自减
	a++//自增运算符
	//a--//自减运算符
	//a:=a--  不能将自增运算符运用在表达式中
	fmt.Println(a)
}

