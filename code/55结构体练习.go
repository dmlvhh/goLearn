package main

import "fmt"

type stud struct {
	id    int
	name  string
	score []int //结构体成功数组或切片
}

func main() {
	arr := []stud{stud{101, "小明", []int{100, 99, 88}},
		stud{102, "小红", []int{88, 79, 84}},
		stud{103, "小刚", []int{100, 90, 81}},
		stud{104, "小强", []int{86, 96, 83}},
		stud{105, "小花", []int{70, 90, 77}}}
	for i :=0;i<len(arr);i++ {
		sum:=0
		for j:=0;j<len(arr[i].score);j++ {
			sum+=arr[i].score[j]
		}
		fmt.Printf("第%d名学生总成绩为：%d 平均成绩：%d\n",i+1,sum,sum/3)
	}

}
