package main

import (
	"fmt"
)

var println = fmt.Println
var printf = fmt.Printf

type Tnode struct {
	data   int
	left   *Tnode
	right  *Tnode
	height int
}

type Root struct {
	root *Tnode
}

func main() {
	var i int
	println("Enter Root node")
	fmt.Scanf("%d\n", &i)
	root := newNode(i)
	populateTree(root)
	displayTree(root)
}

func populateTree(node *Tnode) *Tnode{
	var isLeft, isRight bool
	var i,j int
	println("Do you want to insert Left node")
	fmt.Scanf("%t\n", &isLeft)

	if isLeft{
		println("Enter Left Node")
	    fmt.Scanf("%d\n", &i)
		node.left = newNode(i)
		populateTree(node.left)
	}
	println("Do you want to insert Right node")
	fmt.Scanf("%t\n", &isRight)
	if isRight{
		println("Enter Right Node")
	    fmt.Scanf("%d\n", &j)
		node.right = newNode(j)
		populateTree(node.right)
	}
	return node
}

func newNode(data int) *Tnode {
	node := &Tnode{data: data}
	node.left = nil
	node.right = nil
	node.height = 1
	return node
}

func displayTree(node *Tnode){
	println("node data -->", node.data)

	if node.left != nil {
		displayTree(node.left)
	}
	if node.right != nil {
		displayTree(node.right)
	}
	return node
}