package routers

import (
	"GoodsManage/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterUser = func(ctx *context.Context) {
	uid, ok := ctx.Input.Session("uid").(int64)
	fmt.Println("auth check:", uid, ok)
	if !ok {
		// ctx.WriteString("<script language='javascript'>window.top.location='/login.html';</script>")
		ctx.Redirect(302, "/login")
		return
	}
}

func init() {
	beego.InsertFilter("/main", beego.BeforeRouter, FilterUser)
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.AutoRouter(&controllers.LoginController{})
}
