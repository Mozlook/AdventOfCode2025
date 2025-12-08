package main

import (
	"bufio"
	"fmt"
	"os"
)

func tachyonBeam(x int, y int, diagram []string) int {
	sum := 0
	for i := x; i < len(diagram)-1; i++ {
		switch diagram[i+1][y] {
		case '^':
			sum++
			leftSum := 0
			rightSum := 0
			if y-1 >= 0 {
				leftSum = tachyonBeam(i+1, y-1, diagram)
			}
			if y+1 < len(diagram[i+1]) {
				rightSum = tachyonBeam(i+1, y+1, diagram)
			}
			return leftSum + rightSum + sum
		case '|':
			return 0

		case '.':
			row := []rune(diagram[i+1])
			row[y] = '|'
			diagram[i+1] = string(row)

		default:
			return sum
		}
	}
	return sum
}

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}
	defer rawInput.Close()
	var diagram []string
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		line := scanner.Text()
		diagram = append(diagram, line)
	}
	start := 0
	for i := range diagram[0] {
		if diagram[0][i] == 'S' {
			start = i
			break
		}
	}
	fmt.Print(tachyonBeam(0, start, diagram))
}
