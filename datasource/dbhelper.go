package datasource
//创建数据库相关的对象 实例
import (
	"fmt"
	"github.com/waomao/hubula/conf"
	"log"
	"sync"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//dbLock 互斥锁
var dbLock sync.Mutex

//整个系统里面 一直保持一个
var masterInstance *xorm.Engine
var slaveInstance *xorm.Engine

//InstanceDbMaster 单例的模式 得到唯一的主库实例
//在整个程序会不断的调用数据库操作 不希望每次都实例化一个
func InstanceDbMaster() *xorm.Engine {
	//如果已经连接 就返回
	if masterInstance != nil {
		return masterInstance
	}
	//如果存在并发
	//加锁 创建之前锁定
	dbLock.Lock()
	//解锁
	defer dbLock.Unlock()
	//也许存在排队的可能
	if masterInstance != nil {
		return masterInstance
	}
	//创建连接
	return NewDbMaster()
}

//NewDbMaster 实例化xorm数据库的操作引擎 这个方法是每一次都会实例化一个
func NewDbMaster() *xorm.Engine {
	sourcename := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database)
	//连接数据库 得到一个实例
	instance, err := xorm.NewEngine(conf.DriverName, sourcename)
	if err != nil {
		log.Fatal("dbhelper.InstanceDbMaster NewEngine error ", err)
		return nil
	}
	//调试用的。展示每一条调试语句调试时间
	instance.ShowSQL(true)
	//instance.ShowSQL(false)
	//返回实例
	masterInstance = instance
	return masterInstance
}
