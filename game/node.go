package game

import (
	"fmt"
)

type Node struct {
	value    *State
	children []*Node
	parent   *Node
	heuris   int
}

func NewNode(value *State) *Node {
	return &Node{value, []*Node{}, nil, 0}
}

func CreateNodes(values []*State, parent *Node) []*Node {
	nodes := []*Node{}
	for _, v := range values {
		nodes = append(nodes, &Node{v, []*Node{}, parent, 0})
	}
	return nodes
}

// type FindFunc func(State) bool

// func (root *Node) Find(comparator FindFunc) *Node {
// 	queue := make([]*Node, 0)
// 	queue = append(queue, root)
// 	for len(queue) > 0 {
// 		nextUp := queue[0]
// 		queue = queue[1:]
// 		if comparator(*nextUp.value) {
// 			return nextUp
// 		}
// 		queue = append(queue, nextUp.children...)
// 	}
// 	return nil
// }

// func (node *Node) CreateChildren(value *State) *Node {
// 	children := NewNode(value)
// 	node.AddChildren(children)
// 	return children
// }

// func (node *Node) AddChildren(children *Node) {
// 	children.parent = node
// 	node.children = append(node.children, children)
// }

// func (node *Node) Remove() {
// 	for idx, sibling := range node.parent.children {
// 		if sibling == node {
// 			node.parent.children = append(
// 				node.parent.children[:idx],
// 				node.parent.children[idx+1:]...,
// 			)
// 		}
// 	}

// 	if len(node.children) != 0 {
// 		for _, child := range node.children {
// 			child.parent = nil
// 		}
// 		node.children = []*Node{}
// 	}
// }

func (node *Node) String() string {
	if node.heuris != 0 {
		return fmt.Sprintf("{State: %v, Heuristic: %v}", node.value, node.heuris)
	}
	return fmt.Sprintf("{State: %v, Heuristic: not calculated}", node.value)
}

func (node *Node) Value() *State {
	return node.value
}

func (node *Node) Parents() []Node {
	tmp := []Node{}
	var cur *Node = node
	for cur != nil {
		tmp = append(tmp, *cur)
		cur = cur.parent
	}
	return tmp
}

func (node *Node) ValueParents() []State {
	tmp := []State{}
	var cur *Node = node
	for cur != nil {
		tmp = append(tmp, *cur.value)
		cur = cur.parent
	}
	return tmp
}

func (node *Node) StringValueParents() string {
	tmp := ""
	var cur *Node = node
	for cur != nil {
		tmp = fmt.Sprintln(cur.value) + tmp
		cur = cur.parent
	}
	return tmp[:len(tmp)-1]
}

func (node *Node) StringParents() string {
	tmp := ""
	var cur *Node = node
	for cur != nil {
		tmp = fmt.Sprintln(cur) + tmp
		cur = cur.parent
	}
	return tmp[:len(tmp)-1]
}

func (node *Node) HeuristicValue(numberDisks, step int, times [][]int) int {
	if node.heuris != 0 {
		return node.heuris
	}
	if node.parent == nil {
		node.heuris = 0
		return node.heuris
	}
	disk, _, weight, err := Movement(*node.parent.value, *node.value)
	if err != nil {
		return -1
	}

	heuris := 0
	if intInSlice(step, times[disk-1]) {
		heuris += 1
	} else {
		heuris += 3
	}

	if weight == ExpectedDirection(disk, numberDisks) {
		heuris += 1
	} else {
		heuris += 3
	}
	heuris += step
	node.heuris = heuris
	return node.heuris
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
