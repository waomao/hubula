package bootstrap

import (
	"github.com/waomao/hubula/conf"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

//Configurator 定义配置器 类型是func
type Configurator func(*Bootstrapper)

//Bootstrapper 使用Go内建的嵌入机制(匿名嵌入)，允许类型之前共享代码和数据
// （Bootstrapper继承和共享 iris.Application ）
// 参考文章： https://hackthology.com/golangzhong-de-mian-xiang-dui-xiang-ji-cheng.html
type Bootstrapper struct {
	//内置继承iris类
	*iris.Application
	AppName  string
	AppOwner string
	//创建时间
	AppSpawnDate time.Time
}

// New returns a new Bootstrapper.
//实例化
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application:  iris.New(),
		AppSpawnDate: time.Now(),
		AppName:      appName,
		AppOwner:     appOwner,
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// SetupViews loads the templates.
//初始化模板 传入一个目录
func (b *Bootstrapper) SetupViews(viewsDir string) {
	htmlEngine := iris.HTML(viewsDir, ".html").Layout("shared/layout.html")
	// 每次重新加载模版（线上关闭它）调试时改模板随时生效
	htmlEngine.Reload(true)
	// 给模版内置各种定制的方法 时间转换的
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeformShort)
	})
	htmlEngine.AddFunc("FromUnixtime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeform)
	})
	//注册进去
	b.RegisterView(htmlEngine)
}

//异常处理
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		//拿到错误信息 放到网页上
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		//判断输出方式 json
		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			_, _ = ctx.JSON(err)
			return
		}
		//否则就用模板输出
		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		_ = ctx.View("shared/error.html")
	})
}

// Configure accepts configurations and runs them inside the Bootstraper's context.
//给web的配置方法
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// 启动计划任务服务
func (b *Bootstrapper) setupCron() {
	// 服务类应用

}

//定义两个常量 站点的对外目录
const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./public/"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)

// Bootstrap prepares our application.
//
// Returns itself.
//初始化
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	//模板目录
	b.SetupViews("./views")

	//异常信息
	b.SetupErrorHandlers()

	// static files
	//默认图标
	b.Favicon(StaticAssets + Favicon)
	//静态站点 把目录的.去掉
	b.HandleDir(StaticAssets[1:len(StaticAssets)-1], StaticAssets)
	// crontab
	//启动计划任务
	b.setupCron()
	// middleware, after static files
	//出异常
	b.Use(recover.New())
	//日志
	b.Use(logger.New())

	return b
}

// Listen starts the http server with the specified "addr".
//监听
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	_ = b.Run(iris.Addr(addr), cfgs...)
}
