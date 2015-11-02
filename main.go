package main

import (
	_ "GoodsManage/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SessionOn = true
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/easyui", "static/easyui")
	beego.SetStaticPath("/bootstrap", "static/bootstrap")
	beego.Run()
}
