package common

import (
	"fmt"
	"os"
)

/*
   判断文件或文件夹是否存在 存在为true
   如果返回的错误为nil,说明文件或文件夹存在
   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
                    使用os.IsExist()判断为true,说明文件或文件夹存在
   如果返回的错误为其它类型,则不确定是否在存在
*/
func PathExists(path string) bool {
	//os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("common->path_file : line 20 error![%v]\n", err)
			return true
		}
		fmt.Printf("common->path_file : line 23 error![%v]\n", err)
		return false
	}
	return true
}

// 简化版  直接用os.IsNotExist(err)
func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func IsHuBuLaOSS() bool{
	_dir := "./HuBuLaOSS"
	ok := PathExists(_dir)
	return ok
}

//func DDd()  {
//	_dir := "./gzFiles2"
//	exist:= PathExists(_dir)
//
//	if exist {
//		fmt.Printf("has dir![%v]\n", _dir)
//	} else {
//		fmt.Printf("no dir![%v]\n", _dir)
//		// 创建文件夹
//		err := os.Mkdir(_dir, os.ModePerm)
//		if err != nil {
//			fmt.Printf("mkdir failed![%v]\n", err)
//		} else {
//			fmt.Printf("mkdir success!\n")
//		}
//	}
//}