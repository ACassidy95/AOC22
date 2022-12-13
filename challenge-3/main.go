package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// Challenge 3-2
	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	prioritySum = prioritiseGroupBadges(file)
	file.Close()
	fmt.Printf("Sum of group badge priorities: %d\n", prioritySum)
}

func prioritiseRucksacks(F *os.File) int {
	scanner := bufio.NewScanner(F)
	prioritySum := 0
	for scanner.Scan() {
		rucksackContents := []byte(scanner.Text())
		compartmentAContents := rucksackContents[:len(rucksackContents)/2]
		compartmentBContents := rucksackContents[len(rucksackContents)/2:]

		compartmentASet := contentsSet(compartmentAContents)
		compartmentBSet := contentsSet(compartmentBContents)
		compartmentSetIntersection := contentSetIntersection(compartmentASet, compartmentBSet)

		for k := range compartmentSetIntersection {
			prioritySum += calculateItemTypePriority(k)
		}
	}
	return prioritySum
}

func prioritiseGroupBadges(F *os.File) int {
	var currentGroup [][]byte
	scanner := bufio.NewScanner(F)
	prioritySum := 0
	for scanner.Scan() {
		rucksack := []byte(scanner.Text())
		currentGroup = append(currentGroup, rucksack)
		if len(currentGroup) == 3 {
			groupBadge := findGroupBadge(currentGroup)
			prioritySum += calculateItemTypePriority(groupBadge)
			currentGroup = nil
		}
	}
	return prioritySum
}

func findGroupBadge(rucksackGroup [][]byte) byte {
	var groupBadge byte
	rucksackGroupIntersection := make(map[byte]bool)

	// Iteratively intersect sets constructed from the rucksack contents
	// in order to find the common element across all rucksacks in group
	for _, rucksack := range rucksackGroup {
		rucksackSet := contentsSet(rucksack)
		rucksackGroupIntersection = contentSetIntersection(rucksackGroupIntersection, rucksackSet)
	}

	// Iterate over intersection to assign badge value to var
	// (it is expected this map will have only 1 element)
	for k := range rucksackGroupIntersection {
		groupBadge = k
	}

	return groupBadge
}

func contentsSet(contents []byte) map[byte]bool {
	compartmentSet := make(map[byte]bool)
	for _, item := range contents {
		_, ok := compartmentSet[item]
		if ok {
			continue
		}
		compartmentSet[item] = true
	}
	return compartmentSet
}

func contentSetIntersection(c1, c2 map[byte]bool) map[byte]bool {
	intersection := make(map[byte]bool)

	if len(c1) < len(c2) {
		c1, c2 = c2, c1
	}

	for k := range c1 {
		if c2[k] {
			intersection[k] = true
		}
	}
	return intersection
}

func calculateItemTypePriority(itemType byte) int {
	itemValOffset := 0
	if byte('A') <= itemType && itemType <= byte('Z') {
		itemValOffset = 38
	} else if byte('a') <= itemType && itemType <= byte('z') {
		itemValOffset = 96
	}
	return int(itemType) - itemValOffset
}
