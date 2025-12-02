// Package day01 contains solutions for Avent of Code 2025 Day 1.
package day01

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	inputLines, err := ReadLinesFromInputFile("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rotations, err := parseInstructions(inputLines)
	if err != nil {
		log.Fatal(err)
	}

	dial := newDial()

	partOnePassword := 0
	partTwoPassword := 0

	for _, rotation := range rotations {
		zeroClicks := dial.rotate(rotation)
		if dial.position == 0 {
			partOnePassword++
		}
		partTwoPassword += zeroClicks
	}

	fmt.Printf("Part 1 - The final password is: %d\n", partOnePassword)
	fmt.Printf("Part 2 - The final password is: %d\n", partTwoPassword)
}

func ReadLinesFromInputFile(filepath string) ([]string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

type rotationDirection int

const (
	rotationRight rotationDirection = iota
	rotationLeft
)

type rotation struct {
	direction rotationDirection
	distance  int
}

func newRotation(instruction string) (rotation, error) {
	var direction rotationDirection
	instruction = strings.ToUpper(strings.TrimSpace(instruction))

	switch instruction[0] {
	case 'R':
		direction = rotationRight
	case 'L':
		direction = rotationLeft
	default:
		return rotation{}, errors.New("invalid rotation direction")
	}

	distance, err := strconv.Atoi(instruction[1:])
	if err != nil {
		return rotation{}, errors.New("invalid distance in rotation instruction")
	}

	return rotation{
		direction: direction,
		distance:  distance,
	}, nil
}

func parseInstructions(lines []string) ([]rotation, error) {
	rotations := make([]rotation, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		rotation, err := newRotation(line)
		if err != nil {
			return rotations, err
		}
		rotations = append(rotations, rotation)
	}

	return rotations, nil
}

type dial struct {
	position     int
	numPositions int
}

func newDial() *dial {
	return &dial{position: 50, numPositions: 100}
}

func (d *dial) rotate(r rotation) (zeroClicks int) {
	startPosition := d.position
	remainingDistance := r.distance % d.numPositions
	zeroClicks = r.distance / d.numPositions

	if r.direction == rotationRight {
		d.position = (d.position + remainingDistance) % d.numPositions
	} else {
		d.position = (d.position - remainingDistance + d.numPositions) % d.numPositions
	}
	if r.direction == rotationRight {
		if startPosition+remainingDistance >= d.numPositions {
			zeroClicks++
		}
	} else {
		if remainingDistance >= startPosition && startPosition > 0 {
			zeroClicks++
		}
	}

	rotationWord := map[rotationDirection]string{
		rotationRight: "R",
		rotationLeft:  "L",
	}
	if zeroClicks >= 1 {
		fmt.Printf("The dial is rotated %s%d to point at %d; during this rotation, it points at 0 %d times.\n", rotationWord[r.direction], r.distance, d.position, zeroClicks)
	} else {
		fmt.Printf("The dial is rotated %s%d to point at %d.\n", rotationWord[r.direction], r.distance, d.position)
	}
	return zeroClicks
}
