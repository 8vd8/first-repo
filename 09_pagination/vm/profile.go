package vm

import "chapter9/model"

// 创建个人主页
type ProfileViewModel struct {
	BaseViewModel
	Posts          []model.Post
	Editable       bool
	IsFollow       bool
	FollowersCount int
	FollowingCount int
	ProfileUser    model.User
	BasePageViewModel
}

type ProfileViewModelOp struct {
}

func (ProfileViewModelOp) GetVM(sUser, pUser string, page, limit int) (ProfileViewModel, error) {
	/*
		sUser:表示当前会话的用户，就是正在浏览网站的用户
		pUser:表示要查看的人，这个用户可以是任何人
	*/
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, total, _ := model.GetPostByUserIDPageAndLimit(u.ID, page, limit)
	v.ProfileUser = *u
	v.Editable = (sUser == pUser)
	v.SetBasePageViewModel(total, page, limit)
	if !v.Editable { //不是查看自己
		//就需要关注是否关注了这个用户
		v.IsFollow = u.IsFollowedByUser(sUser)
	}
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
