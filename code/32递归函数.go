package main

import "fmt"

//在函数定义时调用函数本身 递归函数
func test03(a int){
	//在递归函数中一个要有出口 return
	if a==0 {
		return
	}
	a--
	fmt.Println(a)
	test03(a)
}
func main1() {
	test03(10)
}
//-------------------------------------
//计算n的阶乘
var sum int =1
func test04(n int)  {
	if n==1 {
		return
	}
	test04(n-1)
	sum*=n
}
func main()  {
	test04(5)
	fmt.Println(sum)
}
