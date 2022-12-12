package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Move int
type Result int

const (
	ROCK Move = iota + 1
	PAPER
	SCISSORS
)

const (
	LOSS Result = 3 * iota
	DRAW
	WIN
)

var moveMapping = map[string]Move{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var encodedMoveMapping = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var ruleMapping = map[Move][]Result{
	//          R  ,  P   , S
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
		totalPoints += calculateTurnPoints(moves[1], moves[0])
	}
	return totalPoints
}

func calculateTurnPoints(self, opponent string) int {
	myMove := moveMapping[encodedMoveMapping[self]]
	oppMove := moveMapping[opponent]
	myResult := ruleMapping[myMove][moveValue(oppMove)-1]
	return moveValue(myMove) + resultValue(myResult)
}

func moveValue(m Move) int {
	return int(m)
}

func resultValue(r Result) int {
	return int(r)
}
