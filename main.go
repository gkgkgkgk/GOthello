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

func main(){
    flag.StringVar(&loadpath, "savedpath", "", "The path to a save file from which to load a game.")
	flag.Parse()

	beginGame()
	gameLoop()
}