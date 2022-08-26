package main

import "fmt"

func main() {
	chicken:=0
	count:=0
	for cock:=0;cock<=20;cock++ {
		for hen:=0;hen<=33;hen++ {
			count++
			chicken=100-cock-hen
			if cock*5+hen*3+chicken/3==100 && chicken%3==0 {
				fmt.Printf("公鸡：%d 母鸡：%d 小鸡：%d\n",cock,hen,chicken)
			}
		}
	}
	fmt.Println("执行次数为：",count)
}
