package main

import "fmt"

func sum(a []int, result chan int)  {
	sum := 0
	for _,v := range a{
		sum += v
	}
	result <- sum

}

func main()  {
	a := []int{1,2,3,4,5,6,7,8}
	result := make(chan int)
	go sum (a[:len(a)/2],result)
	go sum(a[len(a)/2:],result)
	x,y := <-result,<-result

	fmt.Println(x,y,x+y)

}
