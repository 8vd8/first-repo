package vm

import "chapter8/model"

// 创建个人主页
type ProfileViewModel struct {
	BaseViewModel
	Posts          []model.Post
	Editable       bool
	IsFollow       bool
	FollowersCount int
	FollowingCount int
	ProfileUser    model.User
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
	u, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, _ := model.GetPostByUserID(u.ID)
	v.ProfileUser = *u
	v.Editable = (sUser == pUser)
	v.FollowersCount = u.FollowersCount()
	v.FollowingCount = u.FollowingCount()
	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, nil
}

// a follow b
func Follow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.Follow(b)
}

func UnFollow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.Unfollow(b)
}
