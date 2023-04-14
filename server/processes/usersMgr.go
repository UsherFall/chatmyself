package processes

var usersMgr *UsersMgr

type UsersMgr struct {
	OnlineUsers map[string]*Userprocess
}

//初始化userMgr
func init() {
	usersMgr = &UsersMgr{
		OnlineUsers: make(map[string]*Userprocess, 1024),
	}
}

//添加onlineUsers
func (um *UsersMgr) AddOnlineUsers(userName string, up *Userprocess) {
	um.OnlineUsers[userName] = up
}
