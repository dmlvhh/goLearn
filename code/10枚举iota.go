package main

import "fmt"

func main1()  {

	const(
		a=iota
		b=iota
		c=iota
		d=iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	//定义变量 为状态
	value:=a
	fmt.Println(value)
	value=b
	fmt.Println(value)


}
func main()  {
	//如果定义枚举是常量写在同一行值相同
	//const(
	//	a=iota //0
	//	b //1
	//	c //2
	//	d //3
	//	e //4
	//)
	const(
		a=10
		b,c=iota,iota
		d,e
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}