package main

import "fmt"

func main1()  {
	//逻辑非 ！ 非真为假 非假为真
	var a bool = false
	a=true
	b:=20
	c:=20
	fmt.Println(!(b==c))
	//fmt.Println(!b)//error
	fmt.Println(!a)
}
func main2()  {
	//逻辑与 &&  表达式1&&表达式2 同真为真 其余为假
	a:=10
	b:=20
	c:=30
	d:=20
	fmt.Println(a>b && c<d)
}

func main3()  {
	//逻辑或||   表达式1 || 表达式2  同假为假 其余为真
	a:=10
	b:=20
	c:=30
	d:=20
	fmt.Println(a<b || c<d)

}
func main()  {
	//取地址运算符
	//var a int = 10
	//fmt.Println(&a)
	a:=10
	//指针变量 *取值运算符
	p:=&a
	fmt.Println(*p)
}
