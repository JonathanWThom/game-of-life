package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	var game Game
	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	game.Init(width, height)
}
