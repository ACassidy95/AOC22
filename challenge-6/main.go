package main

import (
	"bufio"
	"fmt"
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
	packetStart := findEndOfPacketStartMarker(signal)
	fmt.Printf("Packet starts after character %d\n", packetStart)
}

func readInput(F *os.File) string {
	var signal string
	scanner := bufio.NewScanner(F)
	for scanner.Scan() {
		signal += scanner.Text()
	}
	return signal
}

func findEndOfPacketStartMarker(signal string) int {
	return 1
}
