package common

import (
	"github.com/wangbin/jiebago"
	"regexp"
	"strings"
)

var Seg jiebago.Segmenter

func init() {
	_ = Seg.LoadDictionary("configs/dict.txt")
}

/*
函数名：delete_extra_space(s string) string
功  能:删除字符串中多余的空格(含tab)，有多个空格时，仅保留一个空格，同时将字符串中的tab换为空格
参  数:s string:原始字符串
返回值:string:删除多余空格后的字符串
创建时间:2018年12月3日
修订信息:
*/
func Delete_extra_space(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	//替换非utf8字符为空格
	s1 := strings.ToValidUTF8(s," ")
	//替换tab为空格
	s1 = strings.Replace(s1, "	", " ", -1)
	//替换换行为空格
	s1 = strings.Replace(s1,"\n"," ",-1)
	//替换|为空格
	s1 = strings.Replace(s1,"|"," ",-1)

	//替换|为空格
	s1 = strings.Replace(s1,"/"," ",-1)
	//替换|为空格
	s1 = strings.Replace(s1,"\""," ",-1)
	//替换|为空格
	s1 = strings.Replace(s1,"("," ",-1)
	//替换|为空格
	s1 = strings.Replace(s1,")"," ",-1)

	//，空格
	s1 = strings.Replace(s1,","," ",-1)
	//。空格
	s1 = strings.Replace(s1,"."," ",-1)
	//,空格
	s1 = strings.Replace(s1,"，"," ",-1)
	//.空格
	s1 = strings.Replace(s1,"。"," ",-1)

	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}

//字符串【搜索引擎模式】分词
func PrintKeyWord(ch <-chan string) string{
	var str  = ""
	for word := range ch {
		str += "|" + word
		//fmt.Printf("|%s", word)
	}
	return str
}

//博客文章关键检索
func BlogKeyWord(s string) string{
	cachestr := s
	//去除多余空格 只保留一个空格
	s = Delete_extra_space(s)

	//按空格分割
	a := strings.Fields(s)
	str := cachestr

	if len(a) > 0 {
		for i:=0;i<len(a) ;i++  {
			//fmt.Println(a[i])
			//str += "|" + a[i]
			str += PrintKeyWord(Seg.CutForSearch(a[i], true))
		}
	}
	return str
}