package vm

import "chapter9/model"

type ExploreViewModel struct {
	BaseViewModel
	Posts []model.Post
	BasePageViewModel
}

type ExploreViewModelOp struct {
}

func (ExploreViewModelOp) GetVM(username string, page, limit int) ExploreViewModel {
	posts, total, _ := model.GetPostByPageAndLimit(page, limit)
	v := ExploreViewModel{}
	v.SetTitle("Explore")
	v.Posts = *posts
	v.SetCurrentUser(username)
	v.SetBasePageViewModel(total, page, limit)
	return v
}
