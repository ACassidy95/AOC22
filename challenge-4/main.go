package main

import (
	"fmt"
	"log"
	"os"
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
	return redundantRotaCount
}
