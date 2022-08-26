package main

import "fmt"

func test(s []int)  {
	s[2]=123
}
func main01() {
	s:=[]int{1,2,3,4,5}
	test(s)
	//fmt.Printf("%p\n",s)
	fmt.Println(s)
}
func BubbleSort(s []int)  {

	for i:=0;i<len(s)-1;i++ {
		for j:=0;j<len(s)-1-i;j++{
			if s[j]>s[j+1] {
				s[j],s[j+1]=s[j+1],s[j]
			}
		}
	}
}
func main02() {
	s:=[]int{9,1,5,6,3,10,2,4,8}
	//切片作为函数参数是地址传递 形参可以改变实参
	BubbleSort(s)
	fmt.Println(s)
}
func test1(s []int) []int {
	//如果使用append操作 可能会改变地址 需要用返回值
	s=append(s,456)
	fmt.Printf("%p\n",s)
	fmt.Println(s)
	return s
}
func main()  {
	s:=[]int{1,2,3}
	fmt.Printf("%p\n",s)
	s=test1(s)
	fmt.Println(s)
	fmt.Printf("%p\n",s)
}