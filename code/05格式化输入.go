package main

import "fmt"

func main1()  {
	//var a int
	// & 运算符 取地址运算符
	// 可以用键盘输入
	//fmt.Scan(&a)
	//fmt.Println(a)
	//%p 占位符 表示输出一个数据对应的内存地址 &a
	//0x表示十六进制数据
	//fmt.Printf("%p",&a)

	// 空格或者回车作为接收结束
	//var str string
	//fmt.Scan(&str)
	//fmt.Println(str)

	// 接收2个数据
	var s1,s2 string
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	fmt.Println(s1)
	fmt.Println(s2)
}

func main2()  {
	var r float64
	var PI float64

	//通过键盘获取半径
	fmt.Scan(&r,&PI)
	fmt.Printf("面积为：%.2f\n",PI*r*r)
	fmt.Printf("周长为：%.2f\n",PI*2*r)
}
func main3()  {
	var a int
	var b string
	fmt.Scanf("%3d",&a)
	fmt.Scanf("%s",&b)
	fmt.Println(a)
	fmt.Println(b)
}
func main()  {
	// 命名变量规则
	var name string
	var passwd string
	fmt.Println("请输入用户名")
	fmt.Scanf("%s",&name)
	fmt.Println("请输入密码")
	fmt.Scanf("%s",&passwd)
	fmt.Println(name,passwd)
}