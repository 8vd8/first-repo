package controller

import (
	"log"
	"net/http"

	"chapter6/vm"
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", middleAuth(indexHandler))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", middleAuth(logoutHandler))
	http.HandleFunc("/register", registerHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetVM()

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v) //template[a] a是指渲染的html文件 ，v是后端的处理文件，这两个会结合在一起，反馈成w
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if len(username) < 3 {
			v.AddError("username must longer than 3")
		}

		if len(password) < 6 {
			v.AddError("password must longer than 6")
		}

		if !vm.CheckLogin(username, password) {
			v.AddError("username password not correct, please input again")
		}

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v) //重新回到这个页面，并且注意此时 返回的v是有错误了的
		} else {
			setSessionUser(w, r, username)                //开启session
			http.Redirect(w, r, "/", http.StatusSeeOther) //进入主页
		}
	}

}
func indexHandler(w http.ResponseWriter, r *http.Request) { //进入主页将用到的处理器
	tpName := "index.html"
	vop := vm.IndexViewModelOp{} //这个功能和上面那个功能不一样了
	//原来他每一个功能都给了不同的结构体
	//这样在每一个html文件中才好渲染
	username, _ := getSessionUser(r)
	//现在我们这里直接从会话中得到user的信息
	v := vop.GetVM(username) //得到user,posts,设置currentUser = user
	templates[tpName].Execute(w, &v)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w, r)                                          //将会话结束
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect) //直接重新定向到登录页面
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "register.html" //原始模板
	vop := vm.RegisterViewModelOp{}
	v := vop.GetVM() //只包含标题和errs

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errs := checkRigister(username, email, pwd1, pwd2)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			if err := addUser(username, pwd1, email); err != nil {
				log.Println("add User error:", err)
				w.Write([]byte("Error insert database"))
				return
			}
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

}
