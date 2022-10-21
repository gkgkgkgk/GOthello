package main

import (
	"fmt"
	"time"
	"math"
)

func getAITurn(board[8][8] int, legalMoves []int, turn int, timeLimit int) int{
	start := time.Now()

	if len(legalMoves) == 1 { // if there is only one available move
		return 0
	}
	
	bestMove := 0
	depth := 1
	completeDepth := true
	// fmt.Printf("checking %d different turns\n", len(legalMoves))

	for {
		if float64(timeLimit) - time.Since(start).Seconds() < float64(timeLimit)/2.0 {
			break
		}

		value := math.MinInt
		bestMoveDepth := 0
		// fmt.Printf("Depth: %d, found %d moves\n", depth, len(legalMoves))

		for i, move := range legalMoves {
			// fmt.Printf("Placing move on board: %d\n", move)
			tmpBoard := placePiece(board, move, turn)
			tmpTurn := turn

			if turn == 1 {
				tmpTurn = 2
			} else {
				tmpTurn = 1
			}

			score := alphaBeta(tmpBoard, math.MinInt, math.MaxInt, depth, tmpTurn, turn)
			// fmt.Printf("Score: %d\n", score)

			if score > value {
				value = score
				bestMoveDepth = i
			}

			if time.Since(start).Seconds() > float64(timeLimit) {
				completeDepth = false
				break
			}
		}

		if !completeDepth {
			fmt.Printf("AI was interrupted at depth %d.\n", depth)
			break
		}

		bestMove = bestMoveDepth
		depth++
	}

	fmt.Printf("AI reached a depth of %d after %f seconds.\n", depth, time.Since(start).Seconds())
	
	return bestMove
}

func alphaBeta(board[8][8] int, alpha int, beta int, depth int, turn int, maxTurn int) int{
	// check here for terminal state
	if depth == 0 { // also check time
		// fmt.Printf("depth: %d alpha: %d beta: %d\n", depth, alpha, beta)
		return calculateHeuristicScore(board, turn)
	}

	if turn == maxTurn {
		return maxScore(board, alpha, beta, depth, turn, maxTurn)

	} else {
		return minScore(board, alpha, beta, depth, turn, maxTurn)
	}
}

// of course, this is a thing: https://github.com/golang/go/issues/20517
func maxScore(board[8][8] int, alpha int, beta int, depth int, turn int, maxTurn int) int{
	v := math.MinInt
	legalMoves := getAllLegalMoves(turn, board)

	for _, element := range legalMoves {
		tmpBoard := placePiece(board, element, turn)
		
		if turn == 1 {
			turn = 2
		} else {
			turn = 1
		}

		// fmt.Printf("%d and %d\n", v, ab)
		v = maximum(v, alphaBeta(tmpBoard, alpha, beta, depth - 1, turn, maxTurn))

		if v >= beta {
			return v
		}

		alpha = maximum(alpha, v)
	}
	
	return v
}

func minScore(board[8][8] int, alpha int, beta int, depth int, turn int, maxTurn int) int{
	v := math.MaxInt
	legalMoves := getAllLegalMoves(turn, board)

	for _, element := range legalMoves {
		tmpBoard := placePiece(board, element, turn)

		if turn == 1 {
			turn = 2
		} else {
			turn = 1
		}

		v = minimum(v, alphaBeta(tmpBoard, alpha, beta, depth - 1, turn, maxTurn))

		if v <= alpha {
			return v
		}

		beta = minimum(beta, v)
	}
	
	return v
}

// given a board and a player, return a heuristic score for the player
func calculateHeuristicScore(board [8][8] int, player int) int{
	scoreBoard := [8][8]int{
		{200, -10, 10, 10, 10, 10, -10, 200},
		{-10, -10, -1, -1, -1, -1, -1, -10},
		{10, -1, 1, 1, 1, 1, -1, 10},
		{10, -1, 1, 1, 1, 1, -1, 10},
		{10, -1, 1, 1, 1, 1, -1, 10},
		{10, -1, 1, 1, 1, 1, -1, 10},
		{-10, -10, -1, -1, -1, -1, -10, -10},
		{200, -10, 10, 10, 10, 10, -10, 200}}
	
	score := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == player {
				score += scoreBoard[i][j]
			}
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