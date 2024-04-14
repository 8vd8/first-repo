package vm

import "chapter10/model"

type ResetPasswordViewModel struct {
	LoginViewModel
	Token string
}
type ResetPasswordViewModelOp struct {
}

func (ResetPasswordViewModelOp) GetVM(token string) ResetPasswordViewModel {
	v := ResetPasswordViewModel{}
	v.SetTitle("Reset Password")
	v.Token = token
	return v
}

func CheckToken(tokenString string) (string, error) {
	return model.CheckToken(tokenString)
}

func ResetUserPassword(username, password string) error {
	return model.UpdatePassword(username, password)
}
