package main

import (
	"fmt"
	"log"

	"github.com/kuronosu/Tower-of-Hanoi/game"
)

func main() {
	var i int
	fmt.Print("Number of discs: ")
	if _, err := fmt.Scanf("%d", &i); err != nil || i <= 0 {
		log.Fatal("Invalid input")
	}
	game := game.SetUpGame(i)
	if result := game.Search(); result != nil {
		fmt.Println(result.StringParents())
	} else {
		fmt.Println("No solution found")
	}
}
