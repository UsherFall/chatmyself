package processes

import (
	"chatmyself/common/message"
	"chatmyself/utils"
	"encoding/json"
	"fmt"
	"net"
)

type Userprocess struct {
	Conn net.Conn
}

func (up *Userprocess) Login(mes *message.Message) {

	//将loginMes反序列化
	var loginMes message.LoginMes
	err := json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("反序列化loginMes错误", err)
		return
	}

	//构建Type为loginResMes的message
	var resMes message.Message
	var loginResMes message.LoginResMes
	mes.Type = message.LoginResMesType

	//将上线用户添加至OnlineUsers
	usersMgr.AddOnlineUsers(loginMes.UserName, up)

	//构建loginResMes
	for name, _ := range usersMgr.OnlineUsers {
		loginResMes.UsersName = append(loginResMes.UsersName, name)
	}
	loginResMes.UserName = loginMes.UserName

	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	//序列化并构建mes
	data := tf.MarshalMes(&resMes, loginResMes)

	//发送mes给客户端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送消息错误")
		return
	}
	fmt.Println(loginMes.UserName, "已上线")

}
