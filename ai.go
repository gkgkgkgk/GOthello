package main

import (
	"fmt"
	"time"
	"math"
)

func getAITurn(legalMoves []int) int{
	start := time.Now()

	fmt.Printf("%f", time.Since(start).Seconds())
	return 0
}

func alphaBeta(alpha int, beta int, depth int, board[8][8] int, maxPlayer bool, turn int) int{
	heuristic := calculateHeuristic(board, turn)

	if depth == 0 || isTerminalNode(board, turn){
		return heuristic
	}

	if turn == 1 {
		turn = 2
	} else {
		turn = 1
	}

	if maxPlayer {
		value := int(math.Inf(-1))
		childrenNodes := getChildrenNodes(board)

		for i := 0; i < len(childrenNodes); i++{
			value := maximum(value, alphaBeta(alpha, beta, depth -1, childrenNodes[i], false, turn))
			if value >= beta {
				break
			}
	
			alpha = maximum(alpha, value)
		}
		
		return value
	} else {
		value := int(math.Inf(1))
		childrenNodes := getChildrenNodes(board)

		for i := 0; i < len(childrenNodes); i++{
			value := minimum(value, alphaBeta(alpha, beta, depth -1, childrenNodes[i], true,  turn))
			if value <= alpha {
				break
			}
	
			beta = minimum(beta, value)
		}
		
		return value
	}
}

func getChildrenNodes(board[8][8] int) [][8][8] int {
	var children [][8][8]int
	children = append(children, board)

	return children
}

func isTerminalNode(board[8][8] int, turn int) bool{
	if len(getAllLegalMoves(turn, board)) == 0{
		return true
	}
	
	return true
}

func calculateHeuristic(board [8][8] int, player int) int{
	score := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == player {
				score++
			} 
		}
	}

	if board[0][0] == player {
		score += 100
	}

	if board[0][7] == player {
		score += 100
	}

	if board[7][7] == player {
		score += 100
	}

	if board[7][0] == player {
		score += 100
	}

	for i := 0; i < 8; i++{
		if board[i][0] == player {
			score += 100
		}
	}

	return score
}