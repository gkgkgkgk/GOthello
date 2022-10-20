package main

import (
	"flag"
)

var loadpath string
var defaultSettings bool
var gameOver bool

func main(){
    flag.StringVar(&loadpath, "savedpath", "", "The path to a save file from which to load a game.")
	flag.BoolVar(&defaultSettings, "default", false, "Load default settings: Human player 1, AI player 2, 5 second time limit, game begins with player 1.")
	flag.Parse()

	// for p1/p2 0 is human, 1 is AI
	p1, p2, turn, timeLimit := beginGame()
	gameLoop(p1, p2, turn, timeLimit)
}