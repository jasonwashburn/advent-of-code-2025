// Package day02 contains solutions for Avent of Code 2025 Day 2.
package day02

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	inputLine, err := readInputLineFromFile("day02/sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	rangeStrings := strings.Split(inputLine, ",")
	for numIndex, numStr := range rangeStrings {
		fmt.Printf("Num %d: %s\n", numIndex, numStr)
	}

	productIDRanges := make([][2]int, len(rangeStrings))
	for i, rangeString := range rangeStrings {
		splitString := strings.Split(rangeString, "-")
		startStr := splitString[0]
		endStr := splitString[1]

		startNum, err := strconv.Atoi(startStr)
		if err != nil {
			log.Fatal(err)
		}

		endNum, err := strconv.Atoi(endStr)
		if err != nil {
			log.Fatal(err)
		}

		productIDRanges[i] = [2]int{startNum, endNum}

		fmt.Printf("Range: %d to %d\n", productIDRanges[i][0], productIDRanges[i][1])
	}

	invalidIDSum := 0
	for _, productIDRange := range productIDRanges {
		rangeStart := productIDRange[0]
		rangeEnd := productIDRange[1]
		for productID := rangeStart; productID <= rangeEnd; productID++ {
			if isSillyPattern(productID) {
				fmt.Printf("Silly pattern found: %d\n", productID)
				invalidIDSum += productID
			}
		}
	}

	fmt.Printf("Sum of invalid product IDs: %d\n", invalidIDSum)
}

func readInputLineFromFile(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	inputLine := strings.ReplaceAll(string(content), "\n", "")
	return inputLine, nil
}

func isSillyPattern(n int) bool {
	numStr := strconv.Itoa(n)

	splitIndex := len(numStr) / 2
	firstHalf := numStr[:splitIndex]
	secondHalf := numStr[splitIndex:]

	return firstHalf == secondHalf
}
