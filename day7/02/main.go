package main

import (
	"bufio"
	"fmt"
	"os"
)

func tachyonBeam(x, y int, diagram []string, memo [][]int) int {
	if x >= len(diagram)-1 {
		return 1
	}
	if y < 0 || y >= len(diagram[x]) {
		return 1
	}

	if memo[x][y] != -1 {
		return memo[x][y]
	}

	for i := x; i < len(diagram)-1; i++ {
		if y < 0 || y >= len(diagram[i+1]) {
			memo[x][y] = 1
			return 1
		}

		if diagram[i+1][y] == '^' {
			leftSum := 0
			rightSum := 0

			if y-1 >= 0 {
				leftSum = tachyonBeam(i+1, y-1, diagram, memo)
			}
			if y+1 < len(diagram[i+1]) {
				rightSum = tachyonBeam(i+1, y+1, diagram, memo)
			}

			memo[x][y] = leftSum + rightSum
			return memo[x][y]
		}
	}

	memo[x][y] = 1
	return 1
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

	rows := len(diagram)
	col := len(diagram[0])
	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, col)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	fmt.Print(tachyonBeam(0, start, diagram, memo))
}
