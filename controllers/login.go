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
	var user = new(Tuser).Init()
	fmt.Println("login.user", user.Query("select * from Tuser"))
	fmt.Println("login.Service.User", user.Query("select * from Tuser"))
	fmt.Println("login.Service.User.Get(2)", user.Get(1))
	this.TplNames = "login/login.html"
}
func (this *LoginController) Post() {
	var username = this.GetString("uname")
	var password = this.GetString("psw")

	var user = new(Tuser).Init()
	var b = user.Login(username, password)
	fmt.Println(username, password)
	fmt.Println(b)
	if false == b {
		this.Ctx.WriteString("alert('用户名或密码有误！');")
		return
	}
	this.SetSession("uid", user.ID)
	this.SetSession("uname", username)
	this.Ctx.WriteString("1")

}
