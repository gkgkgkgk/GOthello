package main

import (
	"fmt"
)

func getAllLegalMoves(player int){
	var legalMoves []int
	for row := 0; row < 8; row++{
		for column := 0; column < 8; column++{
			if isLegalMove(player, row, column){
				legalMoves = append(legalMoves, row * 8 + column)
			}
		}
	}

	fmt.Printf("Found %d legal moves.", len(legalMoves))
	fmt.Printf("%v", legalMoves)
}

func isLegalMove(player int, row int, column int) bool{
	opponent := 0

	if player == 1 {
		opponent = 2
	} else {
		opponent = 1
	}

	if board[row][column] == 0 {
		nw, n, ne, w, e, sw, s, se := -1,-1,-1,-1,-1,-1,-1,-1

		if row != 0 {
			n = board[row - 1][column]

			if column != 0 {
				nw = board[row - 1][column - 1]
			}
			
			if column != 7 {
				ne = board[row - 1][column + 1]
			}
		}
		
		if row != 7 {
			s = board[row + 1][column]

			if column != 0 {
				sw = board[row + 1][column - 1]
			}
			
			if column != 7 {
				se = board[row + 1][column + 1]
			}
		}
		
		if column != 0 {
			w = board[row][column - 1]
		}
		
		if column != 7 {
			e = board[row][column + 1]
		}
		

		if nw != opponent && n != opponent && ne != opponent && w != opponent && e != opponent && sw != opponent && s != opponent && se != opponent{
			return false
		}

		if nw == opponent {
			for d := 2; column - d >= 0 && row - d >= 0; d++ {
				nextPiece := board[row - d][column - d]
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
				nextPiece := board[row - d][column]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if nw == opponent {
			for d := 2; column - d >= 0 && row - d >= 0; d++ {
				nextPiece := board[row - d][column - d]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if w == opponent {
			for d := 2; column - d >= 0; d++ {
				nextPiece := board[row][column - d]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if e == opponent {
			for d := 2; column + d <= 7; d++ {
				nextPiece := board[row][column + d]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if sw == opponent {
			for d := 2; column - d >= 0 && row + d <= 7; d++ {
				nextPiece := board[row + d][column - d]
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
				nextPiece := board[row + d][column]
				if nextPiece == 0 {
					break
				}

				if nextPiece == player {
					return true
				}
			}
		}

		if se == opponent {
			for d := 2; column + d <= 7 && row + d <= 7; d++ {
				nextPiece := board[row + d][column + d]
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