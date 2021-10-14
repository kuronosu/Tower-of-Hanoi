package game

import (
	"fmt"
	"sort"
)

type Game struct {
	root   *Node
	target State
}

func SetUpGame(n int) *Game {
	initial := BlankState()
	target := BlankState()
	for i := n; i > 0; i-- {
		initial.a.Push(i)
		target.c.Push(i)
	}
	return NewGame(initial, target)
}

func NewGame(initial, target *State) *Game {
	return &Game{NewNode(initial), *target}
}

func (game *Game) String() string {
	return fmt.Sprintf("Game(%v, %v)", game.root.Value(), game.target)
}

func (game *Game) Search() *Node {
	var current *Node
	open := []*Node{game.root}
	close := []*Node{}
	openLast := game.root
	for len(open) > 0 && !game.target.Eq(*openLast.value) {
		lastIndex := len(open) - 1
		current, open = open[lastIndex], open[:lastIndex] // Pop
		current.children = CreateNodes(current.value.ApplyRules(), current)
		close = append(close, current)
		for _, n := range current.children {
			if !ContainState(close, *n.value) && !ContainState(open, *n.value) {
				open = append(open, n)
			}
		}
		sort.SliceStable(open, func(i, j int) bool {
			return open[i].value.Weight() < open[j].value.Weight()
		})
		openLast = open[len(open)-1]
	}
	if openLast.value.Eq(game.target) {
		return openLast
	}
	return nil
}

func Eq(tower, other []*Node) bool {
	if len(tower) != len(other) {
		return false
	}
	for i, v := range tower {
		if v != other[i] {
			return false
		}
	}
	return true
}

func ContainState(list []*Node, state State) bool {
	for _, v := range list {
		if v.value.Eq(state) {
			return true
		}
	}
	return false
}
