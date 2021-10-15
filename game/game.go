package game

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type Game struct {
	root   *Node
	target State
	Times  [][]int
}

func SetUpGame(n int) *Game {
	initial := BlankState()
	target := BlankState()
	for i := n; i > 0; i-- {
		initial.a.Push(i)
		target.c.Push(i)
	}
	times := [][]int{}
	nn := float64(n)
	for i := 0; i < n; i++ {
		times = append(times, []int{})
	}
	for x := 1.0; x <= math.Pow(2, nn); x++ {
		for y := 1.0; y <= nn; y++ {
			if math.Pow(2, y)*(x)-math.Pow(2, y-1) < math.Pow(2, nn) {
				times[int(y)-1] = append(times[int(y)-1], int(math.Pow(2.0, y)*x-math.Pow(2.0, y-1.0)))
			}
		}
	}
	root := NewNode(initial)
	root.heuris = 1
	return &Game{root, *target, times}
}

func (game *Game) String() string {
	return fmt.Sprintf("Game(%v, %v)", game.root.Value(), game.target)
}

func (game *Game) Search() *Node {
	current := game.root
	step := 1
	n := len(game.target.c)
	str := ""
	for !game.target.Eq(*current.value) {
		current.children = CreateNodes(current.value.ApplyRules(), current)
		sort.SliceStable(current.children, func(i, j int) bool {
			return current.children[i].HeuristicValue(n, step, game.Times) < current.children[j].HeuristicValue(n, step, game.Times)
		})
		str += fmt.Sprintf("%v => %v\n", current, current.children)
		if len(current.children) > 0 {
			current = current.children[0]
		} else {
			return nil
		}
		step += 1
	}
	os.WriteFile("generated_states.txt", []byte(str[:len(str)-1]), 0644)
	return current
}
