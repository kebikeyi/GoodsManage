package controllers

import (
	. "GoodsManage/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	var User = new(User).Init()
	fmt.Println("login.user", User.Query("select * from User"))
	fmt.Println("login.Service.User", User.Query("select * from User"))
	fmt.Println("login.Service.User.Get(2)", User.Get(2))
	this.TplNames = "login/login.html"
}
func (this *LoginController) Post() {
	var username = this.GetString("uname")
	var password = this.GetString("psw")

	var user = new(User).Init()
	var b = user.Login(username, password)
	if b == false {
		this.Ctx.WriteString("用户名或密码有误！")
		return
	}
	this.SetSession("uid", user.ID)
	this.SetSession("uname", username)
	this.Ctx.WriteString("1")

}
