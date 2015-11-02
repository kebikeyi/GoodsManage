package models

type TRoles struct {
	ID       int64  //
	RoleName string `xorm:"varchar(50)"`
	RoleInfo string `xorm:"varchar(200)"`  //
	Rights   string `xorm:"varchar(8000)"` //
	RoleType int    //
	State    int    //
	_godb    `xorm:"-"`
}

//初始化函数
func (this *TRoles) Iint() *TRoles {
	this._godb._bean = this
	this._godb._beans = make([]TRoles, 0)
	return this
}
