package main

import "fmt"

func main() {
	arr:=[10]int{1,-2,3,8,9,40,32,10,22,33}
	//外层控制行
	for i:=0; i<len(arr)-1;i++{
		//内层控制列
		for j:=0;j<len(arr)-1-i;j++ {
			//比较两个相邻元素 满足条件交换数据
			//升序排序大于号  降序小于号
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}
