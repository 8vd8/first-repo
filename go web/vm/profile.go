package vm

import "chapter10/model"

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
func (ProfileViewModelOp) GetPopupVM(sUser, pUser string) (ProfileViewModel, error) {
	//sUser表示当前用户的用户名，pUser表示要我们鼠标某用户停留长时间的那个用户
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	v.ProfileUser = *u
	v.Editable = (sUser == pUser)
	if !v.Editable { //pUser是自己的话，就不会执行关注操作
		// 如果不是自己，就会检查是否关注了这个人
		v.IsFollow = u.IsFollowedByUser(sUser)
		//这样做的意义是提供一种用户友好的界面交互
	}
	v.FollowersCount = u.FollowersCount()
	v.SetCurrentUser(sUser)
	return v, nil
}
