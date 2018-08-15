package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

// Human 实现sayHi 方法

func (human Human) sayHi(){
	fmt.Printf("Hi, I am %s you can call me on %s \n",human.name,human.phone)
}

// Human 实现Sing 方法
func (h Human) sing(lyrics string)  {
	fmt.Println("la .. la .. la .. ",lyrics)
}

type  Student struct{
	Human  //匿名字段
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

// Employee 重载Human的sayHi方法
func (e Employee)sayHi()  {
	fmt.Printf("Hi,I am %s,I work at %s. Call me on %s \n",e.name,e.company,e.phone)
}

// 定义一个接口
type Men interface {
	sayHi()
	sing(lyrics string)
}

func main()  {
	mike := Student{Human{"Mike",25,"222-222-XXX"},"MIT",0.00}
	paul := Student{Human{"paul",24,"333-333-XXX"},"SEU",5.00}
	sam := Employee{Human{"Sam",36,"444-444-xxx"},"Golang Inc",1000}
	Tom := Employee{Human{"Tom",37,"555-555-xxx"},"Things Ltd",5000}

	//定义Men的类型
	var i Men

	//i 能够存储Student
	i = mike
	fmt.Println("This is Mike, a student")
	i.sayHi()
	i.sing("November rain")

	// i也能存储Employee
	i = Tom
	fmt.Println("This is Tom,an Employee")
	i.sayHi()
	i.sing("born to be wild")

	// 定义slice Men
	fmt.Println("Let's use a slice of Men and see what happends")
	x := make([]Men,3)

	//这三个是不同类型的元素，但是他们实现了interface的同一个接口
	x[0],x[1],x[2] = paul,sam,mike

	for _,value := range x{
		value.sayHi()
	}
}