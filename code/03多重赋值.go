package main

import "fmt"

func main1()  {
	// 格式 变量1,变量2,变量3:=值1,值2,值3
	a,b,c,d,e:=10,20,30,40,50
	fmt.Println(a,b,c,d,e)
}
func main2() {
	a,b:=10,20

	// 交换a b 的值
	// 通过自动推到类型为temp赋值
	// temp:=a
	// a=b
	// b=temp

	// 加减计算
	//a=a+b
	//b=a-b
	//a=a-b

	//交换变量的值
	a,b=b,a
	fmt.Println(a)
	fmt.Println(b)
}
