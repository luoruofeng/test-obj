package main

import "github.com/luoruofeng/test-obj/entity"

func main()  {

	var node1 *entity.TreeNode = &entity.TreeNode{}
	node1.Value =22

	node2 := new(entity.TreeNode)
	node2.Value =33

	node3 := &entity.TreeNode{}
	node3.Value = 44

	node4 := &entity.TreeNode{66,nil,nil}

	node5 := entity.CreateNode(99)
	node5.Left = node4
	node5.Right = node3
	node5.Left.Right = node2
	node5.Right.Left = node1

	TreeNodes := []*entity.TreeNode{node5,node4,node3,node2,node1}

	node1.SetNode(3333)

	for _,v := range TreeNodes{
		v.Print()
	}

	node5.PrintAllY()

	node5.PrintAllX(true)
}
