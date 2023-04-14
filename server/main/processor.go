package main

import (
	"chatmyself/common/message"
	"chatmyself/server/processes"
	"chatmyself/utils"
	"fmt"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//根据消息类型选择不同的处理函数
func (ps *Processor) processing() {

	tf := &utils.Transfer{
		Conn: ps.Conn,
	}

	for {
		//创建tf实例，读取客户端信息
		mes := tf.ReadPkg()
		if mes == (message.Message{}) {
			fmt.Println("有人离开！！但我不知道是谁！！")
			return
		}

		//根据消息类型确定调用的函数
		ps.severProcess(&mes)
	}

}

//根据消息类型确定调用的函数
func (ps *Processor) severProcess(mes *message.Message) {

	switch mes.Type {
	case message.LoginMesType:

		//进行登录处理
		up := processes.Userprocess{
			Conn: ps.Conn,
		}
		up.Login(mes)

	case message.ChatMesType:

		//聊天处理
		sp := processes.Smsprocess{}
		sp.PrintChatMes(mes)

	default:
		fmt.Println("啊？")
	}
}
