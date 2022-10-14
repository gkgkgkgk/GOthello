package main

import (
	"flag"
)


var timeLimit int
var turn int

// 0 is human, 1 is AI
var p1 int
var p2 int
var loadpath string
var defaultSettings bool
var gameOver bool

func main(){
    flag.StringVar(&loadpath, "savedpath", "", "The path to a save file from which to load a game.")
	flag.BoolVar(&defaultSettings, "default", false, "Load default settings: Human player 1, AI player 2, 5 second time limit, game begins with player 1.")
	flag.Parse()

	beginGame()
	gameLoop()
}