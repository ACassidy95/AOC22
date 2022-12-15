package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Challenge 4-1
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	redundantRotaCount := findFullyRedundantRotas(file)
	file.Close()
	fmt.Printf("The number of fully redundant rotas is: %d\n", redundantRotaCount)

	// Challenge 4-2
	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	partiallyRedundantRotaCount := findRedundantRotas(file)
	file.Close()
	fmt.Printf("The number of fully redundant rotas is: %d\n", partiallyRedundantRotaCount)
}

func findFullyRedundantRotas(F *os.File) int {
	redundantRotaCount := 0
	scanner := bufio.NewScanner(F)
	for scanner.Scan() {
		rota := strings.Split(scanner.Text(), ",")
		elfRotaA := rota[0]
		elfRotaB := rota[1]
		elfRotaALowerBound, elfRotaAUpperBound := convertRotaToIntBounds(elfRotaA)
		elfRotaBLowerBound, elfRotaBUpperBound := convertRotaToIntBounds(elfRotaB)
		totalOverlap := rotaTotallyRedundant(elfRotaALowerBound, elfRotaAUpperBound, elfRotaBLowerBound, elfRotaBUpperBound)
		if totalOverlap {
			redundantRotaCount++
		}
	}
	return redundantRotaCount
}

func findRedundantRotas(F *os.File) int {
	redundantRotaCount := 0
	scanner := bufio.NewScanner(F)
	for scanner.Scan() {
		rota := strings.Split(scanner.Text(), ",")
		elfRotaA := rota[0]
		elfRotaB := rota[1]
		elfRotaALowerBound, elfRotaAUpperBound := convertRotaToIntBounds(elfRotaA)
		elfRotaBLowerBound, elfRotaBUpperBound := convertRotaToIntBounds(elfRotaB)
		overlap := rotaRedundant(elfRotaALowerBound, elfRotaAUpperBound, elfRotaBLowerBound, elfRotaBUpperBound)
		if overlap {
			redundantRotaCount++
		}
	}
	return redundantRotaCount
}

func convertRotaToIntBounds(rota string) (int, int) {
	rotaBoundInts := strings.Split(rota, "-")

	// Don't Do What Donny Don't Does
	// 1. Ignore errors
	// 2. Trust input makers to not put in errors
	rotaLowerBound, _ := strconv.Atoi(rotaBoundInts[0])
	rotaUpperBound, _ := strconv.Atoi(rotaBoundInts[1])
	return rotaLowerBound, rotaUpperBound
}

func rotaTotallyRedundant(lowerBoundA, upperBoundA, lowerBoundB, upperBoundB int) bool {
	totalOverlap := true
	rotaAWidth := upperBoundA - lowerBoundA
	rotaBWidth := upperBoundB - lowerBoundB
	if rotaAWidth <= rotaBWidth {
		if !(lowerBoundB <= lowerBoundA && upperBoundA <= upperBoundB) {
			totalOverlap = false
		}
	} else {
		if !(lowerBoundA <= lowerBoundB && upperBoundB <= upperBoundA) {
			totalOverlap = false
		}
	}
	return totalOverlap
}

func rotaRedundant(lowerBoundA, upperBoundA, lowerBoundB, upperBoundB int) bool {
	partialOverlap := true
	if !(upperBoundB < lowerBoundA || lowerBoundB > upperBoundA) {
		partialOverlap = false
	}
	return partialOverlap
}
