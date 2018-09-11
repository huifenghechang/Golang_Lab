package main

import (
	"fmt"
	"os"
)

/* 课程大纲：
- func tryDefer    ----- defer 的使用  defer 是堆栈调用函数，先进后出
- func writeFile() ---- 资源管理

*/

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
	}
}

func writeFile(fileName string) {
	file, err := os.Open(fileName,
		os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {

	}

}

func main() {

}
