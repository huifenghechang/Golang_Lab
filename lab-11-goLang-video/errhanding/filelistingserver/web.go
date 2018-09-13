package main

import (
	"lab-11-goLang-video/errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 捕获异常panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		// 处理来自客户端的请求
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error occurred"+
				"handling request: %s",
				err.Error())
		}

		// 系统错误处理
		code := http.StatusOK
		switch {
		case os.IsNotExist(err):
			code = http.StatusNotFound
		case os.IsPermission(err):
			code = http.StatusForbidden
		default:
			code = http.StatusInternalServerError
		}
		http.Error(writer,
			http.StatusText(code), code)
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
