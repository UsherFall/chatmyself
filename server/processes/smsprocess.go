package processes

import (
	"chatmyself/common/message"
	"encoding/json"
	"fmt"
)

type Smsprocess struct{}

func (sp *Smsprocess) PrintChatMes(mes *message.Message) {

	//反序列化chatMes
	var chatMes message.ChatMes
	err := json.Unmarshal([]byte(mes.Data), &chatMes)
	if err != nil {
		fmt.Println("反序列化chatMes错误", err)
		return
	}

	//打印信息
	fmt.Printf("%s：%s\n", chatMes.UserName, chatMes.Content)

}
