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
	tree := &node{val: 3}
	tree.Insert(2)
	tree.right = &node{val: 2}
	tree2 := &node{val: 3}
	tree.Insert(2)
	tree.right = &node{val: 5}
	Valid_BST(tree)
	print_in_order(tree)
	Is_Tree_Symmetrical(tree.left, tree.right)
	Are_Trees_Symmetrical(tree, tree2)
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
	num += rand.Intn(10)
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

func print_in_order(root *node) {
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

func Is_Tree_Symmetrical(left_side *node, right_side *node) {
	Symmetrical := Is_Tree_Symmetrical_Helper(left_side, right_side)
	if Symmetrical {
		fmt.Println("The Tree is symmetrical")
	} else {
		fmt.Println("The Tree is not symmetrical")
	}
}

func Is_Tree_Symmetrical_Helper(left_side *node, right_side *node) bool {
    if (left_side == nil || right_side == nil ) {
        return left_side == right_side
    }
    if (left_side.val != right_side.val) {
        return false
    }
    return Is_Tree_Symmetrical_Helper(left_side.left, left_side.right) && Is_Tree_Symmetrical_Helper(left_side.right, right_side.left)
}

func Are_Trees_Symmetrical(tree1 *node, tree2 *node)  {
    if tree1.val != tree2.val {
		fmt.Println("The Trees are not symmetrical")
	}
	bool1 := Is_Tree_Symmetrical_Helper(tree1.left, tree1.right)
	bool2 := Is_Tree_Symmetrical_Helper(tree2.left, tree2.right)
    Symmetrical := bool1 && bool2
	if Symmetrical {
		fmt.Println("The Trees are symmetrical")
	} else {
		fmt.Println("The Trees are not symmetrical")
	}
}

func Sum_Of_Left_Leaves(root *node) int {
	if root == nil {
		return 0
	}
	sum := 0
	left := root.left
	if (left != nil && left.left == nil && left.right == nil) {
		sum += left.val
	}
	return sum + Sum_Of_Left_Leaves(root.left) + Sum_Of_Left_Leaves(root.right)
}

func Valid_BST(root *node) {
    valid := Valid_BST_Helper(nil, nil, root)
	if valid {
		fmt.Println("The Tree is a Valid BST")
	} else {
		fmt.Println("The Tree is not a Valid BST")
	}
}
func Valid_BST_Helper(low, high, root *node) bool {
    if root == nil {
        return true
    }
    if low != nil && root.val <= low.val {
        return false
    }
    if high != nil && root.val >= high.val {
        return false
    }
    return Valid_BST_Helper(low, root, root.left) && Valid_BST_Helper(root, high, root.right)
}
