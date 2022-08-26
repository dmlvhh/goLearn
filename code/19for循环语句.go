package main

import "fmt"

func main1() {
	//for 表达式1 ;表达式2 ;表达式3{  循环体  }

	//for i:=1;i<=10;i++{
	//	fmt.Println(i)
	//}

	//计算1-100和
	//sum:=0
	//for i:=1;i<=100;i++ {
	//	sum+=i
	//}
	//fmt.Println(sum)

	//计算1-100偶数的和
	//sum:=0
	//for i:=1;i<=100;i++ {
	//	if i%2==0 {
	//		sum+=i
	//	}
	//}
	//fmt.Println(sum)
	sum:=0
	for i:=0;i<=100;i+=2 {
		sum+=i
	}
	fmt.Println(sum)
}

func main()  {
	//for函数格式 i是在函数内部定义的不能在外部使用
	//局部变量
	//for i:=1;i<=10;i++ {
	//	fmt.Println(i)
	//}
	//fmt.Println(i) //err

	//i:=1
	//for ;; {
	//	if i>10 {
	//	跳出循环
	//		break
	//	}
	//	fmt.Println(i)
	//	i++
	//}
	//fmt.Println(i)

	//死循环  在程序中要避免死循环出现  使用break  跳出循环
	for{
		fmt.Println("hello")
		break
	}
}
