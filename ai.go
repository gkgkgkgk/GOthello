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

func alphaBeta(alpha int, beta int, depth int, board[8][8] int, maxPlayer bool) int{
	heuristic := calculateHeuristic()

	if depth == 0 || isTerminalNode(board){
		return heuristic
	}

	if maxPlayer {
		value := int(math.Inf(-1))
		childrenNodes := getChildrenNodes(board)

		for i := 0; i < len(childrenNodes); i++{
			value := maximum(value, alphaBeta(alpha, beta, depth -1, childrenNodes[i], false))
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
			value := minimum(value, alphaBeta(alpha, beta, depth -1, childrenNodes[i], true))
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

func isTerminalNode(board[8][8] int) bool{
	if len(getAllLegalMoves(turn, board)) == 0{
		return true
	}
	
	return true
}

func calculateHeuristic() int{
	score := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == turn {
				score++
			} 
		}
	}

	return score
}