package main

import "fmt"
//函数在调用结束后会从内存中销毁
func testl(a int)  {
	a++
	fmt.Println(a)
}
func main1() {
	a:=0
	for i:=0;i<10;i++{
		testl(a)
	}
}
//----------------------------------------
//可以通过匿名函数和闭包 实现函数在栈区的本地化
func testa() func() int  {
	var a int
	return func() int {
		a++
		return a
	}
}
func main()  {
	//将testa函数类型赋值给f
	//f:=testa
	//函数调用 将testa的返回值给f
	//f:=testa()
	var a int
	f:= func() int{
		a++
		return a
	}
	for i:=0;i<10;i++ {
		fmt.Println(f())
	}
	//fmt.Printf("%T",f)
}
