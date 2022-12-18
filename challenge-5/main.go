package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const crateWidth = 3

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	crateConfig, moveConfig := readInput(file)
	file.Close()
	crates := parseCrateConfig(crateConfig)
	moves := parseMoveConfig(moveConfig)
	topCrates := moveCrates(moves, crates)
	fmt.Printf("The top crates are: %s\n", topCrates)
}

func readInput(F *os.File) (string, string) {
	var crateConfigBuffer, movesBuffer bytes.Buffer
	var crateConfig, moves string
	scanner := bufio.NewScanner(F)
	currentBuffer := crateConfigBuffer
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currentBuffer = movesBuffer
		}
		currentBuffer.WriteString(line)
	}
	crateConfig = crateConfigBuffer.String()
	moves = movesBuffer.String()
	return crateConfig, moves
}

func parseCrateConfig(crateConfig string) [][]string {
	var crates [][]string
	var emptyStack []string
	crateStrata := strings.Split(crateConfig, "\n")
	crateNums := crateStrata[len(crateStrata)]
	stackCount := len(strings.Split(crateNums, " "))

	// Remove line of stack numbers from configs since it is saved
	// in crateNums
	crateStrata = crateStrata[:len(crateStrata)-1]

	for i := 0; i < stackCount; i++ {
		crates = append(crates, emptyStack)
	}

	// The stack each crate should be placed in can be gotten by
	// retrieving the character in crateNums at the same index as
	// the alphabetic char denoting each crate on the current line.
	// Crates are prepended to the 'stack' here since input is read
	// top to bottom
	for _, line := range crateStrata {
		lineBytes := []byte(line)
		for idx, char := range lineBytes {
			if char >= 'A' && char <= 'Z' {
				stackIdx, _ := strconv.Atoi(string(crateNums[idx]))
				crateStack := crates[stackIdx-1]
				crateStack = append([]string{string(char)}, crateStack...)
			}
		}
	}

	return crates
}

func parseMoveConfig(moveConfig string) [][]int {
	var moves [][]int
	var emptyMove []int
	movesSplit := strings.Split(moveConfig, "\n")

	for i := 0; i < len(movesSplit); i++ {
		moves = append(moves, emptyMove)
	}

	// Extracting integer values of indices 1, 3, 5
	// of moves split on whitespace since
	// all instructions are of the form
	// Move X from Y to Z
	for i, move := range movesSplit {
		m := strings.Split(move, " ")
		n, _ := strconv.Atoi(m[1])
		src, _ := strconv.Atoi(m[3])
		dest, _ := strconv.Atoi(m[5])
		moves[i] = append(moves[i], n)
		moves[i] = append(moves[i], src-1)
		moves[i] = append(moves[i], dest-1)
	}

	return moves
}

func moveCrates(moves [][]int, crates [][]string) string {
	var topCrates bytes.Buffer
	for _, move := range moves {
		nCrates := move[0]
		src := move[1]
		dest := move[2]

		srcStack := crates[src]
		destStack := crates[dest]
		for i := 0; i < nCrates; i++ {
			moveCrate(srcStack, destStack)
		}
	}
	return topCrates.String()
}

func moveCrate(srcStack, destStack []string) {
	crate := srcStack[len(srcStack)-1]
	destStack = append(destStack, crate)
}