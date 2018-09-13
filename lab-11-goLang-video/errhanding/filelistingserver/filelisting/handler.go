package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

// 代码的编写，采用了解耦合的设计。HandleFileList()用来处理客户端的请求
func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	// 检查请求链接是否含有特定的前缀
	fmt.Println()
	if strings.Index(
		request.URL.Path, prefix) != 0 {
		return userError(
			fmt.Sprintf("path %s must start "+
				"with %s", request.URL.Path, prefix))
	}

	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// 读取文件，将文件写入writer,最后返回给用户
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
