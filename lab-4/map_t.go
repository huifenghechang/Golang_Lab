package main

import "fmt"

func main()  {
	sxl := make(map[string]string)
	sxl["golang"] = "docker"
	sxl["python"] = "flask web framework"
	sxl["linux"] = "sys administrator"
	fmt.Print("Travel all keys:")
	for key := range sxl{
		fmt.Printf("% s",key)
		fmt.Println()
	}
	fmt.Println()

	delete(sxl,"python")
	sxl["linux"] = " I like Linux ~"

	v,found := sxl["linux"]
	fmt.Printf("Found key \"linux\" Yes or False: %t, value of key \"linux\":\"%s\"",found,v)
	fmt.Println()

}
