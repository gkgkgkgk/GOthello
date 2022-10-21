package main

import (
	"fmt"
	"github.com/fatih/color"
	"bufio"
	"os"
)

func beginGame() (int, int, int, int){
	gameOver = false
	p1 := 0
	p2 := 1
	timeLimit := 1
	turn := 1

	color.Set(color.FgWhite)
	color.Set(color.BgBlack)
	fmt.Println("Welcome to Othello.")

	if defaultSettings {
		fmt.Println("Beginning game with default settings.")
	} else {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("")
		fmt.Print("Is ")
		colorPrint("Player 1 (black pieces)", color.BgGreen, color.FgBlack)
		fmt.Println(" a human (h) or AI (a)?")
		
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

			color.Red("Please type an 'h' for human or an 'a' for AI.\n")
		}

		fmt.Print("Is ")
		colorPrint("Player 2 (white pieces)", color.BgGreen, color.FgWhite)
		fmt.Println(" a human (h) or AI (a)?")

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

			color.Red("Please type an 'h' for human or an 'a' for AI.\n")
		}

		fmt.Println("What is the AI time limit (in seconds) for each turn?")
		
		for {
			timeStr, err := reader.ReadString('\n')	
			timeLimit, err = convertStringToInt(timeStr)

			if err != nil {
				color.Red("Please input a valid number.\n")
			} else {
				break
			}
		}

		fmt.Println("Which player will make the first move?")
		for {
			turnStr, err := reader.ReadString('\n')	
			turn, err = convertStringToInt(turnStr)

			if err != nil || (turn != 1 && turn != 2 ){
				color.Red("Please input a a 1 or 2.\n")
			} else {
				break
			}
		}
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
	color.Unset()

	return p1, p2, turn, timeLimit
}

func gameLoop(p1 int, p2 int, turn int, timeLimit int){
	board := initializeBoard()
	noTurn := false

	if loadpath != "" {
		if loadpath == "autosave" {
			board, turn, timeLimit = loadGame("./saves/autosave.txt")
		} else {
			board, turn, timeLimit = loadGame(loadpath)
		}
	}


	// Begin game loop
	for !gameOver {
		saveGame(board, turn, timeLimit)
		printBoard(board)

		fmt.Printf("It is Player %d's turn.\n", turn)
		legalMoves := getAllLegalMoves(turn, board)

		if len(legalMoves) > 0 {
			playerMove := getPlayerDecision(board, legalMoves, turn, p1, p2, timeLimit)
			fmt.Printf("\nPlayer %d placed a piece on %s.\n", turn, convertIntToCoords(legalMoves[playerMove]))
			board = placePiece(board, legalMoves[playerMove], turn)
		} else {
			if noTurn {
				gameOver = true
			} else {
				fmt.Printf("\nPlayer %d has no moves.\n", turn)
				noTurn = true
			}
		}

		if turn == 1 {
			turn = 2
		} else {
			turn = 1
		}
	}

	fmt.Println("GAME OVER")
	checkWinner(board)
}

func checkWinner(board [8][8] int){
	p1Score := 0
	p2Score := 0

	for i := 0; i < 8; i++{
		for j := 0; j < 8; j++ {
			if board[i][j] == 1 {
				p1Score++
			} else if board[i][j] == 2{
				p2Score++
			}
		}
	}

	if p1Score > p2Score {
		fmt.Printf("Player 1 won the game %d to %d.\n", p1Score, p2Score)
	} else if p2Score > p1Score{
		fmt.Printf("Player 2 won the game %d to %d.\n", p2Score, p1Score)
	} else {
		fmt.Printf("The game ended in a tie %d to %d.\n", p1Score, p2Score)
	}
}