package model

import (
	"fmt"
	"log"
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
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return user.FollowSelf()
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

func (u *User) Follow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Append(u).Error
}

func (u *User) Unfollow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Delete(u).Error
}

func (u *User) FollowSelf() error {
	return db.Model(u).Association("Followers").Append(u).Error
}

func (u *User) FollowersCount() int {
	return db.Model(u).Association("Followers").Count()
}

func (u *User) FollowingIDs() []int { //获取所有粉丝的id
	var ids []int
	rows, err := db.Table("follower").Where("follower_id = ?", u.ID).Select("user_id,follower_id").Rows()
	if err != nil {
		log.Println("Counting Following error:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, followerID int
		rows.Scan(&id, &followerID)
		ids = append(ids, id)
	}
	return ids
}

func (u *User) FollowingCount() int {
	ids := u.FollowingIDs()
	return len(ids)
}

func (u *User) FollowingPosts() (*[]Post, error) { //得到粉丝的帖子
	var posts []Post
	ids := u.FollowingIDs()
	if err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)", ids).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

func (u *User) IsFollowedByUser(username string) bool {
	user, _ := GetUserByUsername(username)
	ids := user.FollowingIDs()
	for _, id := range ids {
		if u.ID == id {
			return true
		}
	}
	return false
}

func (u *User) CreatePost(body string) error {
	post := Post{Body: body, UserID: u.ID}
	return db.Create(&post).Error
}
