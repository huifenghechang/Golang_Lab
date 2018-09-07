package main

import (
	"fmt"
	"lab-11-goLang-video/tree"
)

// 定义树节点
type myTreeNode struct {
	node *tree.Node
}

//  后序遍历
func (myTree *myTreeNode) postOrder() {
	if myTree == nil || myTree.node == nil {
		return
	}
	left := myTreeNode{myTree.node.Left}
	right := myTreeNode{myTree.node.Right}

	left.postOrder()
	right.postOrder()
	myTree.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	fmt.Println("In-Order traversal:")
	root.Traverse()

	fmt.Println("My own post-Order traversal:")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()
	maxNodeValue := 0
	for node := range c {
		if node.Value > maxNodeValue {
			maxNodeValue = node.Value
		}
	}
	fmt.Println("Max node value:", maxNodeValue)
}
