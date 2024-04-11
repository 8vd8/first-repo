package model

import (
	"fmt"
	"time"
)

// User struct
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorm:"type:varchar(128)"` //varchar(n)表示可以容纳n个字符
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
	/*
		Follower and User是多对多的关系
		association_jointable_foreignkey:follower_id表示在关联表中
		指向Follower的外键的列明是：follower_id
	*/
	LastSeen *time.Time
	AboutMe  string `gorm:"type:varchar(140)"`
	Avatar   string `gorm:"type:varchar(200)"`
}

func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", Md5(email))
}

func (u *User) SetPassword(pwd string) {
	u.PasswordHash = GeneratePasswordHash(pwd) //直接对传入的密码进行哈希值加密
}

func (u *User) CheckPassword(pwd string) bool {
	return GeneratePasswordHash(pwd) == u.PasswordHash
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
	/*
		现在我们在
	*/
}

func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)
	user.SetAvatar(email)
	return db.Create(&user).Error
}

func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	item, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(item).Update(contents).Error
	/*
		使用GORM方法指定更新数据库模型
	*/
}

func UpdateLastSeen(username string) error {
	contents := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(username, contents)
}

func UpdateAboutMe(username, text string) error {
	contents := map[string]interface{}{"about_me": text}
	return UpdateUserByUsername(username, contents)
}
