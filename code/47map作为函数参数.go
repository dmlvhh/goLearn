package main

import "fmt"

func dem0(m map[int]string)  {
	m[101]="孙悟空2"
	m[103]="唐老三"
	fmt.Println(len(m))
	delete(m,101)
	fmt.Println(len(m))
}
func main()  {
	m:=make(map[int]string,1)
	m[101]="孙悟空"
	fmt.Println(len(m))
	//map作为函数参数是地址传递(引用传递)
	dem0(m)
	fmt.Println(m)
}