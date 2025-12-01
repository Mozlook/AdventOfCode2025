package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func makeTurn(direction string, clicks int, currentRotation *int) bool {
	switch direction {
	case "R":
		*currentRotation = (*currentRotation + clicks) % 100
	case "L":
		*currentRotation = (*currentRotation - clicks) % 100
	}
	if *currentRotation == 0 {
		return true
	}
	return false
}

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v", err)
		return
	}
	defer rawInput.Close()

	scanner := bufio.NewScanner(rawInput)
	password := 0
	currentRotation := 50
	for scanner.Scan() {
		line := scanner.Text()
		dir := string(line[0])
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("parsing error: %v", err)
			return
		}
		if makeTurn(dir, number, &currentRotation) {
			password++
		}
	}
	fmt.Println(password)
}
