package main

import (
	"fmt"
	"net/http"
)

//文件服务：提供文件的访问，提供下载，上传，
func main() {
	//指定文件路径，创建服务器
	dirpath := http.FileServer(http.Dir("D:\\GoPATH\\src\\FileServerDemo\\file"))

	//指定访问路由
	http.Handle("/", dirpath)
	//http.Handle("/",http.StripPrefix("/test/",dirpath))

	/*http.HandleFunc() 函数类型
	http.Header() 接口类型*/

	//设置端口号，启动服务
	err := http.ListenAndServe(":9091", dirpath)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("创建服务器失败")
		return
	}
}
