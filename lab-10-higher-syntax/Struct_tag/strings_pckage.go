package main

import (
	"strings"
	"fmt"
)

func main()  {
	s1 := "qerror"
	if strings.Index(s1,"error") > -1{
		fmt.Printf("hello---%d",strings.Index(s1,"error"))
	}
}
