package main

import (
	"bufio"
	"errors"
	"fmt"
	"lab-11-goLang-video/chapter7/fib"
	"os"
)

/* 课程大纲：
- func tryDefer    ----- defer 的使用  defer 是堆栈调用函数，先进后出
- func writeFile() ---- 资源管理

*/

// 参数在defer语句时计算，最后打印的是30-29-、、、-0 而不是打印 30 行 30 出来！！！
func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

// 出错处理代码！！！
func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)

	err = errors.New("This is a new error")

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()
	// 使用bufio来写文件，速度比较快，先写到内存，最后一起导入
	writer := bufio.NewWriter(file)
	// 将写入的数据导入到filename文件中
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")

}
