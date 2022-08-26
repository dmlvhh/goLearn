package main

import "fmt"

func main1() {

	//var 变量名 数据类型
	//数组定义和使用
	//var arr [10]int
	//数组名[下标]=值
	//var arr [10]int
	//arr[0]=123
	//arr[1]=10

	var arr [10]int =[10] int{1,2,3,4,5,6,7,8,9,10}
	//fmt.Println(arr[0])
	//fmt.Println(arr[1])
	//fmt.Println(arr[2])
	//fmt.Println(arr[3])
	////打印数组名可以打印出所有元素数据
	//fmt.Println(arr)

	//遍历数组元素信息
	//for i:=0;i<10;i++ {
	//	fmt.Println(arr[i])
	//}

	//len(string) 获取字符串有效长度
	//len(...int) 获取不定参函数参数个数
	//len([10]int) 计算数组元素个数

	//fmt.Println(len(arr))
	//for i:=0;i<len(arr);i++ {
	//	fmt.Println(arr[i])
	//}

	//range 遍历集合信息
	//返回值为 下标 值
	for _,v:=range arr {
		fmt.Println(v)
	}
}

//--------------------------------------------------
func main2()  {
	//数组在定义时 可以初始化部分元素的值
	//var arr [10]int=[10]int{1,2,3,4,5}
	//使用自动推导类型创建数组
	arr:=[...]int{1,2,3,4,5,6,7,8,9,10}

	for _,v:=range arr {
		fmt.Println(v)
	}
	fmt.Printf("%T",arr)
}

//-----------------------------------------------------
func main3()  {
	//arr:=[5]int{1,2,3,4,5}
	//在定义 叫元素个数
	var arr [5]int=[5]int{1,2,3,4,5}
	//在数组使用时叫下标
	//arr[5]=100 //数组下标越界
	//arr[-1]=100
	//数组在定义后 元素个数就已经固定 不能添加
	//fmt.Println(arr)
	//数组是一个常量 不允许赋值 数组名代表整个数组
	//arr=10 //err
	//数组名也可以表示数组首地址
	fmt.Printf("%p\n",&arr)
	fmt.Printf("%p\n",&arr[0])
}
//--------------------------------------------
//10只小猪称体重 谁最大
func main()  {
	arr:=[10]int{1,-2,3,8,9,40,32,10,22,33}

	max:=arr[0]
	for i:=1;i<len(arr);i++ {
		if max<arr[i] {
			max=arr[i]
		}
	}
	fmt.Println("最重的小猪体重是：",max)
}