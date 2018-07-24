package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main(){
	var str string  = "go_lang"
	fmt.Printf("T/F? Does the string \"%s\"Contains \"%s\"?\n",str,"go")
	fmt.Printf("%t\n",strings.Contains(str,"go"))

	var ori string = "123456"
	var i int
	var s string

	fmt.Printf("The size of ints is:%d\n",strconv.IntSize)

	i, _ = strconv.Atoi(ori)
	fmt.Printf("The interger is : %d\n",i)
	i = i+5
	s = strconv.Itoa(i)
	fmt.Printf("The new value of string is %s",s)

}
