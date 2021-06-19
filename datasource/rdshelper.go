package datasource

import (
	"fmt"
	"github.com/waomao/hubula/conf"
	"log"
	"sync"
	"time"
	"github.com/gomodule/redigo/redis"
)

//互斥锁
var rdsLock sync.Mutex

//实例对象
var cacheInstance *RedisConn

//RedisConn 封装成一个redis资源池 连接池
type RedisConn struct {
	//连接池
	pool *redis.Pool
	//是否需要debug 自己用的
	showDebug bool
}

//Do 对外只有一个命令，封装了一个redis的命令
func (rds *RedisConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	//从连接池拿到连接
	conn := rds.pool.Get()
	//放回连接池
	defer conn.Close()

	//因为有debug功能 需要一个时间的检查
	t1 := time.Now().UnixNano()
	//执行 效率
	reply, err = conn.Do(commandName, args...)
	if err != nil {
		e := conn.Err()
		if e != nil {
			log.Println("rdshelper Do", err, e)
		}
	}
	//执行完的时间
	t2 := time.Now().UnixNano()
	if rds.showDebug {
		//打印执行时间
		fmt.Printf("[redis] [info] [%dus]cmd=%s, err=%s, args=%v, reply=%s\n", (t2-t1)/1000, commandName, err, args, reply)
	}
	return reply, err
}

//ShowDebug 设置是否打印操作日志
//前面用到名字是小写的这里要大写实例它
func (rds *RedisConn) ShowDebug(b bool) {
	rds.showDebug = b
}

//InstanceCache 得到唯一的redis缓存实例
func InstanceCache() *RedisConn {
	if cacheInstance != nil {
		return cacheInstance
	}
	rdsLock.Lock()
	defer rdsLock.Unlock()

	if cacheInstance != nil {
		return cacheInstance
	}
	return NewCache()
}

//NewCache 重新实例化
func NewCache() *RedisConn {
	//连接池
	pool := redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", conf.Configs().Redis["slave"].Host, conf.Configs().Redis["slave"].Port))
			if err != nil {
				log.Fatal("rdshelper.NewCache Dial error ", err)
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		//最多连接数
		MaxIdle: 10000,
		//最大活跃数
		MaxActive: 10000,
		//超时时间
		IdleTimeout: 0,
		//等待
		Wait: false,
		//连接的活跃时间 一直活跃
		MaxConnLifetime: 0,
	}
	//实例化连接对象
	instance := &RedisConn{
		pool: &pool,
	}

	cacheInstance = instance
	//设置debug
	cacheInstance.ShowDebug(true)
	//cacheInstance.ShowDebug(false)
	return cacheInstance
}
