package viewsmodels

import (
	"time"
)

//后台session
type AdminSession struct {
	Aid        int       `orm:"column(aid);auto"`
	Username   string    `orm:"column(username);size(30)"`
	Mail       string                        `orm:"column(mail);size(80)"`
	TimeAdd    time.Time                     `orm:"column(time_add);type(timestamp);null;auto_now_add"`
	TimeUpdate time.Time                     `orm:"column(time_update);type(timestamp);null"`
	Ip         string                        `orm:"column(ip);size(15)"`
	JobNo      string                        `orm:"column(job_no);size(15)"`
	NickName   string                        `orm:"column(nick_name);size(50)"`
	TrueName   string                        `orm:"column(true_name);size(50)"`
	Qq         string                        `orm:"column(qq);size(50)"`
	Phone      string                        `orm:"column(phone);size(50)"`
	Mobile     string                        `orm:"column(mobile);size(20)"`
	//Role_id    map[int]models.AdminRoleAccess //扩展角色
}
