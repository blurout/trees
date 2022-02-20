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
	tree := &node{val: 0}
	print_in_order(tree)
	fmt.Println(Is_Tree_Symmetrical(tree.left, tree.right))
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
	print_in_order_helper(root, &arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
func print_in_order_helper(root *node, arr *[]int) {
	if root == nil {
		return
	}
	print_in_order_helper(root.left, arr)
	*arr = append(*arr, root.val)
	print_in_order_helper(root.right, arr)
}

func Is_Tree_Symmetrical(left_side *node, right_side *node) bool {
    if (left_side == nil || right_side == nil ) {
        return left_side == right_side
    }
    if (left_side.val != right_side.val) {
        return false
    }
    return Is_Tree_Symmetrical(left_side.left, left_side.right) && Is_Tree_Symmetrical(left_side.right, right_side.left)
}

func Are_Trees_Symmetrical(tree1 *node, tree2 *node) bool {
    if tree1.val != tree2.val {
		return false
	}
	bool1 := Is_Tree_Symmetrical(tree1.left, tree1.right)
	bool2 := Is_Tree_Symmetrical(tree2.left, tree2.right)
    return bool1 && bool2
}

