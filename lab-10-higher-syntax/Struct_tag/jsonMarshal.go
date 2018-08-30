package main

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	Name string		`json:"name"`
	Size string		`json:"size"`
	Profits int		`json:"profits"`
}

func main(){
	company := Company{Name:"lh",Size:"Large",Profits:10000}
	companyAsByte, err := json.Marshal(&company)
	if err != nil{
		fmt.Printf("Failed to json.Marshal")
	}
	fmt.Printf(string(companyAsByte))

	/*companyAsByte2, err := json.Marshal(&company)
	if err != nil {
		fmt.Printf("err"+err.Error())
	}

	fmt.Printf(string(companyAsByte2))*/
}

