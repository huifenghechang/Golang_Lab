package main

import "fmt"

/*
	struct 没有构造函数，使用工厂函数来定制结构体实例
*/

type TreeNode struct {
	Value       int
	left, right *TreeNode
}

/*
	Go 语言没有构造函数
	虽然函数返回的是局部变量，但是因为GO语言编译器自带垃圾回收机制，
	所以，该函数依旧可以正常使用

*/
func createTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func main() {
	var root TreeNode

	// 给结构体赋值
	root = TreeNode{6, nil, nil}
	root.left = &TreeNode{}
	root.right = &TreeNode{7, nil, nil}
	// 在结构体中，无论是成员，还是指针，一律用.来表示
	root.right.left = new(TreeNode)
	nodes := []TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)

}
