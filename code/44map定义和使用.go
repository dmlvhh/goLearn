package main

import "fmt"

func main01()  {
	m:=make(map[int]string,1)
	/**
	map[keyType]valueType
	map[key] key 是一个基本数据类型
	map的长度是自动扩容的
	map中的数据是无序存储的
	 */
	m[100]="张三"
	m[2]="李四"
	m[5]="王五"
	//for i,v index value
	for k,v:=range m {
		fmt.Println(k,v)
	}
	//fmt.Println(m)
}

func main()  {
	//map在定义时 key是唯一的 不允许重复
	//m:=map[string]int{"赵四":10,"赵四":20}//err
	//m:=map[string]int{"赵四":10,"王五":20}
	m:=make(map[string]int,1)
	//重新赋值
	m["赵四"]=10
	_,ok:=m["赵四"]
	if ok {
		fmt.Println("该key已存在")
	}else{
		m["赵四"]=20
	}
	fmt.Println(m)
}