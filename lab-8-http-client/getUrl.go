package main

import (
	"net/http"
	"log"
	"fmt"
	"reflect"
	"bytes"
)

func main()  {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil{
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	headers := resp.Header

	for k,v := range headers{
		fmt.Printf("k=%v, v=%v", k, v)
	}

	fmt.Println()
	fmt.Printf("resp status %s,statusCode %d \n",resp.Status,resp.StatusCode)
	fmt.Printf("resp Proto %s \n",resp.Proto) // proto 协议
	fmt.Printf("resp content length %d \n",resp.ContentLength)
	fmt.Printf("resp transfer encoding %v \n",resp.TransferEncoding)
	fmt.Printf("resp Uncompresed %t \n", resp.Uncompressed)

	fmt.Println(reflect.TypeOf(resp.Body))

	buf := bytes.NewBuffer(make([]byte,0,512))

	length, _ := buf.ReadFrom(resp.Body)
	fmt.Println(len(buf.Bytes()))
	fmt.Println(length)
	//fmt.Println(string(buf.Bytes()))



}

