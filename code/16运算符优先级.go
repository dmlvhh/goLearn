package main

import "fmt"

//运算顺序
/*
括号() 结构体成员  数组下标[]
单目运算符 逻辑非！ 取地址& 取值* 自增++ 自减--
双目运算符 乘除 加减 关系 逻辑 || &&  赋值 = += -= *= /= %=

*/
func main1()  {
	a:=10
	b:=20
	c:=30
	//d:=a+b*c
	//var d int
	//d=(a+b)*c

	//fmt.Println(d)
	fmt.Println(a+b>=c && !(b>c))
}

func main()  {
	fmt.Println("请输入年份")
	var year int
	fmt.Scan(&year)
	//1.能被400整除  2.能够被4整除 不能被100整除
	//逻辑与优先级高于逻辑或
	b:=(year%400)==0 || (year%4==0 && year%100!=0)
	fmt.Println(b)


}
