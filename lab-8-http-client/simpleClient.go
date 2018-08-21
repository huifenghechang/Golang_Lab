package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("The first show is getUrl \n")
	getUrl()
	fmt.Printf("Next show is postUrl \n")
	postUrl()

}
func getUrl() string {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// 错误处理
		return "get Url Failed"
	}

	defer resp.Body.Close() //关闭链接

	headers := resp.Header
	for k, v := range headers {
		fmt.Printf("k=%v, v= %v\n", k, v)
	}

	fmt.Printf("resp status %s, statusCode %d \n", resp.Status, resp.StatusCode)
	fmt.Printf("resp Proto %s \n", resp.Proto) // proto 协议
	fmt.Printf("resp content length %d \n", resp.ContentLength)

	return "get Url Succeed!"
}

func postUrl() string {
	fmt.Printf("Hello-World!This is BlockChain World!")
	return "Post Succeed!"
}
