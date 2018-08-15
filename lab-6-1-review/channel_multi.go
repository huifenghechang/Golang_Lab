package main

import "fmt"

func multiy(a []int, result chan int)  {
	z := 1
	for _,v := range a{
		z *= v
	}
	result <- z
}

func main()  {
	a := []int{1,2,3,4,5,6,7,8}
	result := make(chan int)
	go multiy(a[:len(a)/2],result)
	go multiy(a[len(a)/2:],result)

	x,y := <-result,<-result

	fmt.Println(x,y,x*y)


}
