package tree

import "fmt"

// 在Go语言中，可以将结构体的成员函数和成员变量写在两个文件中；
// 这种解耦合，可以增加程序的可读性
func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()
}

// 函数式编程，将函数作为参数，用来回调,该函数实现的是中序遍历
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	// 中序遍历
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
