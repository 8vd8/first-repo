package vm

import (
	"chapter10/config"
	"chapter10/model"
)

type EmailViewModel struct {
	Username string
	Token    string
	Server   string
}

type EmailViewModelOp struct {
}

func (EmailViewModelOp) GetVM(email string) EmailViewModel {
	v := EmailViewModel{}
	u, _ := model.GetUserByEmail(email)
	v.Username = u.Username
	v.Token, _ = u.GenerateToken() //jwt令牌是一个字符串
	v.Server = config.GetServerURL()
	return v

}
