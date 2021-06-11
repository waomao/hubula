package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/waomao/hubula/models"
)

//UserDao 结构体在这里相当于是一个类的概念
type UserDao struct {
	//数据库相关的操作 xorm引擎
	engine *xorm.Engine
}

//NewUserDao 实例化公共方法
func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

//Get id 返回模型
func (d *UserDao) Get(id int) *models.User {
	data := &models.User{ID: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}

	data.ID = 0
	return data
}
