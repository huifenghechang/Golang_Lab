package main

import (
	"fmt"
	"lab-11-goLang-video/chapter5/mock"
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
	var r Retriever
	mockRetriever := mock.Retriever{
		Contents: "this is a fake imooc.com"}

	r = &mockRetriever

	fmt.Println(r.Get(url))

}
