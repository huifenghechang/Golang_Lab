package main

import (
"fmt"
"os"
)

func  main()  {
	target := "Hello-World"
	// os.Args 是一个切片参数
	if len(os.Args) >1 {
		target = os.Args[1]
	}
	fmt.Println(target)

}
