package game

import "fmt"

type Node struct {
	value    *State
	children []*Node
	parent   *Node
}

func NewNode(value *State) *Node {
	return &Node{value, []*Node{}, nil}
}

func CreateNodes(values []*State, parent *Node) []*Node {
	nodes := []*Node{}
	for _, v := range values {
		nodes = append(nodes, &Node{v, []*Node{}, parent})
	}
	return nodes
}

type FindFunc func(State) bool

func (root *Node) Find(comparator FindFunc) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if comparator(*nextUp.value) {
			return nextUp
		}
		queue = append(queue, nextUp.children...)
	}
	return nil
}

func (node *Node) CreateChildren(value *State) *Node {
	children := NewNode(value)
	node.AddChildren(children)
	return children
}

func (node *Node) AddChildren(children *Node) {
	children.parent = node
	node.children = append(node.children, children)
}

func (node *Node) Remove() {
	// Remove the node from it's parents children collection
	for idx, sibling := range node.parent.children {
		if sibling == node {
			node.parent.children = append(
				node.parent.children[:idx],
				node.parent.children[idx+1:]...,
			)
		}
	}

	// If the node has any children, set their parent to nil and set the node's children collection to nil
	if len(node.children) != 0 {
		for _, child := range node.children {
			child.parent = nil
		}
		node.children = []*Node{}
	}
}

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.value)
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

func (node *Node) StringParents() string {
	tmp := ""
	var cur *Node = node
	for cur != nil {
		tmp = fmt.Sprintln(cur.value) + tmp
		cur = cur.parent
	}
	return tmp[:len(tmp)-1]
}
