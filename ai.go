package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func getAITurn(board[8][8] int, legalMoves []int, turn int, timeLimit int) int{
	if len(legalMoves) == 1 { // if there is only one avaailable move
		return legalMoves[0]
	}
	
	var scores []int
	branchTime := timeLimit / len(legalMoves) 

	for i := range legalMoves {
		start := time.Now()
		tmpBoard := placePiece(board, legalMoves[i], turn)

		score := alphaBeta(tmpBoard, int(math.Inf(-1)), int(math.Inf(1)), 1, false, turn)
		scores = append(scores, score)
		fmt.Printf("%f %f", time.Since(start).Seconds(), branchTime)
	}

	bestScore := int(math.Inf(-1))
	for _, element := range scores {
		if element > bestScore {
			bestScore = element
		}
	}

	var indices []int
	for i, element := range scores {
		if element == bestScore {
			indices = append(indices, i)
		}
	}

	if len(indices) == 1 {
		return indices[0]
	} 
	
	return indices[rand.Intn(len(indices))]
}

func alphaBeta(board[8][8] int, maxDepth int, turn int) int{
	return maxScore(board, int(math.Inf(-1)), int(math.Inf(1)), maxDepth)
}

func maxScore(board[8][8] int, alpha int, beta int, depth int, turn int) int{
	legalMoves := getAllLegalMoves(turn, board)
	v := math.Inf(-1)

	for i, element := range legalMoves {
		tmpBoard := placePiece(board, element, turn)
	}
	
	return 0
}

func minScore(board[8][8] int, alpha int, beta int, depth int, color int) int{
	return 0
}

// given a board and a player, return a heuristic score for the player
func calculateHeuristicScore(board [8][8] int, player int) int{
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
			score += 10
		}

		if board[0][i] == player {
			score += 10
		}

		if board[i][7] == player {
			score += 10
		}

		if board[7][i] == player {
			score += 10
		}
	}

	return score
}

func getChildrenNodes(board[8][8] int) [][8][8] int {
	var children [][8][8]int
	children = append(children, board)

	return children
}

// this is probably wrong
func isTerminalNode(board[8][8] int, turn int) bool{
	if len(getAllLegalMoves(turn, board)) == 0{
		return true
	}
	
	return true
}