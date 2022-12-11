package main

import (
	"fmt"
	"log"
	"os"
)

const (
	ROCK = iota
	PAPER
	SCISSORS
)

const (
	LOSS = 3 * (iota - 1)
	DRAW
	WIN
)

var move_mapping = map[string]int{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	score := calculateScore(file)
	fmt.Printf("Total score: %d\n", score)
}

func calculateScore(F *os.File) int {

}
