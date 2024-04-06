package main

import (
	"chapter5/controller"
	"chapter5/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	//Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db) //设置全局db

	controller.Startup()

	http.ListenAndServe(":8888", nil)
	
}
