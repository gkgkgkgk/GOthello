package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func getAITurn(board[8][8] int, legalMoves []int, turn int, timeLimit int) int{
	fmt.Printf("Getting AI turn.\nThere are %d available moves: ", len(legalMoves))
	fmt.Printf("[")
	for i, move := range legalMoves {
		if i == len(legalMoves) - 1 {
			fmt.Printf("(%s)", convertIntToCoords(move))
		} else {
			fmt.Printf("(%s), ", convertIntToCoords(move))
		}
	}
	fmt.Printf("]\n")

	start := time.Now()

	if len(legalMoves) == 1 { // if there is only one available move
		return 0
	}

	depth := 0
	bestI := 0
	scores := make([]int, len(legalMoves))
	depths := make([]int, len(legalMoves))
	cutOff := false

	out:
	for {
		if float64(timeLimit) - time.Since(start).Seconds() < float64(timeLimit)/2.0 {
			break out
		}

		for i, move := range legalMoves {
			if float64(timeLimit) - time.Since(start).Seconds() < 0.25 {
				cutOff = true
				break out
			}

			var p = []int{move}
			tmpBoard := placePiece(board, move, turn)
			tmpTurn := turn
			if turn == 1 {
				tmpTurn = 2
			} else {
				tmpTurn = 1
			}

			score := minimax(tmpBoard, depth, math.MinInt, math.MaxInt, turn, tmpTurn, p)
			scores[i] = score
			depths[i] = depth
		}

		depth++
	}

	bestScore := math.MinInt
	for i, score := range(scores){
		if score > bestScore {
			bestScore = score
			bestI = i
		} else if score == bestScore && rand.Float64() > 0.5{
			bestScore = score
			bestI = i
		}
	}
	
	fmt.Printf("Scores evaluated to be %v. Finished in %f seconds.\n", scores, time.Since(start).Seconds())

	if cutOff {
		fmt.Printf("Cut off at depth %d. Turn selected from depth %d", depth, depths[bestI])
	} else {
		fmt.Printf("Successfully evaluated until depth %d.", depth)
	}

	return bestI
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
		{200, -10, 15, 10, 10, 15, -10, 200},
		{-10, -10, -1, -1, -1, -1, -1, -10},
		{15, -1, 1, 1, 1, 1, -1, 15},
		{10, -1, 1, 1, 1, 1, -1, 10},
		{10, -1, 1, 1, 1, 1, -1, 10},
		{15, -1, 1, 1, 1, 1, -1, 15},
		{-10, -10, -1, -1, -1, -1, -10, -10},
		{200, -10, 15, 10, 10, 15, -10, 200}}

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