package main

import "fmt"

func test01(arr [10]int)  {
	//for i:=0;i<len(arr);i++{
	//	fmt.Println(arr[i])
	//}
	arr[2]=123
	fmt.Printf("%p\n",&arr)

}
func main1() {
	//指定数组下标进行初始化操作
	//var arr [10]int=[10]int{1,2,7:10}
	arr:=[10]int{9,1,5,6,7,3,10,2,4,8}
	//数组作为函数参数是传递数组名
	//数组作为函数参数是值传递的
	test01(arr)
	//fmt.Println(arr)
	fmt.Printf("%p\n",&arr)
	//fmt.Println(arr)
}

//-----------------------------------------
func BubbleSort(arr [10]int) [10]int{
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-1-i;j++ {
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]

			}

		}
	}
	return arr
}
func main() {
	arr:=[10]int{9,1,5,6,7,3,10,2,4,8}
	fmt.Printf("%T\n",arr)//[10]int
	arr=BubbleSort(arr)
	fmt.Println(arr)
}
