package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Username string
}
type Post struct {
	User
	Body string
}

// IndexViewModel struct
type IndexViewModel struct {
	Title string
	User
	Posts []Post
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		u1 := User{Username: "bofy"}
		u2 := User{Username: "rene"}
		posts := []Post{
			{User: u1, Body: "Hardwroking..."},
			{User: u2, Body: "加油努力，还有很长的路要走"},
		}
		v := IndexViewModel{Title: "Homepage", User: u1, Posts: posts}
		tpl, err := template.ParseFiles("C:\\Users\\Lenovo\\Desktop\\learn_web\\templateBasic\\templates\\_base.html")
		if err != nil {
			panic(err)
		}
		tpl.Execute(writer, &v)

	})
	http.ListenAndServe("localhost:8080", nil)
}
