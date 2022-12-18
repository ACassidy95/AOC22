package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const PACKET_WINDOW_SIZE = 4
const MESSAGE_WINDOW_SIZE = 14

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	signal := readInput(file)
	packetStart := findEndOfProtocolPortionStartMarker(signal, PACKET_WINDOW_SIZE)
	messageStart := findEndOfProtocolPortionStartMarker(signal, MESSAGE_WINDOW_SIZE)
	fmt.Printf("Packet starts after character %d\nMessage starts after character %d\n", packetStart, messageStart)
}

func readInput(F *os.File) []byte {
	var signal string
	scanner := bufio.NewScanner(F)
	for scanner.Scan() {
		signal += scanner.Text()
	}
	return []byte(signal)
}

func findEndOfProtocolPortionStartMarker(signal []byte, windowSize int) int {
	var eopsIdx int
	for i := 0; i < len(signal)-windowSize; i++ {
		window := signal[i : i+windowSize]
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
