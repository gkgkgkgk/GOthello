package main

import (
	"fmt"
	"github.com/fatih/color"
	"bufio"
	"os"
)

func beginGame(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Othello.")
	fmt.Println("")
	fmt.Print("Is ")
	color.Set(color.FgBlack)
	color.Set(color.BgGreen)
	fmt.Print("Player 1 (black pieces)")
	color.Unset()

	fmt.Println(" a human (h) or AI (a)?")
	color.Unset()
	
	for {
		p1str, _ := reader.ReadString('\n')

		if p1str[0:1] == "h" || p1str[0:1] == "a" {
			if p1str[0:1] == "h"{
				p1 = 0
			} else {
				p1 = 1
			}
			break
		}

		color.Set(color.FgRed)
		fmt.Println("Please type an 'h' for human or an 'a' for AI.")
		color.Unset()
	}

	

	fmt.Print("Is ")
	color.Set(color.FgWhite)
	color.Set(color.BgGreen)
	fmt.Print("Player 2 (white pieces)")
	color.Unset()

	
	fmt.Println(" a human (h) or AI (a)?")
	color.Unset()
	for {
		p2str, _ := reader.ReadString('\n')

		if p2str[0:1] == "h" || p2str[0:1] == "a" {
			if p2str[0:1] == "h"{
				p2 = 0
			} else {
				p2 = 1
			}
			break
		}

		color.Set(color.FgRed)
		fmt.Println("Please type an 'h' for human or an 'a' for AI.")
		color.Unset()
	}

	fmt.Println("What is the AI time limit (in seconds) for each turn?")
	
	for {
		timeStr, err := reader.ReadString('\n')	
		timeLimit, err = convertStringToInt(timeStr)

		if err != nil {
			color.Set(color.FgRed)
			fmt.Println("Please input a valid number.")
		} else {
			break
		}
		color.Unset()
	}

	fmt.Println("Which player will make the first move?")
	for {
		turnStr, err := reader.ReadString('\n')	
		turn, err = convertStringToInt(turnStr)

		if err != nil || (turn != 1 && turn != 2 ){
			color.Set(color.FgRed)
			fmt.Println("Please input a a 1 or 2.")
		} else {
			break
		}
		color.Unset()
	}

	fmt.Print("\nThe game will begin with Player 1 as a")

	if p1 == 0 {
		fmt.Print(" human")
	} else {
		fmt.Print("n AI")
	}

	fmt.Print(" and Player 2 as a")
	if p2 == 0 {
		fmt.Print(" human")
	} else {
		fmt.Print("n AI")
	}

	fmt.Println(".")
	
	fmt.Printf("Player %d will go first and the AI time limit is %d seconds.\nGood Luck!\n\n", turn, timeLimit)
}

func gameLoop(){
	initializeBoard()
	if loadpath != "" {
		if loadpath == "autosave" {
			loadGame("./saves/autosave.txt")
		} else {
			loadGame(loadpath)
		}
	}

	printBoard()
	getAllLegalMoves(turn)
}