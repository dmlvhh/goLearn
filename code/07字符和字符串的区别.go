package main

import "fmt"

func main1()  {
	// 存储值不同
	//var a byte ='a'
	//var b string ="a" //'a' '\0'
	//fmt.Println(a==b)
	//%s 遇到\0停止
	var b string = "hello\nworld"
	fmt.Printf("%s",b)

}

func main()  {
	//var str string = "hello world"
	//在go语言中一个汉子算作3个字符 为了和linux统一处理
	var str string = "传智播客"
	//计算字符串个数
	num:=len(str)
	fmt.Println(num)
}