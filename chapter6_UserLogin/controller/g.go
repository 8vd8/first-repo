package controller

import (
	"github.com/gorilla/sessions"
	"html/template"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	store          *sessions.CookieStore
	pageLimit      int
)

func init() {
	templates = PopulateTemplates()
	//基于cookie的会话存储对象，将其存储在全局变量中
	//提供一个字节数组作为参数，用于加密和验证Cookie
	store = sessions.NewCookieStore([]byte("something-very-secret"))
	sessionName = "go-mega"
}

// Startup func
func Startup() {
	homeController.registerRoutes()
}
