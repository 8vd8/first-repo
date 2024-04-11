package vm

import (
	"chapter8/model"
	"log"
)

type RegisterViewModel struct {
	LoginViewModel
}

type RegisterViewModelOp struct {
}

func (RegisterViewModelOp) GetVM() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

func CheckUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username:", username)
		return true
	}
	return false
}

func AddUser(username, password, email string) error {
	return model.AddUser(username, password, email) //再次封装
}
