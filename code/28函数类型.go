package main

import "fmt"

func test6()  {
	fmt.Println("瓜娃子")
}
func test7(a int,b int)  {
	fmt.Println(a+b)
}
func test8()  {
	fmt.Println("你好golang")
}
func test9(a int,b int) int {
	return a+b
}

//type 可以定义函数类型
//type 可以为已存在类型起别名
//函数类型其实就是指针
type FUNCTYPE func()
type FUNCTEST func(int, int)
type funcDemo func(int, int) int
func main1() {
	//定义函数类型变量
	var f FUNCTYPE
	f=test6
	//通过f()函数变量调用函数
	f()
	f=test8
	f()

	var f1 FUNCTEST
	f1=test7
	f1(10,20)

	var f2  funcDemo
	f2=test9
	v:=f2(20,30)
	fmt.Println(v)
}
//-------------------------------------------
func test10(a int,b int){
	fmt.Println(a+b)
}

func main() {
	//函数调用
	//test10(10,20)
	//如果使用print打印函数名是一个地址
	//函数名本身就是一个指针类型数据 在内存中代码区进行存储
	//fmt.Println(test10)
	//使用自动推到类型
	//f:=test10
	//f(10,20)
	//fmt.Printf("%T",f)
	//直接定义函数类型
	var f func(int,int)
	f=test10
	f(10,20)
}