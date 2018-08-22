package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	UserId   int    `json:"user_id" bson:"b_user_id"`
	UserName string `json:"user_name" bson:"b_user_name"`
}

func main()  {

	u := &User{UserId: 1, UserName: "tony"}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))

	// 获取tag中的内容
	t := reflect.TypeOf(u)
	field := t.Elem().Field(0)
	fmt.Println(field.Tag.Get("json"))
	// 输出：user_id
	fmt.Println(field.Tag.Get("bson"))
}