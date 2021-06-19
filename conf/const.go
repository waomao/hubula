package conf

const (
	//站点的对外目录
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./assets/"
	StaticPublic = "./public/"

	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"

	//SysTimeform 时间
	SysTimeform = "2006-01-02 15:04:05"

	//SysTimeformShort 时间
	SysTimeformShort = "2006-01-02"
	////////////////////TYPE/////////////////////
	//站点配置ID
	SITE_ID      = 10007 //站点配置ID
	ADMIN_ROLE   = 420   //角色ID
	MEMBER_GROUP = 400   //用户组

	//博客模块ID
	TYPE_ID = 10006 //博客模块ID
	//原创
	ORIGINAL = 10003 //原创
	//栏目 博客分类属性 栏目ID
	TYPE_CAT = 10001 //栏目 博客分类属性 栏目ID
	//文章
	TYPE_ARTICLE = 0 //文章
	//是否阅读
	READ_FINISH = 10016 //已看
	READ_NOW    = 10015 //在看
	READ_NOT    = 10014 //未看
	///////////////////////////////////////////
	APP_CSDN = 10011 //csdn
	//////////
	ADMIN_YES = 301 //后台
	ADMIN_NO  = 302 //前台
	//
	MODULE_ID_WORK_OTHER = 10018 //技术之外文章 生活
	MODULE_ID_WORK       = 10019 //技术
	//
	TYPE_FROM           = 600 //表单标签
	TYPE_INPUT_RADIO    = 603 //单选框
	TYPE_INPUT_TEXTAREA = 604 //多行文本框
	//
	APP_API = 10023 //接口
)

