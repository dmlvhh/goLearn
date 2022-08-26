package main

import "fmt"

func main1()  {
	// 输出格式 println 打印数据时自带换行
	/*
	fmt.Println("hello")
	fmt.Println(10)
	fmt.Println(3.14)
	*/

	//输出格式 print 打印数据时不带换行
	/*
	fmt.Print("hello")
	fmt.Print(10)
	fmt.Print(3.14)
	*/
	a:=10
	b:=3.1415926
	/*
	%d 占位符 表示输出一个整型数据
	%f 占位符 表示输出一个浮点型数据
	\n 占位符 表示转义字符 相当于换行符
	*/
	fmt.Printf("==%05d==\n",a)
	// %f 占位符默认保留6位小数 %.3f表示小数点后面保留3位小数 会对第4位四舍五入
	fmt.Printf("%.3f\n",b)
}
func main()  {
	//bool布尔 string字符串 byte字符
	//声明bool类型变量 默认输出值为 false
	var a bool
	a=true
	//fmt.Println(a)
	//%t 占位符 表示输出一个布尔类型数据
	fmt.Printf("%t",a)
	var b string
	b="你好呀golang"
	//%s 占位符 表示输出一个字符串类型数据
	fmt.Printf("%s\n",b)
	var c byte
	c='a'
	//字符型变量 对应的ASCII码值
	//fmt.Println(c)
	//%c 占位符 表示输出一个字符型数据
	fmt.Printf("%c",c)
}