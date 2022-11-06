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
			if float64(timeLimit) - time.Since(start).Seconds() < 0.35 {
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

			score, cutOff := minimax(tmpBoard, depth, math.MinInt, math.MaxInt, turn, tmpTurn, p, timeLimit, start)

			if !cutOff {
				scores[i] = score
				depths[i] = depth
			}
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

func minimax(board[8][8] int, depth int, a int, b int, maxTurn int, currentTurn int, previousMoves []int, timeLimit int, start time.Time) (int, bool){
	max := currentTurn
	min := 1
	if max == 1 {
		min = 2
	}

	if float64(timeLimit) - time.Since(start).Seconds() < 0.1 {
		return 0, true
	}
	
	
	alpha := a
	beta := b

	if depth == 0 {
		score := calculateHeuristicScore(board, maxTurn)
		return score, false
	}

	

	if len(getAllLegalMoves(max, board)) == 0 && len(getAllLegalMoves(min, board)) == 0 {
		minScore := 0
		maxScore := 0

		for i := 0; i < 8; i++{
			for j := 0; j < 8; j++ {
				if board[i][j] == max {
					maxScore++
				} else if board[i][j] == min {
					minScore++
				}
			}
		}

		if maxScore > minScore {
			return 100000, false
		} else if minScore > maxScore {
			return -100000, false
		}

		return 0, false
	}

	legalMoves := getAllLegalMoves(currentTurn, board)

	if maxTurn == currentTurn {
		maxVal := math.MinInt
		
		if len(legalMoves) == 0 {
			maxVal = calculateHeuristicScore(board, maxTurn)
		}

		for _, move := range(legalMoves) {
			tmpBoard := placePiece(board, move, currentTurn)
			tmpPreviousMoves := append(previousMoves, move)
			tmpTurn := currentTurn

			if tmpTurn == 1 {
				tmpTurn = 2
			} else {
				tmpTurn = 1
			}

			val, _ := minimax(tmpBoard, depth - 1, alpha, beta, maxTurn, tmpTurn, tmpPreviousMoves, timeLimit, start)
			maxVal = maximum(maxVal, val)
			alpha = maximum(alpha, val)

			if beta <= alpha {
				break
			}
		}
		return maxVal, false
	} else {
		minVal := math.MaxInt

		if len(legalMoves) == 0 {
			minVal = calculateHeuristicScore(board, min)
		}
		
		for _, move := range(legalMoves) {
			tmpBoard := placePiece(board, move, currentTurn)
			tmpPreviousMoves := append(previousMoves, move)
			tmpTurn := currentTurn

			if tmpTurn == 1 {
				tmpTurn = 2
			} else {
				tmpTurn = 1
			}

			val, _ := minimax(tmpBoard, depth - 1, alpha, beta, maxTurn, tmpTurn, tmpPreviousMoves, timeLimit, start)
			minVal = minimum(minVal, val)
			beta = minimum(beta, val)

			if beta <= alpha {
				break
			}
		}
		return minVal, false
	}
}

// given a board and a player, return a heuristic score for the player
func calculateHeuristicScore(board [8][8] int, player int) int{
	max := player
	min := 1

	if max == 1 {
		min = 2
	}

	scoreBoard := [8][8]int{
		{10, -5, 5, 5, 5, 5, -5, 10},
		{-5, -7, -1, -1, -1, -1, -7, -5},
		{5, -1, 10, 1, 1, 10, -1, 5},
		{5, -1, 1, 5, 5, 1, -1, 5},
		{5, -1, 1, 5, 5, 1, -1, 5},
		{5, -1, 10, 1, 1, 10, -1, 5},
		{-5, -7, -1, -1, -1, -1, -7, -5},
		{10, -5, 5, 5, 5, 5, -5, 10}}
	
	empties := 0
	maxCorners := 0
	minCorners := 0
	maxPieces := 0
	minPieces := 0

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 0 {
				empties++
			} else if board[i][j] == max {
				maxPieces++
				if empties > 40 {
					maxPieces += scoreBoard[i][j]
				}
			} else {
				minPieces++
				if empties > 40 {
					minPieces += scoreBoard[i][j]
				}
			}

			if (i == 0 && j == 0) || (i == 0 && j == 7) || (i == 7 && j == 7) || (i == 7 && j == 0) {
				if board[i][j] == max {
					maxCorners += 1
				} else if board[i][j] == min {
					minCorners += 1
				} 
			}
		}
	}

	maxLegalMoves := getAllLegalMoves(max, board)
	minLegalMoves := getAllLegalMoves(min, board)

	maxMobility := len(maxLegalMoves)
	minMobility := len(minLegalMoves)

	emptyStop := 0

	if board[0][0] == 0 {
		if board[0][1] == max {
			emptyStop--
		}
		if board[1][1] == max {
			emptyStop--
		}
		if board[1][0] == max {
			emptyStop--
		}
		if board[0][1] == min {
			emptyStop++
		}
		if board[1][1] == min {
			emptyStop++
		}
		if board[1][0] == min {
			emptyStop++
		}
	}
	if board[0][7] == 0 {
		if board[0][6] == max {
			emptyStop--
		}
		if board[1][7] == max {
			emptyStop--
		}
		if board[1][6] == max {
			emptyStop--
		}
		if board[0][6] == min {
			emptyStop++
		}
		if board[1][7] == min {
			emptyStop++
		}
		if board[1][6] == min {
			emptyStop++
		}
	}
	if board[7][7] == 0 {
		if board[6][6] == max {
			emptyStop--
		}
		if board[7][6] == max {
			emptyStop--
		}
		if board[6][7] == max {
			emptyStop--
		}
		if board[6][6] == min {
			emptyStop++
		}
		if board[7][6] == min {
			emptyStop++
		}
		if board[6][7] == min {
			emptyStop++
		}
	}
	if board[7][0] == 0 {
		if board[6][0] == max {
			emptyStop--
		}
		if board[6][1] == max {
			emptyStop--
		}
		if board[7][1] == max {
			emptyStop--
		}
		if board[6][0] == min {
			emptyStop++
		}
		if board[6][1] == min {
			emptyStop++
		}
		if board[7][1] == min {
			emptyStop++
		}
	}


	score := 0
	pieces := 0
	if maxPieces + minPieces != 0 {
		pieces = 100 * (maxPieces - minPieces) / (maxPieces + minPieces)
	}
	mobility := 0
	if maxMobility + minMobility != 0 {
		mobility = 100 * (maxMobility - minMobility) / (maxMobility + minMobility)
	}
	corners := 0
	if maxCorners + minCorners != 0 {
		corners = 100 * (maxCorners - minCorners) / (minCorners + maxCorners)
	}

	if empties > 40 {
		score = corners * 15 + mobility * 3 - pieces * 3 + emptyStop * 50
	} else if empties > 15 {
		score = corners * 10 + mobility * 2  + emptyStop * 50
	} else {
		score = corners * 10 + mobility * 2 + pieces * 10 + emptyStop * 10
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