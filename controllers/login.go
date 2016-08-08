package controllers

import (
	. "GoodsManage/models"
	"fmt"
	_ "io/ioutil"
	_ "net/http"
	_ "strconv"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session"
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
	fmt.Println("login.Service.User.Get(1)", user.Get(1))
	this.TplNames = "login/login.html"
}
func (this *LoginController) Post() {
	var username = this.GetString("uname")
	var password = this.GetString("psw")

	var user = new(Tuser).Init()
	var b = user.Login(username, password)
	fmt.Println(user)
	fmt.Println(b)
	if false == b {
		this.Ctx.WriteString("alert('用户名或密码有误！');")
		return
	}
	this.SetSession("uid", user.ID)
	this.SetSession("uname", username)
	// this.Ctx.SetCookie("uid", user.ID)
	// this.Ctx.SetCookie("uname", username)
	this.Ctx.WriteString("1")

}
func (this *LoginController) Uppwd() {
	//this.Layout = "main/index.html"
	this.TplNames = "login/uppwd.html"
}
func (this *LoginController) UppwdPost() {
	var pwd = this.GetString("pwd")
	var pwd1 = this.GetString("pwd1")
	var pwd2 = this.GetString("pwd2")
	//var id = this.GetSession("uid")
	//var id = this.Ctx.GetCookie("uid")
	var id, _ = this.GetInt64("id", 0)
	fmt.Println("id is:", id)
	// var uid, _ = strconv.ParseInt(id, 10, 64)
	// fmt.Println("uid is", uid)
	fmt.Println(pwd, pwd1, pwd2)
	if pwd1 == "" || pwd2 == "" || pwd1 != pwd2 {
		this.Ctx.WriteString("alert('两次密码不形同！');")
	}
	var user = new(Tuser).Init()

	var b = user.HuoQU(pwd)
	if b == false {
		this.Ctx.WriteString("alert('原密码不正确！');")
	}
	user.PASSWORD = pwd1
	var h = user.Update(id)
	fmt.Println()
	if h > 0 {
		this.Ctx.WriteString("alert('修改成功！');")
	} else {
		this.Ctx.WriteString("alert('修改失败！');")
	}

}
