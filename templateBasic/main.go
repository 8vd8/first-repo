package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

// User struct
type User struct {
	Username string
}

// Post struct
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

// PopulateTemplates func
// Create map template name to template.Template
func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates" //定义模板文件的基本路径
	result := make(map[string]*template.Template)
	//将基本布局模板‘_base.html’，并将其存储在变量‘layout’中
	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content") //打开目标文件夹
	if err != nil {
		panic("Failed to open template blocks directory:" + err.Error())
	}
	fis, err := dir.ReadDir(-1) //Readdir(-1)会读取目录中的文件信息，并返回一个文件信息的切片
	/*
			如果ReadDir里面传入的是一个正整数n，表示最多读取n个文件
			fis是读取到的文件信息切片，每个元素代表目录中的一个文件的信息
		type DirEntry interface {
		    Name() string
		    IsDir() bool
		    Type() FileMode
		    Info() (FileInfo, error)
		}
	*/
	if err != nil {
		panic("Failed to read contents of content directory:" + err.Error())
	}

	for _, fi := range fis {
		func() {
			f, err := os.Open(basePath + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template '" + fi.Name() + "'")
			}
			defer f.Close()
			content, err := ioutil.ReadAll(f) //读这个文件
			if err != nil {
				panic("Failed to read content from file'" + fi.Name())
			}
			tmpl := template.Must(layout.Clone())
			/**
			使用layout.Clone()复制基本布局模板，然后将读取的内容解析为模板
			并于基本布局合并
			*/
			_, err = tmpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "' as template")
			}
			result[fi.Name()] = tmpl
		}() //这里定义函数
	}
	return result
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := User{Username: "bonfy"}
		u2 := User{Username: "rene"}

		posts := []Post{
			Post{User: u1, Body: "Beautiful day in Portland!"},
			Post{User: u2, Body: "The Avengers movie was so cool!"},
		}

		v := IndexViewModel{Title: "Homepage", User: u1, Posts: posts}

		templates := PopulateTemplates()
		templates["index.html"].Execute(w, &v)
	})
	http.ListenAndServe(":8888", nil)
}
