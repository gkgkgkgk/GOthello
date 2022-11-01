package main


// func minimax(board[8][8] int, depth int, alpha int, beta int, maxTurn int, currentTurn int, previousMoves []int) (int){
// 	if depth == 0 {
// 		score := calculateHeuristicScore(board, maxTurn)
// 		// fmt.Printf("Best heuristic had a score of %d.\n", score)
// 		// fmt.Println(previousMoves)
// 		return score
// 	}

// 	if maxTurn == currentTurn {
// 		// fmt.Println("Max turn.")
// 		maxVal := math.MinInt
// 		legalMoves := getAllLegalMoves(currentTurn, board)

// 		for _, move := range(legalMoves) {
// 			// fmt.Printf("Checking turn %s\n", convertIntToCoords(move))
// 			tmpBoard := placePiece(board, move, currentTurn)
// 			tmpPreviousMoves := append(previousMoves, move)
// 			tmpTurn := currentTurn

// 			if tmpTurn == 1 {
// 				tmpTurn = 2
// 			} else {
// 				tmpTurn = 1
// 			}

// 			val := minimax(tmpBoard, depth - 1, alpha, beta, maxTurn, tmpTurn, tmpPreviousMoves)
// 			// fmt.Printf("found val %d\n", val)
// 			maxVal = maximum(maxVal, val)
// 			alpha = maximum(alpha, val)

// 			if beta <= alpha {
// 				break
// 			}
// 		}
// 		return maxVal
// 	} else {
// 		// fmt.Println("Min turn.")
// 		minVal := math.MaxInt
// 		legalMoves := getAllLegalMoves(currentTurn, board)

// 		for _, move := range(legalMoves) {
// 			// fmt.Printf("Checking turn %s\n", convertIntToCoords(move))
// 			tmpBoard := placePiece(board, move, currentTurn)
// 			tmpPreviousMoves := append(previousMoves, move)
// 			tmpTurn := currentTurn

// 			if tmpTurn == 1 {
// 				tmpTurn = 2
// 			} else {
// 				tmpTurn = 1
// 			}

// 			val := minimax(tmpBoard, depth - 1, alpha, beta, maxTurn, tmpTurn, tmpPreviousMoves)
// 			// fmt.Printf("found val %d\n", val)
// 			minVal = minimum(minVal, val)
// 			beta = minimum(beta, val)

// 			if beta <= alpha {
// 				break
// 			}
// 		}
// 		return minVal
// 	}
// }