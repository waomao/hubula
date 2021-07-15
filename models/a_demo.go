package models

import (
	"time"
)

type ADemo struct {
	Id          int64     `xorm:"not null pk autoincr comment('记录标识 pk') BIGINT(20)"`
	ToId          int64     `xorm:"not null comment('所属组织 fk') BIGINT(20)"`
	LoginName     string    `xorm:"not null comment('登录帐号') VARCHAR(64)"`
	Password      string    `xorm:"not null comment('用户密码') VARCHAR(64)"`
	Vsername      string    `xorm:"not null comment('用户姓名') VARCHAR(64)"`
	Mobile        string    `xorm:"comment('手机号') VARCHAR(20)"`
	Email         string    `xorm:"comment('电子邮箱') VARCHAR(64)"`
	GenTime       time.Time `xorm:"not null comment('创建时间') DATETIME"`
	LoginTime     time.Time `xorm:"comment('登录时间') DATETIME"`
	LastLoginTime time.Time `xorm:"comment('上次登录时间') DATETIME"`
	Count         int64     `xorm:"not null comment('登录次数') BIGINT(20)"`
	IsDel      int64       `xorm:"not null default 0 comment('删除0否1是') BIGINT(20)"`
	SuperiorId int           `xorm:"not null default 0 comment('上级id 默认0') INT(5)"`
	Power      int           `xorm:"not null comment('操作权限') INT(2)"`
	Children   []*ADemo `xorm:"-"`
}