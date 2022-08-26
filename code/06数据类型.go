package main

import "fmt"

func main1()  {
	// bool类型
	var a bool
	fmt.Println(a)
	a=true
	fmt.Println(a)
	b:=true
	fmt.Println(b)
}
func main2()  {
	// float32默认保留7位有效数据
	var a float32
	// float32默认保留15位有效数据
	var b float64
	a=3.141592611
	b=3.14159261111111455
	fmt.Println(a)
	fmt.Println(b)
}

func main3()  {
	var a byte = '0'
	var b byte = 'a'
	var c byte = 'A'
	// 打印字符串对应的ASCII值
	// '0' 48 'a' 97 'A' 65
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
func main()  {
	//"\0" 表示字符串的结束标志 "\n"表示换行
	//var a string
	//fmt.Println("==",a,"==")

	var a string="你好"
	var b string="golong"
	//c:=a+b
	//fmt.Println(c)
	//==运算符比较2值是否相同
	fmt.Println(a==b)
}