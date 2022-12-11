package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ROCK = iota
	PAPER
	SCISSORS
)

const (
	LOSS = 3 * iota
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

var rule_mapping = map[int][]int{
	ROCK:     {DRAW, LOSS, WIN},
	PAPER:    {WIN, DRAW, LOSS},
	SCISSORS: {LOSS, WIN, DRAW},
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	score := calculateScore(file)
	file.Close()
	fmt.Printf("Total score: %d\n", score)
}

func calculateScore(F *os.File) int {
	scanner := bufio.NewScanner(F)
	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Split(line, " ")

		// Add 1 to each hand's points since move constants start at 0
		totalPoints += calculateTurnPoints(moves[1], moves[0]) + 1
	}
	return totalPoints
}

func calculateTurnPoints(m1, m2 string) int {
	myMove := move_mapping[m1]
	oppMove := move_mapping[m2]
	myState := rule_mapping[myMove][oppMove]
	return myMove + myState
}
