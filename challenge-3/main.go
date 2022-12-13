package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Challenge 3-1
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	prioritySum := prioritiseRucksacks(file)
	file.Close()
	fmt.Printf("Sum of item priorities: %d\n", prioritySum)
}

func prioritiseRucksacks(F *os.File) int {
	scanner := bufio.NewScanner(F)
	prioritySum := 0
	for scanner.Scan() {
		rucksackContents := strings.Split(scanner.Text(), "")
		compartmentAContents := rucksackContents[:len(rucksackContents)/2]
		compartmentBContents := rucksackContents[len(rucksackContents)/2:]
		fmt.Printf("Compartment A: %x\nCompartment B: %x\n", compartmentAContents, compartmentBContents)
	}
	return prioritySum
}
