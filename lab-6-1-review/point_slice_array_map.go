package main

import "fmt"

// 数组、切片等数据类型，属于符合类型


// 关于指针的操作
func swap1(x int,y int){
	var temp int
	temp = x
	x = y
	y = temp
}

func swap2(x *int,y *int )  {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

// 关于数组、切片的操作
/*
	数组的特点是：定长、且数组内的元素都是相同的

	切片的特点是：长度可以提调整、是引用类型
	创建语法:
		- make([]Type,length,Capacity)
		- make([]Type,length)


	关于map的操作
	map是一种内置的数据结构，属于键值对类型，Map是引用类型
	创建语法
		- make(map[KeyType]ValueType,initialCapacity)
		- make(map[KeyType]ValueType)
*/

func main()  {

	//测试指针的代码
	var x,y  = 10000,5
	swap1(x,y)
	fmt.Println("x|y",x,y)
	swap2(&x,&y)
	fmt.Println("x|y ",x,y)

	//测试切片的代码
	s := make([]int,10,20)
	s[4] = 8
	fmt.Printf("len and Cap of slice : %v is:%d and %d \n",s,len(s),cap(s))

	//测试map的代码
	hdu := make(map[string]string)
	hdu["location"] = "hangzhou"
	hdu["name"] = "HangZhouDianZi University"
	hdu["score"] = "80-90"
	//查询value的值
	v, found := hdu["location"]
	fmt.Println("Found key \"location\" in hdu \n",found,v)
	//遍历输出map的值
	for k,v := range hdu{
		fmt.Printf("\"%s\":\"%s\"\n",k,v)
	}


}
