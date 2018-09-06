package main

import "fmt"

// 函数式编程，参数和返回值都可以是函数！函数是第一等公民。
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// 将函数定义为iAdder类型
type iAdder func(int) (int, iAdder)

// 正统函数式编程，没有变量，只有函数和常量
func add2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, add2(base + v)

	}
}

func main() {
	a := adder()
	b := add2(10)
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d = %d\n", i, a(i))
	}

	fmt.Println("")

	for i := 10; i < 15; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("10+...+ %d = %d \n", i, s)
	}

}
