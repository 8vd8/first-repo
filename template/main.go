package main

import (
	"html/template"
	"net/http"
)

type User struct {
	UserName string
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		user := User{
			"bonfy",
		}
		tpl, err := template.ParseFiles("C:\\Users\\Lenovo\\Desktop\\learn_web\\template\\templates\\index.html")
		if err != nil {
			panic(err)
		}
		tpl.Execute(writer, &user)
	})
	http.ListenAndServe("localhost:8080", nil)
}
