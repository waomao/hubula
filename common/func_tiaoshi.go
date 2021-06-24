package common

import (
	"fmt"
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

//// 从cookie中得到当前登录的用户
//func GetLoginUser(request *http.Request) *models.ObjLoginuser {
//	//获取cookie lottery_loginuser
//	c, err := request.Cookie("waomao_loginuser")
//	if err != nil {
//		return nil
//	}
//	//用url格式解析cookie
//	params, err := url.ParseQuery(c.Value)
//	if err != nil {
//		return nil
//	}
//	//转换uid
//	uid, err := strconv.Atoi(params.Get("uid"))
//	if err != nil || uid < 1 {
//		return nil
//	}
//	// Cookie最长使用时长
//	now, err := strconv.Atoi(params.Get("now"))
//	//当前时间 - cookie 时间 大于 30天 cookie失效
//	if err != nil || NowUnix()-now > 86400*30 {
//		return nil
//	}
//	//// IP修改了是不是要重新登录
//	//ip := params.Get("ip")
//	//if ip != ClientIP(request) {
//	//	return nil
//	//}
//	// 登录信息
//	loginuser := &models.ObjLoginuser{}
//	loginuser.Uid = uid
//	loginuser.Username = params.Get("username")
//	loginuser.Now = now
//	loginuser.Ip = ClientIP(request)
//	//签名
//	loginuser.Sign = params.Get("sign")
//	// if err != nil {
//	// 	log.Println("fuc_web GetLoginUser Unmarshal ", err)
//	// 	return nil
//	// }
//	//验证签名
//	sign := createLoginuserSign(loginuser)
//	if sign != loginuser.Sign {
//		log.Println("fuc_web GetLoginUser createLoginuserSign not sign", sign, loginuser.Sign)
//		return nil
//	}
//
//	return loginuser
//}

//// 将登录的用户信息设置到cookie中
//func SetLoginuser(writer http.ResponseWriter, loginuser *models.ObjLoginuser) {
//	//如果不存在或id小于1 清理cookie
//	if loginuser == nil || loginuser.Uid < 1 {
//		c := &http.Cookie{
//			Name:  "waomao_loginuser",
//			Value: "",
//			Path:  "/",
//			//过期
//			MaxAge: -1,
//			Domain: "paodj.com",
//		}
//		http.SetCookie(writer, c)
//		return
//	}
//	//如果签名为空 则生成签名
//	if loginuser.Sign == "" {
//		loginuser.Sign = createLoginuserSign(loginuser)
//	}
//	//构造一个参数
//	params := url.Values{}
//	//构造好写入cookie
//	params.Add("uid", strconv.Itoa(loginuser.Uid))
//	params.Add("username", loginuser.Username)
//	params.Add("now", strconv.Itoa(loginuser.Now))
//	params.Add("ip", loginuser.Ip)
//	params.Add("sign", loginuser.Sign)
//	c := &http.Cookie{
//		Name:   "waomao_loginuser",
//		Value:  params.Encode(),
//		Path:   "/",
//		Domain: "paodj.com",
//	}
//	http.SetCookie(writer, c)
//}
//
//// 根据登录用户信息生成加密字符串 签名
//func createLoginuserSign(loginuser *models.ObjLoginuser) string {
//	str := fmt.Sprintf("uid=%d&username=%s&secret=%s&now=%d", loginuser.Uid, loginuser.Username, conf.CookieSecret, loginuser.Now)
//	//fmt.Println(str)
//	//md5
//	sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))
//	//fmt.Println(sign)
//	return sign
//	//return CreateSign(str)
//}
