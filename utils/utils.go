package utils

import (
	"chatmyself/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
}

//发送消息
func (tf *Transfer) WritePkg(data []byte) (err error) {

	//将mes发送至服务器
	_, err = tf.Conn.Write(data)
	if err != nil {
		fmt.Println("发送mes错误", err)
		return
	}

	return

}

//读取消息
func (tf *Transfer) ReadPkg() (mes message.Message) {

	//读取客户端发送的信息
	buf := make([]byte, 2048)
	n, err := tf.Conn.Read(buf)
	if err != nil {
		if buf[0] == 0 {
			return
		}
		fmt.Println("读取信息错误", err)
		return
	}

	//将信息反序列化,并用mes接收
	err = json.Unmarshal(buf[:n], &mes)
	if err != nil {
		fmt.Println("反序列化Message错误", err)
		return
	}
	return

}

//序列化并构建消息
func (tf *Transfer) MarshalMes(mes *message.Message, mesData interface{}) (data []byte) {

	//将mesData序列化
	data, err := json.Marshal(mesData)
	if err != nil {
		fmt.Printf("序列化%v错误 %v\n", mes.Type, err)
		return
	}

	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化mes错误", err)
		return
	}
	return

}
