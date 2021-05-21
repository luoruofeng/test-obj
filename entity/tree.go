package entity

import "fmt"

type TreeNode struct{
	Value int
	Left,Right *TreeNode
}

func CreateNode(value int)*TreeNode {
	return &TreeNode{value,nil,nil}
}

func (node TreeNode) Print(){ //node值传递，拷贝新值用
	fmt.Println(node.Value)
}

func (node *TreeNode)SetNode(newValue int){ //node引用传递
	node.Value = newValue
}

func (node *TreeNode) PrintAllY(){
	node.Print()
	if node.Left != nil {
		node.Left.PrintAllY()
	}
	if node.Right != nil {
		node.Right.PrintAllY()
	}
}

func (node *TreeNode) PrintAllX(isRoot bool){
	if isRoot{
		node.Print()
	}

	if node.Left != nil {
		node.Left.Print()
	}
	if node.Right != nil {
		node.Right.Print()
	}

	if node.Left != nil {
		node.Left.PrintAllX(false)
	}
	if node.Right != nil {
		node.Right.PrintAllX(false)
	}
}
