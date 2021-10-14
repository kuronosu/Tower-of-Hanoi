package main

import (
	"fmt"

	"github.com/kuronosu/Tower-of-Hanoi/game"
)

func main() {
	game := game.SetUpGame(4)
	if result := game.Search(); result != nil {
		fmt.Println(result.StringParents())
	} else {
		fmt.Println("Not found")
	}
}
