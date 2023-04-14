package processes

import (
	"bufio"
	"chatmyself/common/message"
	"chatmyself/utils"
	"fmt"
	"net"
	"os"
)

type Smsprocess struct {
	Conn     net.Conn
	UserName string
}

func (sp *Smsprocess) SendChatMes() {

	var content string
	var mes message.Message
	var chatMes message.ChatMes
	mes.Type = message.ChatMesType
	chatMes.UserName = sp.UserName

	for {

		//获取消息内容
		stdin := bufio.NewReader(os.Stdin)
		_, err := fmt.Fscan(stdin, &content)
		if err != nil {
			fmt.Println("获取内容错误", err)
			return
		}

		//组建ChatMes
		chatMes.Content = content

		//序列化并组建消息
		tf := &utils.Transfer{
			Conn: sp.Conn,
		}
		data := tf.MarshalMes(&mes, chatMes)

		//发送消息
		err = tf.WritePkg(data)
		if err != nil {
			return
		}

	}
}
