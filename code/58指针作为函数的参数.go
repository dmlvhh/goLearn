package main

import "fmt"

//func swap(a int,b int)  {
//	a,b=b,a
//}
func swap(a *int,b *int)  {
	*a,*b=*b,*a
}
func main()  {
	a:=10
	b:=20
	//不能交换a b的值 值传递
	//swap(a,b)
	//指针作为函数参数是地址传递
	swap(&a,&b)
	fmt.Println(a)
	fmt.Println(b)
}