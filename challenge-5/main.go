package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

const crateWidth = 3

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	crateConfig, moves := readInput(file)
	file.Close()
	crates := parseCrateConfig(crateConfig)
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

	// Remove line of craate numbers from configs since it is saved
	// in crateNums
	crateStrata = crateStrata[:len(crateStrata)-1]

	for i := 0; i < stackCount; i++ {
		crates = append(crates, emptyStack)
	}

	for _, line := range crateStrata {
		lineBytes := []byte(line)
		for idx, char := range lineBytes {
			if char >= 'A' && char <= 'Z' {
				crateStack := crates[crateNums[idx]]
				crateStack = append([]string{string(char)}, crateStack...)
			}
		}
	}

	return crates
}
