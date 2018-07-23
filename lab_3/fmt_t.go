package main

import "fmt"

func main()  {
	text := "\u5B9E\u9A8C\u697C"
	fmt.Printf("bool output :\n%t\n%t\n",true,false)
	fmt.Println("number output, origin value: 64")
	fmt.Printf("|%b|%8b|%-8b|%08b|% 8b|\n",64,64,64,64,64)
	fmt.Printf("|%x|%8x|%-8x|%08x|% 8x|\n",64,64,64,64,64)
	fmt.Println(`text output, origin value: \u5B9E\u9A8C\u697C`)
	fmt.Printf("content: %s\n", text)
	fmt.Printf("hex value: %X \n Unicode value:",text)
	for _, char := range text{
		fmt.Printf("%U",char)
	}
	fmt.Println()
	bytes := []byte(text)
	fmt.Printf("value of bytes: %s",bytes)
	fmt.Println()
	fmt.Printf("hex value of bytes:% X \n",bytes)
	fmt.Println()
	fmt.Printf("origin value in bytes: %v\n",bytes)

}
