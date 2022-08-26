package main

import "fmt"

//if 可以嵌套 可以判断区间 执行效率低
//switch 执行效率高 不能嵌套区间判断
func main1()  {
	var w int
	fmt.Scan(&w)
	//switch中的值不能是浮点类型的数据
	//switch 选择项可以是一个整型变量
	switch w {
	case 1:
		fmt.Println("星期1")
	case 2:
		fmt.Println("星期2")
	case 3:
		fmt.Println("星期3")
	case 4:
		fmt.Println("星期4")
	case 5:
		fmt.Println("星期5")
	case 6:
		fmt.Println("星期6")
	case 7:
		fmt.Println("星期日")
		//如果输入的未找到则进入default
	default:
		fmt.Println("输入错误")
	}

	//switch score>700 {
	//case true:
	//	fmt.Println(score)
	//case false:
	//	fmt.Println(score)
	//}
}

func main()  {
	var score int
	fmt.Println("请输入分数")
	fmt.Scan(&score)

	switch score/10 {
	case 10:
		//fmt.Println("A")
		fallthrough //让switch执行下一分支的代码 如果不写 执行到下一分支就会自动停止
	case 9: fmt.Println("A")
	case 8: fmt.Println("B")
	case 7: fmt.Println("C")
	case 6: fmt.Println("D")
	default: fmt.Println("E")

	}
}