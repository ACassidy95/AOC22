package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	caloriesPerElf, err := countCaloriesPerElf(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	file.Close()
	fattestElfIdx := findElfWithMostCalories(caloriesPerElf)
	fmt.Printf("The fattest elf is elf number: %d. He is carrying %d calories\n", fattestElfIdx+1, caloriesPerElf[fattestElfIdx])
}

func countCaloriesPerElf(F *os.File) ([]int, error) {
	var caloriesPerElf []int
	currentElfCalories := 0
	scanner := bufio.NewScanner(F)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			caloriesPerElf = append(caloriesPerElf, currentElfCalories)
			currentElfCalories = 0
			continue
		}
		calorieVal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		currentElfCalories += calorieVal
	}

	return caloriesPerElf, nil
}

func findElfWithMostCalories(caloriesPerElf []int) int {
	maxElf := 0
	maxCalories := 0
	for elf, elfCalories := range caloriesPerElf {
		if elfCalories > maxCalories {
			maxElf = elf
			maxCalories = elfCalories
		}
	}
	return maxElf
}
