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
	root := &node{val: 50}
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	root.Insert(Randomint())
	fmt.Print(root.Search(15))
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
