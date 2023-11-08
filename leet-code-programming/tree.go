package main

import (
	"fmt"
	"math"
)

var println = fmt.Println
var printf = fmt.Printf

type Tnode struct {
	data   int
	left   *Tnode
	right  *Tnode
	height int
}

type RNode struct {
	root *Tnode
}

func main() {
	 r := &RNode{}
	values := []int{64,55,45,56,67,65,68}
	for _, v := range values {
		r.root = populateBinarySearchTree(r.root,v)
	}
	println("Height of tree is -->> ", r.root.height)
	// displayTree(r.root)
	println("Pre-Order traversal is -->>")
	displayPreOrder(r.root)
	println()
	println("In-Order traversal is -->>")
	displayInOrder(r.root)
	println()
	println("Post-Order traversal is -->>")
	displayPostOrder(r.root)
	println()
}

func populateBinarySearchTree(node *Tnode, value int) *Tnode {
	if node == nil{
		return newNode(value)
	}

	if value <= node.data {
		node.left = populateBinarySearchTree(node.left, value)
	}

	if value > node.data {
		node.right = populateBinarySearchTree(node.right, value)
	}

	node.height = int(math.Max(float64(height(node.left)),float64(height(node.right)))+1)

	return node
}

func populateTree(node *Tnode) *Tnode {
	var isLeft, isRight bool
	var i,j int
	println("Do you want to insert Left node of", node.data)
	fmt.Scanf("%t\n", &isLeft)
	if isLeft{
		println("Enter Left Node")
	    fmt.Scanf("%d\n", &i)
		node.left = newNode(i)
		populateTree(node.left)
	}
	println("Do you want to insert Right node of", node.data)
	fmt.Scanf("%t\n", &isRight)


	if isRight{
		println("Enter Right Node")
	    fmt.Scanf("%d\n", &j)
		node.right = newNode(j)
		populateTree(node.right)
	}

	node.height = int(math.Max(float64(height(node.left)),float64(height(node.right)))+1)
	return node
}

func newNode(data int) *Tnode {
	node := &Tnode{}
	node.data = data
	node.left = nil
	node.right = nil
	return node
}

func displayTree(node *Tnode) *Tnode {

	if node != nil {
		println("Current Tree node is ------->> ", node.data)
	}
	if node.left != nil {
		println("Left child of Node", node.data, "is", node.left.data)
		displayTree(node.left)
	}
	if node.right != nil {
		println("Right child of Node", node.data, "is", node.right.data)
		displayTree(node.right)
	}
	return node
}

func displayPreOrder(node *Tnode) *Tnode {
	if node == nil {
	   return node
	}
	fmt.Print(" ", node.data)
	displayPreOrder(node.left)
	displayPreOrder(node.right)
	return node
}
func displayInOrder(node *Tnode) *Tnode {
	if node == nil {
	   return node
	}
	displayInOrder(node.left)
	fmt.Print(" ", node.data)
	displayInOrder(node.right)
	return node
}
func displayPostOrder(node *Tnode) *Tnode {
	if node == nil {
	   return node
	}
	displayPostOrder(node.left)
	displayPostOrder(node.right)
	fmt.Print(" ", node.data)
	return node
}



func height(node *Tnode) int {
	if node == nil {
		return -1
	}
	return node.height
}