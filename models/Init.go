package models

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var X *xorm.Engine
var Usera *Tuser

func init() {
	var err error
	X, err = xorm.NewEngine("mssql", "server=.;database=goodsmanage;user id=sa;password=yiche2014")
	fmt.Println("-------11111")
	if err != nil {
		log.Fatalf("failt to carete engine")
	}
	fmt.Println("-------11111")
	X.SetMapper(core.SameMapper{})
	//用户表
	if err = X.Sync2(new(Tuser)); err != nil {
		log.Fatalf("failt Tuser to sync", err.Error())
	}
	//菜单表
	if err = X.Sync2(new(TMenus)); err != nil {
		log.Fatalf("failt TMenus to sync", err.Error())
	}
	//角色表
	if err = X.Sync2(new(TRoles)); err != nil {
		log.Fatalf("failt TRoles to sync", err.Error())
	}
	//初始管理员
	// var tuser = new(Tuser).Init()
	// var d = tuser.GetUser("admin")
	// if d == false {
	// 	_, _ = X.Exec("insert into [Tuser]([uname],password)values('admin','admin')")
	// }
	var tuser = new(Tuser).Init()
	if !tuser.Get(1) {
		_, _ = X.Exec("insert into [Tuser]([uname],password)values('admin','admin')")
	}

}
