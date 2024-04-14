package vm

import (
	"chapter10/model"
	"log"
)

type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

type LoginViewModelOp struct {
}

/*
使用空结构体来定义另一个结构体
从软件工程的角度，我们通常希望看到的是对数据结构的操作
而不是对数据结构本身的操作。
注意：我们不单单是指这个空结构体，而且是指这个空结构体要操作这个数据结构所带来的方法
可扩展就是我们以后对这个结构体需要进行更多的操作也好说，因为控制权
实际是在GetVm上，扩展开放，修改封闭
这里解释一下：扩展开放指的是
当需求发生改变的时候，我们不需要修改原有的代码，而是添加新的代码
对修改封闭是指一旦一个模块被测试好，就不应该再修改它。所有的修改都应该
通过添加新的代码来实现
*/
func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// CheckLogin func
func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	//这里其实是对user的一次封装，得到私有的字段
	if err != nil {
		log.Println("Can not find username:", username)
		log.Println("Error:", err)
		return false
	}
	return user.CheckPassword(password) //对user方法进行封装
}
