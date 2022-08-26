package main

import "fmt"

func main01() {
	var s []int
	/**
	追加操作
	 */
	//在使用append添加数据时 切片的地址可能或发生变化 如果容量扩充导致输出存储溢出 切片会自动寻找新的空间存储数据
	s=append(s,1,2,3,4,5)
	fmt.Println(len(s))//0
	fmt.Println(cap(s))//0
	fmt.Printf("%p\n",s)

	s=append(s,1,2,3,4,5)
	fmt.Println(len(s))//0
	fmt.Println(cap(s))//0
	fmt.Printf("%p\n",s)
}

func main02()  {
	s1:=[]int{1,2,3,4,5}
	s2:=make([]int,5)
	//将s1切片中的数据拷贝到s2中  s2中要有足够的容量
	//使用拷贝操作后s1 s2是两个独立空间 不会相互影响
	copy(s2,s1)
	s2[2]=123
	fmt.Println(s2)
	fmt.Println(s1)

	fmt.Printf("%p\n",s1)
	fmt.Printf("%p\n",s2)
}

func main()  {
	s1:=[]int{1,2,3,4,5}
	s2:=make([]int,6)
	//如果想拷贝具体的切片中的片段 需要使用切片中的截取
	copy(s2,s1[1:])
	fmt.Println(s2)
}