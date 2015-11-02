package models

type TMenus struct {
	ID          int64  //
	NID         int    //
	ParentID    int    //
	Title       string `xorm:"varchar(50)"`  //
	Label       string `xorm:"varchar(50)"`  //
	Description string `xorm:"varchar(50)"`  //
	URL         string `xorm:"varchar(500)"` //
	Roles       string `xorm:"varchar(50)"`  //
	Orders      int    //	4)"` //
	Target      string `xorm:"varchar(50)"`  //
	Memo        string `xorm:"varchar(200)"` //
	Images      string `xorm:"varchar(100)"` //
	State       int    //	4)"` //
	_godb       `xorm:"-"`
}

//初始化参数
func (this *TMenus) Init() *TMenus {
	this._godb._bean = this
	this._godb._beans = make([]TMenus, 0)
	return this
}
