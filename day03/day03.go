// Package day03 contains solutions for Advent of Code 2025 Day 3.
package day03

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	inputLines, err := readInputLinesFromFile("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalJoltage := 0
	for lineIndex, line := range inputLines {
		if line == "" {
			continue
		}

		batteries := getBatteriesFromBankStr(line)

		highIndex, highestNum := findHighestNumber(batteries[:len(batteries)-1])
		_, secondHighestNum := findHighestNumber(batteries[highIndex+1:])

		bankJoltage := (highestNum * 10) + secondHighestNum
		fmt.Printf("Line %d: Highest: %d, Second Highest: %d, Bank Joltage: %d\n", lineIndex+1, highestNum, secondHighestNum, bankJoltage)
		totalJoltage += bankJoltage

	}
	fmt.Println("Total Joltage:", totalJoltage)
}

func readInputLinesFromFile(filepath string) ([]string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(content), "\n"), nil
}

func getBatteriesFromBankStr(s string) []int {
	s = strings.TrimSpace(s)
	batteries := make([]int, len(s))
	for i, r := range s {
		d, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}

		batteries[i] = d
	}
	return batteries
}

func findHighestNumber(numbers []int) (index int, value int) {
	highest := -1
	for i, n := range numbers {
		if n > highest {
			highest = n
			index = i
		}
	}
	return index, highest
}
