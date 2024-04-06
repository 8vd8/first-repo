package main

import (
	"chapter4/model"
	"log"
)

func main() {
	//初始化数据库中的表
	log.Println("DB init")
	db := model.ConnectToDB()
	defer db.Close()
	//设置全局数据库
	model.SetDB(db)
	//每次都是一个重新来的感觉
	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})

	users := []model.User{
		{
			Username:     "bonfy",
			PasswordHash: model.GeneratePasswordHash("abc213"),
			Posts: []model.Post{
				{Body: "Beautiful day in China"},
			},
		},
		{
			Username:     "JackieChen",
			PasswordHash: model.GeneratePasswordHash("chaojie666"),
			Posts: []model.Post{
				{Body: "I hope I will be braver"},
				{Body: "Hope you so"},
			},
		},
	}
	for _, u := range users {
		db.Debug().Create(&u) //这里启用调试模式，意味着执行数据库操作会输出相应的SQL语句
	}

}
