package main

import "fmt"

func main()  {
	m:=make(map[int]string,1)
	m[0]="刘备"
	m[1]="张飞"
	m[2]="张飞"

	//fmt.Println(m[2])
	////在map打印中如果没有出现定义的map 默认为空
	//fmt.Println("----",m[3],"----")

	//v,ok:=m[3]
	//if ok {
	//	fmt.Println(v)
	//}else{
	//	fmt.Println("可以不存在")
	//}

	for k,v:=range m {
		//v,ok:=m[k]
		fmt.Println(k,v)
	}

}