package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	par   *Node
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil, nil}
}

func (n *Node) Value() int {
	return n.value
}

type BinSearchTree struct {
	root *Node
}

func NewBinSearchTree() *BinSearchTree {
	return &BinSearchTree{nil}
}

func (bst *BinSearchTree) Parents(node *Node) []*Node {
	parents := make([]*Node, 0)

	tmp := node
	for {
		if tmp.par == nil {
			break
		} else {
			parents = append(parents, tmp.par)
			tmp = tmp.par
		}
	}

	return parents
}

func (bst *BinSearchTree) LowestCommonParent(n1, n2 int) *Node {
	root := bst.root

	for root != nil {
		if n1 < root.Value() && n2 < root.Value() {
			root = root.left
		} else if n1 > root.Value() && n2 > root.Value() {
			root = root.right
		} else {
			return root
		}
	}

	return nil

}

func (bst *BinSearchTree) Find(val int) *Node {
	node := bst.root

	for node != nil {
		if node.Value() == val {
			return node
		} else if val < node.Value() {
			node = node.left
		} else {
			node = node.right
		}
	}

	return nil
}

func (bst *BinSearchTree) Insert(val int) *Node {
	node := NewNode(val)
	if bst.root == nil {
		bst.root = node
		return node
	}

	current := bst.root
	for {
		if val < current.Value() {
			if current.left == nil {
				current.left = node
				node.par = current
				break
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = node
				node.par = current
				break
			}
			current = current.right
		}
	}

	return node
}

func (bst *BinSearchTree) String() string {

	var print func()

	print = func(node *Node) (string int) {
		if node == nil {
			return "", 0
		}

		rigth, rLen := print(node.right)
		left, lLen := print(node.left)
	}

}

func HardCodedTree() *BinSearchTree {
	datas := [7]int{30, 8, 52, 3, 20, 10, 29}

	tree := NewBinSearchTree()
	for _, i := range datas {
		tree.Insert(i)
	}

	return tree
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tree := HardCodedTree()

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		key1, _ := strconv.Atoi(fields[0])
		key2, _ := strconv.Atoi(fields[1])

		lwcp := tree.LowestCommonParent(key1, key2)
		if lwcp != nil {
			fmt.Println(lwcp.Value())
		}
	}
}
