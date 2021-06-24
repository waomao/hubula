package bootstrap

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/waomao/hubula/common"
	"github.com/waomao/hubula/conf"
	"time"
)

//Configurator 定义配置器 类型是func
type Configurator func(bootstrapper *Bootstrapper)

//Bootstrapper 使用Go内建的嵌入机制(匿名嵌入)，允许类型之前共享代码和数据
// （Bootstrapper继承和共享 iris.Application ）
// 参考文章： https://hackthology.com/golangzhong-de-mian-xiang-dui-xiang-ji-cheng.html
type Bootstrapper struct {
	//内置继承iris类
	*iris.Application

	//两个基本的标识信息
	AppName  string
	AppOwner string
	//创建时间
	AppSpawnDate time.Time
}

// New returns a new Bootstrapper.
//实例化 cfgs ...Configurator更多参数
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application:  iris.New(),
		AppSpawnDate: time.Now(),
		AppName:      appName,
		AppOwner:     appOwner,
	}

	//更多配置是一个切片。循环取出
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
	htmlEngine.Reload(conf.Configs().Stage[conf.Configs().Runmode].Admin_load)

	// 设置页面的函数
	htmlEngine.AddFunc("greet", func(s string) string {
		return "Greetings, " + s + "!"
	})
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
		//fmt.Println(ctx.GetStatusCode())
		//拿到错误信息 放到网页上
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"code":  ctx.Values().GetStringDefault("code","404"),
			"message": ctx.Values().GetStringDefault("message","访问出错啦 : ("),
			"trace": ctx.Values().GetStringDefault("trace","返回首页吧"),
		}

		//这里我们不考虑cookie和自适应，只考虑设备类型
		pathError := "shared/error.html"
		if common.Mobile(ctx.Request()) == true {
			pathError = "shared/error_wap.html"
		}

		//判断输出方式 json
		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			_, _ = ctx.JSON(err)
			return
		}

		//fmt.Println(err)
		//否则就用模板输出
		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		_ = ctx.View(pathError)
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

// Bootstrap prepares our application.
// Returns itself.
//初始化
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	//模板目录
	b.SetupViews(conf.Configs().Webvar.Views)

	//异b常信息
	b.SetupErrorHandlers()

	//默认图标
	b.Favicon(conf.StaticPublic + conf.Favicon)

	//静态站点 把目录的.去掉
	b.HandleDir(conf.StaticAssets[1:len(conf.StaticAssets)-1], conf.StaticAssets)
	b.HandleDir(conf.StaticPublic[1:len(conf.StaticPublic)-1], conf.StaticPublic)

	//后台
	b.HandleDir("/adpcss/", "./assets/pc/css/")
	b.HandleDir("/adpimg/", "./assets/pc/img/")
	b.HandleDir("/adpjs/", "./assets/pc/js/")
	b.HandleDir("/adplib/", "./assets/pc/lib/")

	b.HandleDir("/adwcss/", "./assets/wap/css/")
	b.HandleDir("/adwimg/", "./assets/wap/img/")
	b.HandleDir("/adwjs/", "./assets/wap/js/")
	b.HandleDir("/adwlib/", "./assets/wap/lib/")

	b.HandleDir("/adacss/", "./assets/auto/css/")
	b.HandleDir("/adaimg/", "./assets/auto/img/")
	b.HandleDir("/adajs/", "./assets/auto/js/")
	b.HandleDir("/adalib/", "./assets/auto/lib/")

	//前台
	b.HandleDir("/inpcss/", "./public/pc/css/")
	b.HandleDir("/inpimg/", "./public/pc/img/")
	b.HandleDir("/inpjs/", "./public/pc/js/")
	b.HandleDir("/inplib/", "./public/pc/lib/")

	b.HandleDir("/inwcss/", "./public/wap/css/")
	b.HandleDir("/inwimg/", "./public/wap/img/")
	b.HandleDir("/inwjs/", "./public/wap/js/")
	b.HandleDir("/inwlib/", "./public/wap/lib/")

	b.HandleDir("/inacss/", "./public/auto/css/")
	b.HandleDir("/inaimg/", "./public/auto/img/")
	b.HandleDir("/inajs/", "./public/auto/js/")
	b.HandleDir("/inalib/", "./public/auto/lib/")


	//启动计划任务
	b.setupCron()

	//b.Logger().SetLevel("debug")
	//出异常
	b.Use(recover.New())
	//日志
	b.Use(logger.New())
	// 请求日志记录
	if conf.Configs().CustomLogger {
		b.Use(conf.CustomLogger)
	}

	//中间件
	//注意 Use 和 Done 方法需要写在绑定访问路径的方法之前
	//使用 ‘Use’ 方法作为当前域名下所有路由的第一个处理函数
	//而使用 ‘UseGlobal’ 方法注册的中间件，会在包括所有子域名在内的所有路由中执行
	//b.Use(UseBefore)
	//b.Done(DoneAfter)
	return b
}

// Listen starts the http server with the specified "addr".
//监听
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	_ = b.Run(iris.Addr(addr), cfgs...)
}

//func UseBefore(ctx iris.Context){
//	println("use - before")
//	ctx.Next()
//}
//
//func DoneAfter(ctx iris.Context) {
//	println("Done - after")
//	ctx.Next()
//}