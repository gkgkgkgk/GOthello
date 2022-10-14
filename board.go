package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"io/ioutil"
	"log"
	"os"
)

// 0 is empty, 1 is black, 2 is white
// the board is addressed as board[row_number (y coordinate)][col_number (x coordinate)]
var board [8][8]int

func initializeBoard(){
	board = [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 2, 0, 0, 0},
		{0, 0, 0, 2, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0}}
}

func loadGame(filename string){
	data, err := ioutil.ReadFile(filename)

	if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }

	lines := strings.Split(string(data), "\n")
	for i := 0; i < 8; i++{
		squares := strings.Split(lines[i], " ")
		for j := 0; j < 8; j++ {
			board[i][j], _ = convertStringToInt(squares[j])
		}
	}

	turn, _ = convertStringToInt(lines[8])
	timeLimit, _ = convertStringToInt(lines[9])
}

func saveGame(){
	f, err := os.Create("./saves/autosave.txt")

	if err != nil {
        log.Fatalf("unable to write file: %v", err)
    }

	defer f.Close()

	for i := 0; i < 8; i++{
		line := ""
		for j := 0; j < 8; j++ {
			line += fmt.Sprintf("%d", board[i][j])
			if j != 7 {
				line += " "
			}
		}

		f.WriteString(line + "\n")
	}

	f.WriteString(fmt.Sprintf("%d\n%d", turn, timeLimit))
}

func printBoard(){
	fmt.Print("   ")
	for i:=0; i < 8; i++{
		fmt.Print(i)
		fmt.Print("     ")
	}
	fmt.Print("\n")

	line := ""
	for i:=0; i < 49; i++{
		line += "-"
	}

	for row := 0; row < 8; row++ {
		
		colorPrint(line, color.BgGreen, color.FgBlack)
		fmt.Print("\n")

		for i:=0; i < 2; i++ {
			colorPrint("|", color.BgGreen, color.FgBlack)

			for col:=0; col < 8; col++{
				squareColor := color.BgWhite
				if board[row][col] == 0 {
					squareColor = color.BgGreen
				} else if board[row][col] == 1 {
					squareColor = color.BgBlack
				}
				colorPrint("     ", color.FgBlack, squareColor)
				colorPrint("|", color.BgGreen, color.FgBlack)
			}
			if i == 0 {
				fmt.Printf(" %d", row)
			}
			fmt.Print("\n")
		}
	}

	colorPrint(line, color.BgGreen, color.FgBlack)
	fmt.Print("\n")
}

func getAllLegalMoves(player int) (legalMoves []int) {
	for row := 0; row < 8; row++{
		for col := 0; col < 8; col++{
			if isLegalMove(player, row, col){
				legalMoves = append(legalMoves, row * 8 + col)
			}
		}
	}

	return
}

