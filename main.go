package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kuronosu/Tower-of-Hanoi/game"
)

func main() {
	var n int
	fmt.Print("Number of discs: ")
	if _, err := fmt.Scanf("%d", &n); err != nil || n < 3 {
		log.Fatal("Invalid entry, minimum number of discs 3")
	}
	game := game.SetUpGame(n)
	if result := game.Search(); result != nil {
		resultStr := result.StringParents()
		resultStr = strings.ReplaceAll(resultStr, "{", "")
		resultStr = strings.ReplaceAll(resultStr, "}", "")
		os.WriteFile("solution_states.txt", []byte(resultStr), 0644)
		fmt.Println(resultStr)
	} else {
		fmt.Println("No solution found")
	}
	fmt.Scanln()
	fmt.Print("Finished, press enter to exit")
	fmt.Scanln()
}
