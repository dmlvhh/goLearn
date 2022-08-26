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

func main03()  {
	//不写元素个数叫切片 必须写元素个数的叫数组
	var s []int = []int{1,2,3,4,5}
	//fmt.Println(s)
	//容量大于等于长度
	s=append(s,6,7,8,9)
	fmt.Println("长度：",len(s))
	fmt.Println("容量：",cap(s))

	s=append(s,6,7,8,9)
	fmt.Println("长度：",len(s))
	fmt.Println("容量：",cap(s))

	s=append(s,6,7,8,9)
	fmt.Println("长度：",len(s))
	fmt.Println("容量：",cap(s))
	//如果整体数据没有超过1024字节 则每次扩展容量为上一次的倍数 超过1024 每次扩展是上一次的1/4
	s=append(s,6,7,8,9)
	fmt.Println("长度：",len(s))
	fmt.Println("容量：",cap(s))
}

func main()  {
	s:=make([]int,5)
	s[0]=1
	//使用append在长度后添加数据
	s=append(s,1,2,3)
	//fmt.Println(s)
	//fmt.Println(cap(s))
	//在切片打印时 只能打印有效长度中的数据 cap不能为打印的条件
	for i:=0;i<len(s);i++ {
		fmt.Println(s[i])
	}
}
