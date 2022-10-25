package main

import (
	"fmt"
	// "time"
	"math"
)

func getAITurn(board[8][8] int, legalMoves []int, turn int, timeLimit int) int{
	fmt.Printf("Getting AI turn. There are %d available moves.\n", len(legalMoves))
	// LOOK INTO REORDERING MOVES FROM BEST TO WORSE TO INCREASE PRUNING PERFORMANCE
	depth := 2
	var scores []int

	for _, move := range legalMoves {
		fmt.Printf("Starting move %s.\n", convertIntToCoords(move))
		var p = []int{move}
		tmpBoard := placePiece(board, move, turn)
		tmpTurn := turn
		if turn == 1 {
			tmpTurn = 2
		} else {
			tmpTurn = 1
		}
		score := minimax(tmpBoard, depth, math.MinInt, math.MaxInt, turn, tmpTurn, p)
		fmt.Printf("Final score: %d\n", score)
		scores = append(scores, score)
	}
	
	bestI := 0
	bestScore := math.MinInt
	for i, score := range(scores){
		if score >= bestScore {
			bestScore = score
			bestI = i
		}
	}

	return bestI
	// start := time.Now()

	// if len(legalMoves) == 1 { // if there is only one available move
	// 	return 0
	// }
	
	// bestMove := 0
	// depth := 1
	// completeDepth := true
	// // fmt.Printf("checking %d different turns\n", len(legalMoves))

	// for {
	// 	if float64(timeLimit) - time.Since(start).Seconds() < float64(timeLimit)/2.0 {
	// 		break
	// 	}

	// 	value := math.MinInt
	// 	bestMoveDepth := 0
	// 	// fmt.Printf("Depth: %d, found %d moves\n", depth, len(legalMoves))

	// 	for i, move := range legalMoves {
	// 		// fmt.Printf("Placing move on board: %d\n", move)
	// 		tmpBoard := placePiece(board, move, turn)
	// 		tmpTurn := turn

	// 		if turn == 1 {
	// 			tmpTurn = 2
	// 		} else {
	// 			tmpTurn = 1
	// 		}

	// 		score := alphaBeta(tmpBoard, math.MinInt, math.MaxInt, depth, tmpTurn, turn)
	// 		// fmt.Printf("Score: %d\n", score)

	// 		if score > value {
	// 			value = score
	// 			bestMoveDepth = i
	// 		}

	// 		if time.Since(start).Seconds() > float64(timeLimit) {
	// 			completeDepth = false
	// 			break
	// 		}
	// 	}

	// 	if !completeDepth {
	// 		fmt.Printf("AI was interrupted at depth %d.\n", depth)
	// 		break
	// 	}

	// 	bestMove = bestMoveDepth
	// 	depth++
	// }

	// fmt.Printf("AI reached a depth of %d after %f seconds.\n", depth, time.Since(start).Seconds())
	
	// return bestMove
}

func minimax(board[8][8] int, depth int, alpha int, beta int, maxTurn int, currentTurn int, previousMoves []int) (int){
	if depth == 0 {
		score := calculateHeuristicScore(board, maxTurn)
		// fmt.Printf("Best heuristic had a score of %d.\n", score)
		// fmt.Println(previousMoves)
		return score
	}

	if maxTurn == currentTurn {
		// fmt.Println("Max turn.")
		maxVal := math.MinInt
		legalMoves := getAllLegalMoves(currentTurn, board)

		for _, move := range(legalMoves) {
			// fmt.Printf("Checking turn %s\n", convertIntToCoords(move))
			tmpBoard := placePiece(board, move, currentTurn)
			tmpPreviousMoves := append(previousMoves, move)
			tmpTurn := currentTurn

			if tmpTurn == 1 {
				tmpTurn = 2
			} else {
				tmpTurn = 1
			}

			val := minimax(tmpBoard, depth - 1, alpha, beta, maxTurn, tmpTurn, tmpPreviousMoves)
			// fmt.Printf("found val %d\n", val)
			maxVal = maximum(maxVal, val)
			alpha = maximum(alpha, val)

			if beta <= alpha {
				break
			}
		}
		return maxVal
	} else {
		// fmt.Println("Min turn.")
		minVal := math.MaxInt
		legalMoves := getAllLegalMoves(currentTurn, board)

		for _, move := range(legalMoves) {
			// fmt.Printf("Checking turn %s\n", convertIntToCoords(move))
			tmpBoard := placePiece(board, move, currentTurn)
			tmpPreviousMoves := append(previousMoves, move)
			tmpTurn := currentTurn

			if tmpTurn == 1 {
				tmpTurn = 2
			} else {
				tmpTurn = 1
			}

			val := minimax(tmpBoard, depth - 1, alpha, beta, maxTurn, tmpTurn, tmpPreviousMoves)
			// fmt.Printf("found val %d\n", val)
			minVal = minimum(minVal, val)
			beta = minimum(beta, val)

			if beta <= alpha {
				break
			}
		}
		return minVal
	}
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