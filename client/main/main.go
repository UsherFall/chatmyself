package main

import (
	"chatmyself/client/processes"
	"fmt"
)

func main() {

	//客户端初始界面
	var userName string
	fmt.Println("告诉我你是谁！!")
	fmt.Scanf("%s", &userName)

	//完成用户登录
	up := processes.Userprocess{}
	up.Login(userName)

}
