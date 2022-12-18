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
	var eopsIdx int
	signalBytes := []byte(signal)
	for i := 0; i < len(signalBytes)-windowSize; i++ {
		window := signalBytes[i : i+windowSize]
		uniqueWindowChars := make(map[byte]bool)
		for _, char := range window {
			// If the current char in the window already exists in the
			// set of characters constructed from the window, then the current
			// window cannot be a protocol start marker and the next window can be checked
			_, ok := uniqueWindowChars[char]
			if ok {
				break
			} else {
				uniqueWindowChars[char] = true
			}
		}
		if len(uniqueWindowChars) == windowSize {
			eopsIdx = i + windowSize
			break
		}
	}
	return eopsIdx
}
