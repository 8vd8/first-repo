package vm

import "chapter9/model"

type ProfileEditViewModel struct {
	LoginViewModel
	ProfileUser model.User
}

type ProfileEditViewModelOp struct {
}

func (ProfileEditViewModelOp) GetVM(username string) ProfileEditViewModel {
	v := ProfileEditViewModel{}
	u, _ := model.GetUserByUsername(username)
	v.SetTitle("Profile Edit")
	v.SetCurrentUser(username)
	v.ProfileUser = *u
	return v
}

func UpdateAboutMe(username, text string) error {
	return model.UpdateAboutMe(username, text)
}
