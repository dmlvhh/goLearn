package main

import "fmt"
func main1() {
	//嵌套循环中  执行次数为外层*内层
	for i:=0; i<5; i++{
		for j:=0;j<5;j++ {
			fmt.Println(i,j)
		}
	}
}

//99乘法表
func main2()  {
	for i:=1;i<=9;i++ {
		for j:=1;j<i;j++ {
			fmt.Printf("%d*%d=%d ",j,i,i*j)
		}
		fmt.Println()
	}
}
//--------------------------------------------------------
func main()  {
	//整体循环控制行
	for i:=0;i<=5;i++ {
		//控制空格个数
		for j:=0;j<5-i;j++ {
			fmt.Print(" ")
		}
		//控制星星个数
		for k:=0;k<i*2+1;k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
