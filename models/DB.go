package models

import (
	"fmt"
	"strings"
)

type _godb struct {
	_bean  interface{}
	_beans interface{}
}

//插入一条记录,返回值为影响行数,实体的主键将获取新值
func (this *_godb) Insert() int64 {
	if this._bean == nil {
		fmt.Println("")
		return 0
	}
	i, err := X.Insert(this._bean)
	if err != nil {
		fmt.Println("_godb insert error:the _bean is nil!")
		return 0
	}
	return i
}

//根据ID主键删除一条记录
func (this *_godb) Del(id int64) int64 {
	if this._bean == nil {
		fmt.Println("_godb del error:the _bean is nil!")
		return 0
	}
	rows, err := X.Id(id).Delete(this._bean)
	if err != nil {
		fmt.Println("_godb del error:", err.Error())
		return 0
	}
	return rows
}

//根据主键ID返回数据到调用实体
func (this *_godb) Get(id int64) bool {
	if this._bean == nil {
		fmt.Println("_godb get error:the _bean is nil!")
		return false
	}
	rst, err := X.Id(id).Get(this._bean)
	if err != nil {
		fmt.Println("_godb get error:", err.Error())
		return false
	}
	return rst
}

//根据ID来修改一条记录  0失败 1成功
func (this *_godb) Update(id int64) int64 {
	if this._bean == nil {
		fmt.Println("_godb update error:the _bean is nil!")
		return 0
	}
	rst, err := X.Id(id).Update(this._bean)
	if err != nil {
		fmt.Println("_godb update error:", err.Error())
		return 0
	}
	fmt.Println("_godb update rst :%v 更新结果: %v", this._bean, rst)
	return 1
}

//SQL语句查询
func (this *_godb) Query(sql string) []map[string]string {
	rsts, err := X.Query(sql)
	if err != nil {
		fmt.Println("_godb query error:", err.Error())
		return nil
	}
	rst := ParseByte(rsts)

	return rst
}

//根据条件查询数据
func (this *_godb) Single(sql string) map[string]string {
	rsts, err := X.Query(sql)
	if err != nil {
		fmt.Println("_godb query error:", err.Error())
		return nil
	}
	rst := ParseByte(rsts)
	if len(rst) > 0 {
		return rst[0]
	}
	return nil
}

//根据条件查询数据
func (this *_godb) Find(where string) interface{} {
	if this._beans == nil {
		fmt.Println("_godb find error:the _beans is nil!")
		return nil
	}
	if where != "" {
		err := X.Where(where).Find(this._beans)
		if err != nil {
			fmt.Println("_godb find error1:", err.Error())
			return nil
		}
	} else {
		err := X.Find(this._beans)
		if err != nil {
			fmt.Println("_godb find error2:", err.Error())
			return nil
		}
	}
	return this._beans
}

//字典字节类型转换
func ParseByte(beans []map[string][]byte) []map[string]string {
	_st := make([]map[string]string, 0)
	for _, value := range beans {
		_mt := make(map[string]string)
		for k, v := range value {
			_mt[strings.ToLower(k)] = string(v)
		}
		_st = append(_st, _mt)
	}
	return _st
}
