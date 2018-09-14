package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	// 新建并配置请求头及其相关函数
	request, err := http.NewRequest(
		http.MethodGet, "http://www.imooc.com", nil)

	request.Header.Add("User-Agent",
		"Mozilla/5.0(iphone;CPU iphone os 10_3 like Mac OS X) AppleWebkit/620.1")

	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	// 使用client来发送请求，并获取返回的响应Response
	resp, err := client.Do(request)
	if err != nil {
		panic(nil)
	}

	defer resp.Body.Close()

	// 使用httpUtil工具来解析response
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", s)

}
