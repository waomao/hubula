package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/golog"
	"github.com/waomao/hubula/conf"
	"github.com/waomao/hubula/models"
	"github.com/xormplus/xorm"
	"log"
	"sync"
	"xorm.io/core"
)

//dbLock 互斥锁
var dbLock sync.Mutex
var (
	//保持一个
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
)

// 获取数据库连接的url
// true：master主库
func GetConnURL(mors string) (url string) {
	//"mysql", "user:password@/dbname
	// ?charset=utf8&parseTime=True&loc=Local"
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		conf.Configs().DB[mors].User,
		conf.Configs().DB[mors].Password,
		conf.Configs().DB[mors].Host,
		conf.Configs().DB[mors].Port,
		conf.Configs().DB[mors].Database,
		conf.Configs().DB[mors].Charset,
		conf.Configs().DB[mors].ParseTime,
		conf.Configs().DB[mors].Loc)
	//golog.Infof("@@@ DB conn==>> %s", url)
	return
}

//InstanceDbMaster 单例的模式 得到唯一的主库实例
// 主库，单例 InstanceDbMaster MasterEngine

func MasterEngine() *xorm.Engine {
	//如果已经连接 就返回
	if masterEngine != nil {
		return masterEngine
	}
	//如果存在并发
	//加锁 创建之前锁定
	dbLock.Lock()
	//解锁
	defer dbLock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	//连接数据库
	engine, err := xorm.NewEngine(conf.Configs().DB["master"].DriverName, GetConnURL("master"))
	if err != nil {
		golog.Fatalf("@@@ Instance Master DB error!! %s", err)
		log.Fatal("dbhelper.InstanceDbMaster NewEngine error ", err)
		return nil
	}

	settings(engine, "master")

	engine.SetMapper(core.GonicMapper{})

	//返回实例
	masterEngine = engine
	return masterEngine
}

// 从库，单例
func SlaveEngine() *xorm.Engine {
	//如果已经连接 就返回
	if slaveEngine != nil {
		return slaveEngine
	}
	//如果存在并发
	//加锁 创建之前锁定
	dbLock.Lock()
	//解锁
	defer dbLock.Unlock()
	//也许存在排队的可能
	if slaveEngine != nil {
		return slaveEngine
	}

	//连接数据库
	engine, err := xorm.NewEngine(conf.Configs().DB["slave"].DriverName, GetConnURL("slave"))
	if err != nil {
		log.Fatal("dbhelper.InstanceDbMaster NewEngine error ", err)
		return nil
	}

	settings(engine, "slave")

	//返回实例
	slaveEngine = engine
	return engine
}

//数据库连接后的基本设置
func settings(engine *xorm.Engine, mors string) {
	//打印连接信息
	if err := engine.Ping(); err != nil {
		fmt.Println(err)
	}

	//xorm reverse mysql root:root@tcp(127.0.0.1:3306)/hubula?charset=utf8mb4 templates/goxorm
	//同步创建数据表
	err := engine.Sync2(
		new(models.ADemo))

	if err != nil {
		panic(err.Error())
	}

	//调试用的。展示每一条调试语句调试时间
	engine.ShowSQL(conf.Configs().DB[mors].ShowSql)
	//engine.ShowSQL(false)

	//时区
	engine.SetTZLocation(conf.SysTimeLocation)

	if conf.Configs().DB[mors].MaxIdleConns > 0 {
		engine.SetMaxIdleConns(conf.Configs().DB[mors].MaxIdleConns)
	}
	if conf.Configs().DB[mors].MaxOpenConns > 0 {
		engine.SetMaxOpenConns(conf.Configs().DB[mors].MaxOpenConns)
	}

	// 性能优化的时候才考虑，加上本机的SQL缓存
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//engine.SetDefaultCacher(cacher)
}