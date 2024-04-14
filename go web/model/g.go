package model

import (
	"chapter10/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func ConnectToDB() *gorm.DB {
	/*
		一般来说我们是直接将数据库的地址直接写进去
		但是在工作中的时候其实将配置文件单独放置
	*/
	//这里我们将得到一个数据库的信息包括端口号啥的
	connectingStr := config.GetMysqlConnectingString()
	log.Println("Connect to db...")
	db, err := gorm.Open("mysql", connectingStr) //链接数据库
	if err != nil {
		panic(err)
	}
	db.SingularTable(true) //使用单数表名
	return db
}
