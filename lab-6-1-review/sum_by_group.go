package main

import "fmt"

func Sum(a []int, result chan int){
	sum := 0
	for _, v := range a{
		sum += v
	}
	result <- sum
}

func main()  {
	a := []int{1,2,3,4,5,6,7,8,9}
	result := make(chan int)
	go Sum(a[:len(a)/2],result)
	go Sum(a[len(a)/2:],result)
	x, y := <- result, <- result
	fmt.Println("sum is",x,"+",y,x+y)


}