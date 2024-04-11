package vm

import (
	"chapter9/model"
	"log"
)

type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

type LoginViewModelOp struct {
}

func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// CheckLogin func
func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	//这里其实是对user的一次封装，得到私有的字段
	if err != nil {
		log.Println("Can not find username:", username)
		log.Println("Error:", err)
		return false
	}
	return user.CheckPassword(password) //对user方法进行封装
}
