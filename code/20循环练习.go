package main

import "fmt"

func main1() {
	//水仙花
	for i:=100; i<1000;i++ {
		a:=i/100
		b:=i/10%10
		c:=i%10
		if a*a*a+b*b*b+c*c*c==i {
			fmt.Println("水仙花数为：",i)
		}
	}
}

func main()  {
	//逢7过
	//1-100 逢7 7的倍数 个位有7 十位有7 过
	for i:=1;i<=100;i++ {
		if i%7==0 || i%10==7 || i/10==7 {
			fmt.Println("过")
		}else {
			fmt.Println(i)
		}
	}
}
