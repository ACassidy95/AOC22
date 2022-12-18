package main

import (
	"bufio"
	"log"
	"os"
)

var windowSize = 4

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	signal := readInput(file)
}

func readInput(F *os.File) string {
	var signal string
	scanner := bufio.NewScanner(F)
	for scanner.Scan() {
		signal += scanner.Text()
	}
	return signal
}
