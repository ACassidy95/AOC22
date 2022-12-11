package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var caloriesPerElf []int
	inFile := os.Args[1:]
	file, err := os.Open(inFile[0])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = countCaloriesPerElf(file, caloriesPerElf)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	file.Close()
	fmt.Printf("Calories per elf: %q\n", caloriesPerElf)
	fattestElf := findElfWithMostCalories(caloriesPerElf)
	fmt.Printf("The fattest elf is elf number: %d\n", fattestElf)
}

func countCaloriesPerElf(F *os.File, caloriesPerElf []int) error {
	scanner := bufio.NewScanner(F)
	currentElfCalories := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			caloriesPerElf = append(caloriesPerElf, currentElfCalories)
			currentElfCalories = 0
			continue
		}
		calorieVal, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
			return err
		}
		currentElfCalories += calorieVal
	}

	return nil
}

func findElfWithMostCalories(caloriesPerElf []int) int {
	max := 0
	for elf, elfCalories := range caloriesPerElf {
		if elfCalories > max {
			max = elf
		}
	}
	return max
}
