package main

import ( 
	"fmt"
	"math/rand"
	"time"
)
type node struct {
	val   int
	left  *node
	right *node
}

func main() {
	root := &node{val: 0}
	root.Insert(1)
	root.Insert(2)
	root.Insert(3)
	root.Insert(4)
	root.Insert(5)
	root.Insert(6)
	print_in_order(root)
}

func (n *node) Insert (data int) {
	if n.val < data {
		if n.right == nil {
			n.right = &node{val: data}
		} else {
			n.right.Insert(data)
		}
	} else if n.val > data {
		if n.left == nil {
			n.left = &node{val: data}
		} else {
			n.left.Insert(data)
		}
	}
}

func (n *node) Search (data int) bool {
	if n == nil {
		return false
	}
	if n.val < data {
		return n.right.Search(data)
	} else if n.val > data {
		return n.left.Search(data)

	}
	return true
}

func Randomint() int {
	rand.Seed(time.Now().UnixNano())
	num := 0
	num += rand.Intn(100)
	return num
}

func search_BST(root *node, val int) bool {
    if root == nil {
        return false
    } else if val > root.val {
        return search_BST(root.right, val)
    } else if val < root.val {
        return search_BST(root.left, val)
    }
    return true
}

func invert_tree(root *node) *node {
	if root == nil {
		return nil
	}
	tmp := root.left
	root.left = invert_tree(root.right)
	root.right = invert_tree(tmp)
	return root
}

func print_in_order (root *node) {
	arr:= []int{}
	helper(root, &arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
func helper(root *node, arr *[]int) {
	if root == nil {
		return
	}
	helper(root.left, arr)
	*arr = append(*arr, root.val)
	helper(root.right, arr)
}