package conf

import (
	"github.com/kataras/iris/v12/middleware/logger"
	"time"
)

var (
	//IndexPath 前台模板路径
	MubanIndexPath = "fronted"

	//AdminPath 后台模板路径
	MubanAdminPath = "backend"

	//PcPath PC模板路径
	PcPath = "pc"

	//WapPath WAP模板路径
	WapPath = "wap"

	//AutoPath AUTO模板路径
	AutoPath = "auto"

	//RunningCrontabService 是否需要启动全局计划任务服务
	RunningCrontabService = false

	//SysTimeLocation 中国时区
	SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

	//SignSecret ObjSalesign 签名密钥
	SignSecret = []byte("0123456789abcdef")

	//CookieSecret cookie中的加密验证密钥
	CookieSecret = "waomao.com_login_276566565"

	// 请求日志记录
	CustomLogger = logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	})
)
