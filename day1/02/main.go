package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func makeTurn(direction string, clicks int, currentRotation *int) int {
	inc := 0

	switch direction {
	case "R":
		s := *currentRotation % 100
		if s < 0 {
			s += 100
		}
		clicksAtZero := clicks + s

		if s == 0 {
			inc = clicks / 100
		} else {
			if clicksAtZero >= 100 {
				inc = clicksAtZero / 100
			}
		}

		*currentRotation = (s + clicks) % 100

	case "L":
		s := *currentRotation % 100
		if s < 0 {
			s += 100
		}

		if s == 0 {
			inc = clicks / 100
		} else {
			clicksAtZero := clicks - s
			if clicksAtZero >= 0 {
				inc = 1 + clicksAtZero/100
			}
		}

		*currentRotation = (s - (clicks % 100) + 100) % 100
	}

	return inc
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
		password = password + makeTurn(dir, number, &currentRotation)
	}
	fmt.Println(password)
}
