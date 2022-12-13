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

var desiredResultMapping = map[string]Result{
	"X": LOSS,
	"Y": DRAW,
	"Z": WIN,
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

var moveToAchieveResultMapping = map[Move][]Move{
	//           L  ,  D  ,  W
	ROCK:     {SCISSORS, ROCK, PAPER},
	PAPER:    {ROCK, PAPER, SCISSORS},
	SCISSORS: {PAPER, SCISSORS, ROCK},
}

func main() {
	// Challenge 2-1
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	encodedMoveScore := calculateScoreFromEncodedMove(file)
	file.Close()
	fmt.Printf("Total score from encoded moves: %d\n", encodedMoveScore)

	// Challenge 2-2
	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	desiredResultScore := calculateScoreFromDesiredResult(file)
	file.Close()
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
		totalPoints += calculateTurnPointsFromMoves(selfMove, oppMove)
	}
	return totalPoints
}

func calculateScoreFromDesiredResult(F *os.File) int {
	scanner := bufio.NewScanner(F)
	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()
		turn := strings.Split(line, " ")
		desiredResult := turn[1]
		oppMove := turn[0]
		totalPoints += calculateTurnPointsFromDesiredResult(desiredResult, oppMove)
	}
	return totalPoints
}

func calculateTurnPointsFromMoves(self, opponent string) int {
	myMove := moveMapping[encodedMoveMapping[self]]
	oppMove := moveMapping[opponent]
	myResult := ruleMapping[myMove][moveValue(oppMove)-1]
	return moveValue(myMove) + resultValue(myResult)
}

func calculateTurnPointsFromDesiredResult(result, opponentMove string) int {
	desiredResult := desiredResultMapping[result]
	oppMove := moveMapping[opponentMove]

	// Dividing the desired result value by 3 yields the index of the move to take
	// as given in moveToAchieveResultMapping
	resultToMoveIdx := resultValue(desiredResult) / 3
	myMove := moveToAchieveResultMapping[oppMove][resultToMoveIdx]
	return moveValue(myMove) + resultValue(desiredResult)
}

func moveValue(m Move) int {
	return int(m)
}

func resultValue(r Result) int {
	return int(r)
}
