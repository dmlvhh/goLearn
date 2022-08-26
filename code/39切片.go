package main

import "fmt"

func main01(){
	//数组定义 var 数组名 [元素个数]数据类型
	//切片定义 var 切片名 []数据类型

	//var s []int
	//fmt.Println(s)

	s:=make([]int,5)
	//通过下标为切片赋值
	s[0]=123
	s[1]=234
	s[2]=345
	s[3]=456
	s[4]=567
	//s[6]=567 err
	//通过append 追加切片的信息
	s=append(s,678,789,8910)
	fmt.Println(s)
	//通过len查看切片长度
	fmt.Println(len(s))
}
func main02()  {
	s:=make([]int,5)
	s[0]=123
	s[1]=234
	s[2]=345
	s[3]=456
	s[4]=567
	//遍历
	//for i:=0;i<len(s);i++ {
	//	fmt.Println(s[i])
	//}

	for i,v:=range s{
		fmt.Println(i,v)
	}
	fmt.Println(s)
}

func main()  {
	//不写元素个数叫切片 必须写元素个数的叫数组
	var s []int = []int{1,2,3,4,5}
	fmt.Println(s)
}
