package models

import (
	"fmt"
)

type Tuser struct {
	ID        int64
	UNAME     string `xorm:"varchar(50)"` //用户名
	PASSWORD  string `xorm:"varchar(50)"` //密码
	Nick      string `xorm:"varchar(32)"`
	Type      int    //
	AddTime   string `xorm:"varchar(8)"`    //添加时间
	Rights    string `xorm:"varchar(2000)"` //权限
	Roles     string `xorm:"varchar(500)"`  //角色
	LoginTime string `xorm:"varchar(8)"`    //登录时间
	LoginIP   string `xorm:"varchar(50)"`   //登录IP
	Memo      string `xorm:"varchar(200)"`  //
	State     int    //
	_godb     `xorm:"-"`
}

func (this *Tuser) Init() *Tuser {
	this._godb._bean = this
	this._godb._beans = make([]Tuser, 0)
	return this
}

//登录
func (this *Tuser) Login(uname string, psw string) bool {
	has, err := X.Where("uname=? and password=?", uname, psw).Get(this)
	fmt.Println(has)
	if err != nil {
		fmt.Println("------11111")
		return false
	}
	return has
}

//获取用户
func (this *Tuser) GetUser(uname string) bool {
	has, err := X.Where("uname=?", uname).Get(this)
	if err != nil {
		return false
	}
	return has
}

//获取原始密码
func (this *Tuser) HuoQU(pwd string) bool {
	has, err := X.Where("password=?", pwd).Get(this)
	fmt.Println("has is", has)
	if err != nil {
		return false
	}
	return has
}
