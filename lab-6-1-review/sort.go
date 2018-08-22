package main

import (
	"fmt"
	"sort"
)

func main()  {
	stringList := [] string {"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}
	sort.Strings(stringList)
	fmt.Println(stringList)

}