package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

func main(){
	resp, err := http.Post("http://www.baidu.com",
		"application/x-www-form-urlencoded",
			strings.NewReader("username=xxx&password=hello"))
	if err != nil{
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
