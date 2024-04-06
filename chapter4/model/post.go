package model

import "time"

// Post struct
type Post struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	User      User
	Body      string     `gorm:"type:varchar(180)"`
	Timestamp *time.Time `sql:"DEFAULT:current_timestamp"`
}

func GetPostByUserID(id int) (*[]Post, error) {
	var posts []Post
	if err := db.Preload("User").Where("user_id=?", id).Find(&posts).Error; err != nil {
		return nil, err
	}
	/*
			使用db对象进行查询，在查询之前，
			通过Preload(User)方法预加载用户信息，便于在获取帖子时一并获取相关的用户信息
		    将查询结果存储到posts切片中
	*/
	return &posts, nil
}
