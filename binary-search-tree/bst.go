package main
import (
	"fmt"
)

type Node struct{
	Date int
	LeftChild *Node
	RightChild *Node
}

type BinarySearchTree struct{
	Root *Node
}

func (bst *BinarySearchTree) getRoot() *Node{
	return bst.Root
}

func main(){
	fmt.Println("Printing binary tree")
	bst := new(BinarySearchTree)
	fmt.Println(bst)
}

