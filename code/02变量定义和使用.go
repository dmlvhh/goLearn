package main

import "fmt"

func main01() {
	/*
	变量的定义和使用,声明和定义
	var 变量名 数据类型
	布尔类型 整型 浮点型 字符型 字符串类型
	变量在定义时存储不同的数据 需要不同的数据类型
	*/
	var a int
	a=10
	a=a+30
	fmt.Println(a)
}
func main02() {
	/*计算圆的面积和周长
	*/
	var PI float64 = 3.1415926
	var r float64 = 2.5
	var s float64
	var c float64
	// 面积
	s = PI * r * r
	c = PI * 2 * r

	fmt.Println("面积为：",s)
	fmt.Println("周长为：",c)
}
func main03() {
	/*计算圆的面积和周长
	 */
	 PI := 3.1415926
	 r := 2.5
	// 面积
	s := PI * r * r
	c := PI * 2 * r
	fmt.Println("面积为：",s)
	fmt.Println("周长为：",c)
}
func main04() {
	// 自动推到类型 变量 := 值 会根据值为变量选择类型
	// 去市场买2斤黄瓜 价格为5元
	w := 2.0 //float64
	//var p float64 = 5
	//p : = 5 在go语言中不同的数据类型不能进行计算 可以通过类型转换解决
	p := 5.0 //int
	fmt.Println(w*p)

}
func main()  {
	a:=false //bool
	b:=10 //int
	c:=3.14 //float64
	d:='a' //byte
	e:="看你咋滴" //string
	fmt.Println(a,b,c,d,e)
}