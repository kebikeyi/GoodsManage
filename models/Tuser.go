package models

import (
	"fmt"
)

type Tuser struct {
	ID       int64
	UNAME    string `xorm:"varchar(50)"` //用户名
	PASSWORD string `xorm:"varchar(50)"` //密码
	_godb    `xorm:"-"`
}

func (this *Tuser) Init() *Tuser {
	this._godb._bean = this
	this._godb._beans = make([]Tuser, 0)
	return this
}

//登录
func (this *Tuser) Login(uname string, psw string) bool {
	has, err := X.Where("name=? and psw=?", uname, psw).Get(this)
	fmt.Println(has)
	if err != nil {
		fmt.Println("------11111")
		return false
	}
	return has
}

//获取用户
func (this *Tuser) GetUser(uname string) bool {
	has, err := X.Where("name=?", uname).Get(this)
	if err != nil {
		return false
	}
	return has
}
