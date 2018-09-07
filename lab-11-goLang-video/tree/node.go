package tree

import "fmt"

//定义二叉树节点
type Node struct {
	Value       int
	Left, Right *Node
}

// 为Node节点定义函数
func (node *Node) Print() {
	fmt.Print(node.Value, " ")
}

// 为节点设值
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting value to nil" + "node . Ignored.")
		return
	}
	node.Value = value
}

// 创建一个节点

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