func isLegalMove(player int, row int, col int) bool{
	opponent := 0

	if player == 1 {
		opponent = 2
	} else {
		opponent = 1
	}

	if board[row][col] == 0 {
		nw, n, ne, w, e, sw, s, se := -1,-1,-1,-1,-1,-1,-1,-1

		if row != 0 {
			n = board[row - 1][col]

			if col != 0 {
				nw = board[row - 1][col - 1]
			}
			
			if col != 7 {
				ne = board[row - 1][col + 1]
			}
		}
		
		if row != 7 {
			s = board[row + 1][col]

			if col != 0 {
				sw = board[row + 1][col - 1]
			}
			
			if col != 7 {
				se = board[row + 1][col + 1]
			}
		}
		
		if col != 0 {
			w = board[row][col - 1]
		}
		
		if col != 7 {
			e = board[row][col + 1]
		}
		

		if nw != opponent && n != opponent && ne != opponent && w != opponent && e != opponent && sw != opponent && s != opponent && se != opponent{
			return false
		}

		if nw == opponent {
			for d := 2; col - d >= 0 && row - d >= 0; d++ {
				nextPiece := board[row - d][col - d]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if n == opponent {
			for d := 2; row - d >= 0; d++ {
				nextPiece := board[row - d][col]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if ne == opponent {
			for d := 2; col + d <= 0 && row - d >= 0; d++ {
				nextPiece := board[row - d][col + d]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if w == opponent {
			for d := 2; col - d >= 0; d++ {
				nextPiece := board[row][col - d]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if e == opponent {
			for d := 2; col + d <= 7; d++ {
				nextPiece := board[row][col + d]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if sw == opponent {
			for d := 2; col - d >= 0 && row + d <= 7; d++ {
				nextPiece := board[row + d][col - d]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if s == opponent {
			for d := 2; row + d <= 7; d++ {
				nextPiece := board[row + d][col]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if se == opponent {
			for d := 2; col + d <= 7 && row + d <= 7; d++ {
				nextPiece := board[row + d][col + d]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		return false
	} else {
		return false
	}
}

func placePiece(pos int){
	row := (int)(pos / 8)
	col := pos % 8
	board[row][col] = turn
	
	opponent := 0

	if turn == 1 {
		opponent = 2
	} else {
		opponent = 1
	}

	nw, n, ne, w, e, sw, s, se := -1,-1,-1,-1,-1,-1,-1,-1

	if row != 0 {
		n = board[row - 1][col]

		if col != 0 {
			nw = board[row - 1][col - 1]
		}
		
		if col != 7 {
			ne = board[row - 1][col + 1]
		}
	}
	
	if row != 7 {
		s = board[row + 1][col]

		if col != 0 {
			sw = board[row + 1][col - 1]
		}
		
		if col != 7 {
			se = board[row + 1][col + 1]
		}
	}
	
	if col != 0 {
		w = board[row][col - 1]
	}
	
	if col != 7 {
		e = board[row][col + 1]
	}

	if nw == opponent {
		for d := 2; col - d >= 0 && row - d >= 0; d++ {
			nextPiece := board[row - d][col - d]
			if nextPiece == 0 {
				break
			}

			if nextPiece == turn {
				for d_flip := 1; d_flip < d; d_flip++ {
					board[row - d_flip][col - d_flip] = turn
				} 
				break
			}
		}
	}

	if n == opponent {
		for d := 2; row - d >= 0; d++ {
			nextPiece := board[row - d][col]
			if nextPiece == 0 {
				break
			}

			if nextPiece == turn {
				for d_flip := 1; d_flip < d; d_flip++ {
					board[row - d_flip][col] = turn
				} 
				break
			}
		}
	}

	if ne == opponent {
		for d := 2; col + d <= 0 && row - d >= 0; d++ {
			nextPiece := board[row - d][col + d]
			if nextPiece == 0 {
				break
			}

			if nextPiece == turn {
				for d_flip := 1; d_flip < d; d_flip++ {
					board[row - d_flip][col + d_flip] = turn
				} 
				break
			}
		}
	}

	if w == opponent {
		for d := 2; col - d >= 0; d++ {
			nextPiece := board[row][col - d]
			if nextPiece == 0 {
				break
			}

			if nextPiece == turn {
				for d_flip := 1; d_flip < d; d_flip++ {
					board[row][col - d_flip] = turn
				} 
				break
			}
		}
	}

	if e == opponent {
		for d := 2; col + d <= 7; d++ {
			nextPiece := board[row][col + d]
			if nextPiece == 0 {
				break
			}

			if nextPiece == turn {
				for d_flip := 1; d_flip < d; d_flip++ {
					board[row][col + d_flip] = turn
				} 
				break
			}
		}
	}

	if sw == opponent {
		for d := 2; col - d >= 0 && row + d <= 7; d++ {
			nextPiece := board[row + d][col - d]
			if nextPiece == 0 {
				break
			}

			if nextPiece == turn {
				for d_flip := 1; d_flip < d; d_flip++ {
					board[row + d_flip][col - d_flip] = turn
				} 
				break
			}
		}
	}

	if s == opponent {
		for d := 2; row + d <= 7; d++ {
			nextPiece := board[row + d][col]
			if nextPiece == 0 {
				break
			}

			if nextPiece == turn {
				for d_flip := 1; d_flip < d; d_flip++ {
					board[row + d_flip][col] = turn
				} 
				break
			}
		}
	}

	if se == opponent {
		for d := 2; col + d <= 7 && row + d <= 7; d++ {
			nextPiece := board[row + d][col + d]
			if nextPiece == 0 {
				break
			}

			if nextPiece == turn {
				for d_flip := 1; d_flip < d; d_flip++ {
					board[row + d_flip][col + d_flip] = turn
				} 
				break
			}
		}
	}
}