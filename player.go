package main

import (
	"fmt"
	"os"
	"bufio"
	"github.com/fatih/color"
)

// returns the index of the player's decision
func getPlayerDecision (legalMoves []int, turn int, p1 int, p2 int) int{
	if (turn == 1 && p1 == 0) || (turn == 2 && p2 == 0){ // if it is a real player's turn
		fmt.Println("Select a turn below by typing in a number.")
		for index, element := range legalMoves {
			fmt.Printf("%d) %s\n", index + 1, convertIntToCoords(element))
		}

		for {
			reader := bufio.NewReader(os.Stdin)
			indexStr, _ := reader.ReadString('\n')
			index, err := convertStringToInt(indexStr)

			if err == nil && index - 1 < len(legalMoves) && index > 0{
				return index - 1 
			}

			color.Red("Please enter a valid move.")
		}
	} else { // if it is an AI's turn
		return getAITurn(legalMoves)
	}
}