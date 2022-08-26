package main

import "./userinfo"
//如果调用不同级别目录需要导入包
//全局变量
var num int = 123
//在同级别目录报名要相同
func main() {
	add(10,20)
	userinfo.Login()
	userinfo.DeleteUser()
	userinfo.SelectUser()
}
