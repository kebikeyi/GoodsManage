package models

type User struct {
	ID       int64
	UNAME    string `xorm:"varchar(50)"` //用户名
	PASSWORD string `xorm:"varchar(50)"` //密码
	_godb    `xorm:"-"`
}

func (this *User) Init() *User {
	this._godb._bean = this
	this._godb._beans = make([]User, 0)
	return this
}

//登录
func (this *User) Login(uname string, psw string) bool {
	has, err := X.Where("name=? and pwd=?", uname, psw).Get(this)
	if err != nil {
		return false
	}
	return has
}

//获取用户
func (this *User) GetUser(uname string) bool {
	has, err := X.Where("name=?", uname).Get(this)
	if err != nil {
		return false
	}
	return has
}
