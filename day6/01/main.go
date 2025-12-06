package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}
	defer rawInput.Close()
	var sheet [][]string
	answer := 0
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		sheet = append(sheet, fields)
	}
	n := len(sheet[0])

	for i := range n {
		first, err := strconv.Atoi(string(sheet[0][i]))
		if err != nil {
			fmt.Printf("parsing to int error %v", err)
		}
		second, err := strconv.Atoi(string(sheet[1][i]))
		if err != nil {
			fmt.Printf("parsing to int error %v", err)
		}
		third, err := strconv.Atoi(string(sheet[2][i]))
		if err != nil {
			fmt.Printf("parsing to int error %v", err)
		}
		fourth, err := strconv.Atoi(string(sheet[3][i]))
		if err != nil {
			fmt.Printf("parsing to int error %v", err)
		}
		sum := 0
		if string(sheet[4][i]) == "+" {
			sum = first + second + third + fourth
		} else if string(sheet[4][i]) == "*" {
			sum = first * second * third * fourth
		}
		answer = answer + sum
	}
	fmt.Println(answer)
}
