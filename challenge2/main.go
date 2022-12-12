package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Move int

const (
	ROCK Move = iota + 1
	PAPER
	SCISSORS
)

const (
	LOSS = 3 * iota
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

var ruleMapping = map[Move][]int{
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
		totalPoints += calculateTurnPoints(moves[1], moves[0])
	}
	return totalPoints
}

func calculateTurnPoints(m1, m2 string) int {
	myMove := moveMapping[encodedMoveMapping[m1]]
	oppMove := moveMapping[m2]
	myState := ruleMapping[myMove][moveValue(oppMove)-1]
	return moveValue(myMove) + myState
}

func moveValue(m Move) int {
	intVal := int(m)
	return intVal
}
