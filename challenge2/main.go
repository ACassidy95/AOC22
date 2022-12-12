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

var stateMapping = map[Move][]Move{
	//           L  ,  D  ,  W
	ROCK:     {PAPER, ROCK, SCISSORS},
	PAPER:    {SCISSORS, PAPER, ROCK},
	SCISSORS: {ROCK, SCISSORS, PAPER},
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	encodedMoveScore := calculateScoreFromEncodedMove(file)
	desiredResultScore := calculateScoreFromDesiredResult(file)
	file.Close()
	fmt.Printf("Total score from encoded moves: %d\n", encodedMoveScore)
	fmt.Printf("Total score from desired result: %d\n", desiredResultScore)
}

func calculateScoreFromEncodedMove(F *os.File) int {
	scanner := bufio.NewScanner(F)
	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Split(line, " ")
		selfMove := moves[1]
		oppMove := moves[0]
		totalPoints += calculateTurnPoints(selfMove, oppMove)
	}
	return totalPoints
}

func calculateScoreFromDesiredResult(F *os.File) int {

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
