package common

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/waomao/hubula/web/viewsmodels"
	"net"
	"net/http"
	"regexp"
	"strings"
)

//IsMobile 判断是否手机登录
func IsMobile(header *http.Request) (isMobile bool) {
	// header := c.Ctx.Request().Header

	//iris中获取headers
	//fmt.Printf("mobile header的类型%T\n", header)
	//headers := make(map[string]string)
	// for key, item := range header {
	// 	headers[key] = item[0]
	// 	//fmt.Printf("%s=%s\n", key, item[0])
	// }

	//net/http中获取headers
	//put headers in a map
	headers := make(map[string]string)

	if len(header.Header) > 0 {
		for k, v := range header.Header {
			headers[k] = v[0]
			//fmt.Printf("%s=%s\n", k, v[0])
		}
	}

	isMobile = false
	via := strings.ToLower(headers["VIA"])

	accept := strings.ToUpper(headers["Accept"])

	HTTP_X_WAP_PROFILE := headers["X_WAP_PROFILE"]

	HTTP_PROFILE := headers["PROFILE"]

	HTTP_USER_AGENT := headers["User-Agent"]

	if via != "" && strings.Index(via, "wap") != -1 {

		isMobile = true

	} else if accept != "" && strings.Index(accept, "VND.WAP.WML") != -1 {

		isMobile = true

	} else if HTTP_X_WAP_PROFILE != "" || HTTP_PROFILE != "" {

		isMobile = true

	} else if HTTP_USER_AGENT != "" {

		reg := regexp.MustCompile(`(?i:(blackberry|configuration\/cldc|hp |hp-|htc |htc_|htc-|iemobile|kindle|midp|mmp|motorola|mobile|nokia|opera mini|opera |Googlebot-Mobile|YahooSeeker\/M1A1-R2D2|android|iphone|ipod|mobi|palm|palmos|pocket|portalmmm|ppc;|smartphone|sonyericsson|sqh|spv|symbian|treo|up.browser|up.link|vodafone|windows ce|xda |xda_|MicroMessenger))`)

		fmt.Printf("%q\n", reg.FindAllString(HTTP_USER_AGENT, -1))

		if len(reg.FindAllString(HTTP_USER_AGENT, -1)) > 0 {

			isMobile = true

		}

	}

	return
}

//ClientIP 得到客户端IP地址
func ClientIP(request *http.Request) string {
	host, _, _ := net.SplitHostPort(request.RemoteAddr)
	return host
}

//Redirect 跳转URL
func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}


////////////////////////////////

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	Sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)


// encrypt password 加密密码
func PasswordSalt(pass, salt string) string {
	salt1 := "4%$@w"
	//密码干扰码
	password_salt := "123123123123213123c"
	str :=salt1+pass+salt+password_salt
	//return crypt.Md5(crypt.Sha256(str))
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
//验证
func PasswordVerify(password,  pass, salt string) bool {
	return password == PasswordSalt(pass, salt)
}

// 通用 session 填充
func SessionSet(c iris.Context , sessionname string , sessionvalue interface{}) {
	//存入 Session
	session := Sess.Start(c)
	session.Set(sessionname,sessionvalue)
}

// 用户登录 session 填充
func LoginSessionSet(c iris.Context , sessionname string , sessionvalue *viewsmodels.AdminSession) {
	//存入 Session
	str2,_ := json.Marshal(sessionvalue)
	session := Sess.Start(c)

	//用户名
	session.Set(sessionname, string(str2))
	//登录状态
	session.Set("ISLOGIN", true)

	//格式化显示
	//buff,_ := json.MarshalIndent(sessionvalue,"","   ")
	//fmt.Println("str => ?", string(buff))
}

//获取
func LoginSession(c iris.Context) (*viewsmodels.AdminSession) {
	session := Sess.Start(c)

	user := session.GetString("jname")
	//fmt.Println(usern)


	if user == "" {
		return nil
	}

	stu := &viewsmodels.AdminSession{}
	err := json.Unmarshal([]byte(user),&stu)
	if err != nil{
		return nil
	}

	return stu
}

//获取
func LoginSessionGet(c iris.Context , sessionname string) (bool,*viewsmodels.AdminSession, error) {
	session := Sess.Start(c)
	SESSION_NAME := sessionname

	login,err := session.GetBoolean("ISLOGIN")
	user := session.GetString(SESSION_NAME)
	//fmt.Println(usern)

	if err != nil {
		return false,nil, fmt.Errorf("Msg : Session 不存在")
	}

	if err == nil && user == "" {
		return false,nil, fmt.Errorf("Msg : Session 为空")
	}

	stu := &viewsmodels.AdminSession{}
	err = json.Unmarshal([]byte(user),&stu)
	if err != nil{
		return false,nil, fmt.Errorf("Msg : Session 序列号转换错误" + err.Error())
	}

	return login,stu, nil
}

//转换
//func Convert(admUser *models.Admin) *viewsmodels.AdminSession {
//	//赋值
//	Session := &viewsmodels.AdminSession{}
//	Session.Username = admUser.Username
//	Session.Aid = int(admUser.Aid)
//	Session.Mail = admUser.Mail
//	Session.TimeAdd = admUser.TimeAdd
//	Session.Ip = admUser.Ip
//	Session.NickName = admUser.NickName
//	Session.TrueName = admUser.TrueName
//	Session.Qq = admUser.Qq
//	Session.Phone = admUser.Phone
//	Session.Mobile = admUser.Mobile
//	return Session
//}

func loginNameHandler(ctx iris.Context){
	name := ctx.Params().Get("name")
	println(name)
	ctx.Next()
}

func loginHandler(ctx iris.Context){
	println("login")
	ctx.Next()
}

func before(ctx iris.Context){
	println("before")
	ctx.Next() //继续执行下一个handler，这本例中是mainHandler
}

func mainHandler(ctx iris.Context){
	println("mainHandler")
	ctx.Next()
}

func after(ctx iris.Context){
	println("after")
	ctx.Next()
}

