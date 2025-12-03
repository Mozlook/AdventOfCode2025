package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v", err)
		return
	}
	defer rawInput.Close()
	answer := 0
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		bank := scanner.Text()

		higher := byte('0')
		lower := byte('0')

		for i := 0; i < len(bank); i++ {
			char := bank[i]

			if i < len(bank)-1 {
				if char > higher {
					higher = char
					lower = byte('0')
					continue
				}
			}
			if char > lower {
				lower = char
			}
		}
		sum := int(higher-'0')*10 + int(lower-'0')
		answer = answer + sum
	}

	fmt.Println(answer)
}
