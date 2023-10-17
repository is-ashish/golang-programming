package main

import (
	"fmt"
	"log"
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
	println("Create Root node")
	    _, err := fmt.Scanf("%d", &i)
    if err != nil {
        log.Fatal(err)
    }
	root := newNode(i)
	println("======Display root node data", root.data)
	populateTree(root)
	println("======Display left node data", root.left)
	println("======Display right node data", root.right)

}

func populateTree(node *Tnode) *Tnode{
	var isYes bool
	var i int
	println("Do you want to insert Left node")
	_, err := fmt.Scanf("%t", &isYes)
	if err != nil { 
		log.Fatal(err)
	}
	if isYes{
	    _, err := fmt.Scanf("%d", &i)
    	if err != nil {
        	log.Fatal(err)
    	}
		node.left = newNode(i)
		populateTree(node.left)
	}

	println("Do you want to insert Right node")
	_, err1 := fmt.Scanf("%t", &isYes)
	if err1 != nil {
		log.Fatal(err1)
	}
	if isYes{
	    _, err := fmt.Scanf("%d", &i)
    	if err != nil {
        	log.Fatal(err)
    	}
		node.right = newNode(i)
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
