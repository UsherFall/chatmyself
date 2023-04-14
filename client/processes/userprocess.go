package processes

import (
	"chatmyself/common/message"
	"chatmyself/utils"
	"encoding/json"
	"fmt"
	"net"
)

type Userprocess struct{}

func (up *Userprocess) Login(userName string) {

	//连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接服务器错误", err)
		return
	}

	//组建type为loginMes的message
	var mes message.Message
	var loginMes message.LoginMes

	mes.Type = message.LoginMesType
	loginMes.UserName = userName

	tf := &utils.Transfer{
		Conn: conn,
	}

	//序列化并构建消息
	data := tf.MarshalMes(&mes, loginMes)

	//发送消息
	tf.WritePkg(data)

	//读取返回的消息
	resMes := tf.ReadPkg()

	//对loginResMes反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(resMes.Data), &loginResMes)
	if err != nil {
		fmt.Println("反序列化loginResMes错误", err)
		return
	}

	fmt.Printf("先登录的有:")
	for _, name := range loginResMes.UsersName {
		fmt.Printf("%v\t", name)
	}

	fmt.Printf("\n")
	fmt.Printf("恭喜%v登录成功！说说话吧！\n", loginResMes.UserName)

	//开始聊天
	sp := &Smsprocess{
		Conn:     conn,
		UserName: userName,
	}
	sp.SendChatMes()

}
