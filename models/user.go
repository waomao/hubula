package models

//User 结构体
type User struct {
	//imooc:"id" 自己设置的用于表单
	ID       int    `xorm: "id"json:"id" sql:"ID" imooc:"id"`
	Username string `xorm: "Username"json:"id" sql:"Username" imooc:"id"`
	password string `xorm: "password"json:"id" sql:"password" imooc:"id"`
	email    string `xorm: "email"json:"id" sql:"email" imooc:"id"`
}