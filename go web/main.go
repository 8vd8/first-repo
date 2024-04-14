package main

import (
	"chapter10/controller"
	"chapter10/model"
	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	//Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db) //设置全局db

	controller.Startup()

	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
	//context.ClearHandler(http.DefaultServeMux)时将默认服务器包装在一个上下文清除处理器中
	//防止潜在的内存泄漏问题

}
