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
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	redundantRotaCount := findRedundantRotas(file)
	file.Close()
	fmt.Printf("The number of fully redundant rotas is: %d", redundantRotaCount)
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
		totalOverlap := rotaTotallyRedundant(elfRotaALowerBound, elfRotaAUpperBound, elfRotaBLowerBound, elfRotaBUpperBound)
		if totalOverlap {
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
	return totalOverlap
}
