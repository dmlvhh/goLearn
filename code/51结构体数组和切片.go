package main

import "fmt"

type student struct {
	id      int
	name    string
	sex     string
	age     int
	score   int
	address string
}

func main01() {

	/**
	var 结构体变量名 结构体变量
	定义结构体数组
	var 结构体数组名[元素个数]结构体类型
	*/
	var arr [5]student

	//len(数组名)计算数组元素个数
	//fmt.Println(len(arr))
	/**
	  101 孙尚香 女 16 88 江东
	  102 黄月英 女 28 90 襄阳
	  103 大乔 女 26 70 江东
	  104 小乔 女 24 60 江东
	  105 甄姬 女 20 50 许昌
	*/
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i].id, &arr[i].name, &arr[i].sex, &arr[i].age, &arr[i].score, &arr[i].address)
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			//结构体数组中的元素 允许相互赋值 将结构体成员中的所有数据进行相互交换
			if arr[j].score > arr[j+1].score {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	//[元素个数]数组   []切片len cap
	arr := []student{{101, "曹操", "男", 58, 90, "许昌"},
		{102, "夏侯惇", "男", 40, 100, "荆州"},
		{103, "张辽", "男", 38, 99, "许昌"}}

	//在切片中添加数据
	arr = append(arr, student{104, "许褚", "男", 28, 99, "许昌"})
	arr = append(arr, student{105, "典韦", "男", 30, 98, "许昌"})
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
