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

func main() {
	// Challenge 5-1
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	crateConfig, moveConfig := readInput(file)
	file.Close()
	crates := parseCrateConfig(crateConfig)
	moves := parseMoveConfig(moveConfig)
	movedCrates := moveCratesOneByOne(moves, crates)
	topCrates := getTopCrates(movedCrates)
	fmt.Printf("The top crates are: %s\n", topCrates)

	// Challenge 5-2
	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	crateConfig, moveConfig = readInput(file)
	file.Close()
	crates = parseCrateConfig(crateConfig)
	moves = parseMoveConfig(moveConfig)
	movedCrates = moveCratesInGroups(moves, crates)
	topCrates = getTopCrates(movedCrates)
	fmt.Printf("The top crates are: %s\n", topCrates)
}

func readInput(F *os.File) (string, string) {
	var crateConfigBuffer, movesBuffer, currentBuffer bytes.Buffer
	var crateConfig, moves string
	scanner := bufio.NewScanner(F)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			crateConfigBuffer = currentBuffer
			currentBuffer = movesBuffer
			continue
		}
		line = line + "\n"
		currentBuffer.WriteString(line)
	}
	movesBuffer = currentBuffer
	crateConfig = crateConfigBuffer.String()
	moves = movesBuffer.String()
	return crateConfig, moves
}

func parseCrateConfig(crateConfig string) [][]string {
	var crates [][]string
	var emptyStack []string

	crateStrata := strings.Split(crateConfig, "\n")
	// Remove line of stack numbers from configs since it is saved
	// in crateNums
	crateNums := crateStrata[len(crateStrata)-2]
	crateStrata = crateStrata[:len(crateStrata)-2]

	stackCount := getStackCount(crateNums)

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
				crates[stackIdx-1] = crateStack
			}
		}
	}

	return crates
}

func getStackCount(crateNumbers string) int {
	var stackCount int
	st := strings.Split(crateNumbers, "")
	for _, s := range st {
		if _, err := strconv.Atoi(s); err == nil {
			stackCount++
		} else {
			continue
		}
	}
	return stackCount
}

func parseMoveConfig(moveConfig string) [][]int {
	var moves [][]int
	var emptyMove []int
	movesSplit := strings.Split(moveConfig, "\n")

	// read function leaves dangling newline in moves config
	// which is removed here as a redundant line in the split
	// moves config
	movesSplit = movesSplit[:len(movesSplit)-1]
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

func moveCratesOneByOne(moves [][]int, crates [][]string) [][]string {
	for _, move := range moves {
		nCrates := move[0]
		src := move[1]
		dest := move[2]

		srcStack := crates[src]
		destStack := crates[dest]
		for i := 0; i < nCrates; i++ {
			srcStack, destStack = moveSingleCrate(srcStack, destStack)
		}
		crates[src] = srcStack
		crates[dest] = destStack
	}
	return crates
}

func moveSingleCrate(srcStack, destStack []string) ([]string, []string) {
	crate := srcStack[len(srcStack)-1]
	srcStack = srcStack[:len(srcStack)-1]
	destStack = append(destStack, crate)
	return srcStack, destStack
}

func moveCratesInGroups(moves [][]int, crates [][]string) [][]string {
	for _, move := range moves {
		nCrates := move[0]
		src := move[1]
		dest := move[2]

		srcStack := crates[src]
		destStack := crates[dest]
		srcStack, destStack = moveCrateGroup(srcStack, destStack, nCrates)
		crates[src] = srcStack
		crates[dest] = destStack
	}
	return crates
}

func moveCrateGroup(srcStack, destStack []string, numCrates int) ([]string, []string) {
	crates := srcStack[len(srcStack)-numCrates:]
	srcStack = srcStack[:len(srcStack)-numCrates]
	// Crates acts as a queue for the lifted crates
	// which can simply be appended to destStack in order
	// for the desired effect
	destStack = append(destStack, crates...)
	return srcStack, destStack
}

func getTopCrates(crateStacks [][]string) string {
	var topCrates bytes.Buffer
	for _, stack := range crateStacks {
		topCrate := stack[len(stack)-1]
		topCrates.WriteString(topCrate)
	}
	return topCrates.String()
}
