package main

import (
	"fmt"
	"lab-11-goLang-video/chapter7/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

// 服务器列表读取文件：
/*
	- 基本步骤“读取路径-打开文件-读取文件-写入返回头-关闭文件-！”
*/

// 定义一个类型，用于 handler(writer, request) 这个Handler，实际上就是http.HandleFunc（）需要传入的参数！
type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(
		writer http.ResponseWriter, request *http.Request) {

		//创建defer-recover语句,用于捕获异常
		defer func() {

			fmt.Println("I am in errWrapper ! -- -- !! -- ")

			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
			}
			http.Error(writer,
				http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}()

		err := handler(writer, request) // 这里调用的是fileListingServer里面的部分
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
	fmt.Println("I am first -- -- !! -- ")

	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		fmt.Println("I am main function in web before panic(err)")
		panic(err)
	}

}
