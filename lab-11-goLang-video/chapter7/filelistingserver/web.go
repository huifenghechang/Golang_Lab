package main

import (
	"lab-11-goLang-video/chapter7/filelistingserver/filelisting"
	"net/http"
	"os"
)

// 服务器列表读取文件：
/*
	- 基本步骤“读取路径-打开文件-读取文件-写入返回头-关闭文件-！”
*/

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(
		writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code),
				code)
		}

	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		panic(err)
	}

}
