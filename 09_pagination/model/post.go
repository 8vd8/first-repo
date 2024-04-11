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
	//按照时间戳降序排序
	if err := db.Preload("User").Order("timestamp desc").Where("user_id=?", id).Find(&posts).Error; err != nil {
		return nil, err
	}
	/*
			使用db对象进行查询，在查询之前，
			通过Preload(User)方法预加载用户信息，便于在获取帖子时一并获取相关的用户信息
		    将查询结果存储到posts切片中
	*/
	return &posts, nil
}

func GetPostByUserIDPageAndLimit(id, page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (page - 1) * limit
	if err := db.Preload("User").Order("timestamp desc").Where("user_id=?", id).Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Where("user_id = ?", id).Count(&total)
	return &posts, total, nil
}

func GetPostByPageAndLimit(page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (page - 1) * limit //这里是说我们的起始位置是在最后面的数据开始
	//也就是最新的
	//如果没有limit，就会直接返回所有的匹配数据
	if err := db.Preload("User").Offset(offset).Limit(limit).Order("timestamp desc").Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Count(&total)
	return &posts, total, nil
}
