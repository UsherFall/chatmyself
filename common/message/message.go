package message

const (
	LoginMesType    = "loginMes"
	LoginResMesType = "loginResMes"
	ChatMesType     = "chatMes"
)

type Message struct {
	Type string
	Data string
}

type LoginMes struct {
	UserName string
}

type LoginResMes struct {
	UserName  string
	UsersName []string
}

type ChatMes struct {
	Content  string
	UserName string
}
