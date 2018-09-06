package main

import (
	"fmt"
	"lab-11-goLang-video/chapter5/mock"
	"lab-11-goLang-video/chapter5/real"
	"time"
)

// 定义三个接口
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "http://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "cmouse",
			"course": "golang",
		})
}

// 接口之间的组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another faked imooc.com"})
	return s.Get(url)
}

func main() {
	// Retriever 是一个接口
	var r1 Retriever

	mockRetriever := mock.Retriever{
		Contents: "this is a fake imooc.com"}

	// 因为接口，本身是一个指针，所给给接口赋值的时候，需要将结构体转换为地址形式
	r1 = &mockRetriever
	fmt.Println(r1.Get(url))

	var r2 Retriever

	r2 = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	inspect(r2)

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> Type:%T Value:%v \n", r, r)
	fmt.Printf("> Tpye Switch")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("Contents", v.UserAgent)
	}

	fmt.Println()
}
