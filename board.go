package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"io/ioutil"
	"log"
	"strconv"
	"regexp"
)

// 0 is empty, 1 is black, 2 is white
// the board is addressed as board[row_number (y coordinate)][column_number (x coordinate)]
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

func convertStringToInt(str string) (int, error) {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	num, err := strconv.Atoi(nonAlphanumericRegex.ReplaceAllString(str, ""))
	
	if err != nil{
		return -1, err
	}
	
	return num, nil
}

func printBoard(){
	color.Set(color.FgBlue)
	fmt.Printf("It is Player %d's turn.\n", turn)
	color.Unset()
	fmt.Print("   ")
	for i:=0; i < 8; i++{
		fmt.Print(i)
		fmt.Print("     ")
	}
	fmt.Println("")

	for row := 0; row < 8; row++ {
		color.Set(color.BgGreen)
		color.Set(color.FgBlack)
		for i:=0; i < 49; i++{
			fmt.Print("-")
		}
		fmt.Println("")
		for i:=0; i < 2; i++ {
			color.Set(color.BgGreen)
			color.Set(color.FgBlack)

			fmt.Print("|")
			for column:=0; column < 8; column++{
				if board[row][column] == 0 {
					color.Set(color.BgGreen)
				} else if board[row][column] == 1 {
					color.Set(color.BgBlack)
				} else {
					color.Set(color.BgWhite)
				}
				fmt.Print("     ")
				color.Unset()
				color.Set(color.BgGreen)
				color.Set(color.FgBlack)
				fmt.Print("|")
				color.Unset()

			}
			if i == 0 {
				fmt.Printf(" %d", row)
			}
			fmt.Println("")
		}
	}
	color.Set(color.BgGreen)
	color.Set(color.FgBlack)

	for i:=0; i < 49; i++{
		fmt.Print("-")
	}
	fmt.Println("")
	color.Unset()
}