package vm

import "chapter7/model"

// 创建个人主页
type ProfileViewModel struct {
	BaseViewModel
	Posts       []model.Post
	Editable    bool
	ProfileUser model.User
}

type ProfileViewModelOp struct {
}

func (ProfileViewModelOp) GetVM(sUser, pUser string) (ProfileViewModel, error) {
	/*
		sUser:表示当前会话的用户，就是正在浏览网站的用户
		pUser:表示 要查看个人资料的用户，就是自己查自己
	*/
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u1, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, _ := model.GetPostByUserID(u1.ID)
	v.ProfileUser = *u1
	v.Editable = (sUser == pUser)
	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, nil
}
