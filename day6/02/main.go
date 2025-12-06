package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Add(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func Multiply(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	product := 1
	for _, i := range numbers {
		product *= i
	}
	return product
}

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}
	defer rawInput.Close()
	var sheet []string
	answer := 0
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		line := scanner.Text()
		sheet = append(sheet, line)
	}
	n := len(sheet[0])
	var numbers []int
	for i := n - 1; i >= 0; i-- {
		number := ""
		for j := range 5 {
			if sheet[j][i] == byte('+') {
				intNumber, err := strconv.Atoi(number)
				if err != nil {
					fmt.Printf("parsing error %v", err)
				}
				numbers = append(numbers, intNumber)
				answer += Add(numbers)
				fmt.Println(numbers)
				numbers = numbers[:0]
				number = ""
				break
			}
			if sheet[j][i] == byte('*') {
				intNumber, err := strconv.Atoi(number)
				if err != nil {
					fmt.Printf("parsing error %v", err)
				}
				numbers = append(numbers, intNumber)
				answer += Multiply(numbers)
				fmt.Println(numbers)
				numbers = numbers[:0]
				number = ""
				break
			}
			if sheet[j][i] != byte(' ') {
				number += string(sheet[j][i])
			}
		}
		if number != "" {
			intNumber, err := strconv.Atoi(number)
			if err != nil {
				fmt.Printf("parsing error %v", err)
			}
			numbers = append(numbers, intNumber)
		}
	}
	fmt.Println(answer)
}
