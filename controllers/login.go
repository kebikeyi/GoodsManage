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

//LoginController 控制器
type LoginController struct {
	beego.Controller
}

// Get 登录
func (my *LoginController) Get() {
	my.Data["Website"] = "beego.me"
	my.Data["Email"] = "astaxie@gmail.com"
	var user = new(Tuser).Init()
	fmt.Println("login.user", user.Query("select * from Tuser"))
	fmt.Println("login.Service.User", user.Query("select * from Tuser"))
	fmt.Println("login.Service.User.Get(1)", user.Get(1))
	my.TplName = "login/login.html"
}

//Post 提交
func (my *LoginController) Post() {
	var username = my.GetString("uname")
	var password = my.GetString("psw")

	var user = new(Tuser).Init()
	var b = user.Login(username, password)
	fmt.Println(user)
	fmt.Println(b)
	if false == b {
		my.Ctx.WriteString("alert('用户名或密码有误！');")
		return
	}
	my.SetSession("uid", user.ID)
	my.SetSession("uname", username)
	// my.Ctx.SetCookie("uid", user.ID)
	// my.Ctx.SetCookie("uname", username)
	my.Ctx.WriteString("1")

}

// 密码修改
func (my *LoginController) Uppwd() {
	//this.Layout = "main/index.html"
	my.TplName = "login/uppwd.html"
}

// UppwdPost 修改提交
func (my *LoginController) UppwdPost() {
	var pwd = my.GetString("pwd")
	var pwd1 = my.GetString("pwd1")
	var pwd2 = my.GetString("pwd2")
	//var id = my.GetSession("uid")
	//var id = my.Ctx.GetCookie("uid")
	var id, _ = my.GetInt64("id", 0)
	fmt.Println("id is:", id)
	// var uid, _ = strconv.ParseInt(id, 10, 64)
	// fmt.Println("uid is", uid)
	fmt.Println(pwd, pwd1, pwd2)
	if pwd1 == "" || pwd2 == "" || pwd1 != pwd2 {
		my.Ctx.WriteString("alert('两次密码不形同！');")
	}
	var user = new(Tuser).Init()

	var b = user.HuoQU(pwd)
	if b == false {
		my.Ctx.WriteString("alert('原密码不正确！');")
	}
	user.PASSWORD = pwd1
	var h = user.Update(id)
	fmt.Println()
	if h > 0 {
		my.Ctx.WriteString("alert('修改成功！');")
	} else {
		my.Ctx.WriteString("alert('修改失败！');")
	}

}
