package main

type treeNode struct {
	val int
	left *treeNode
	right *treeNode
}

func main() {
	root := treeNode{val: 9, left: nil, right: nil}
	leftChild := treeNode{val: 6, left: nil, right: nil}
	rightChild := treeNode{val: 12, left: nil, right: nil}
	root.left, root.right = &leftChild, &rightChild
}