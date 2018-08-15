package main

// Go 语言的字符，采用的是UTF8编码格式
import (
	"fmt"
	"strings"
	"strconv"
)

func main()  {
	t1 := "\"Hello World\""
	t2 := `"hello Go Language"`
	fmt.Println(t1)
	fmt.Println(t2)

	// string 包，主要是用来查找、替换、分割字符串等操作
	var str string = "go_lang"
	fmt.Println("str is -->",str)
	var exit = strings.Contains(str,"go")
	fmt.Println("Contain go ?",exit)
	str_new :=strings.Replace(str,"go","java",1)
	fmt.Println(str_new)

	// Strconv包，主要用于字符串和其他类型的数据之间进行转换
	var ori string = "123456"
	var i int
	var s string
	i, _ = strconv.Atoi(ori)
	i += 1000000
	s = strconv.Itoa(i)
	fmt.Println(s)

	//
}
