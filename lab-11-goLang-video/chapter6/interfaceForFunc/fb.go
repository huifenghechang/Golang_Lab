package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

// 为函数实现接口,将函数写成一个文件
func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()
	if next > 20000 {
		return 0, io.EOF
	}

	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

// 打印文件内容
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fibonacci()
	printFileContents(f)

}
