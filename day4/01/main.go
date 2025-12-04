package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		diagram = append(diagram, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scan error: %v\n", err)
		return
	}

	h := len(diagram)
	if h == 0 {
		fmt.Println(0)
		return
	}
	w := len(diagram[0])

	dirs := [8][2]int{
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{0, 1},
		{0, -1},
		{1, 0},
		{1, -1},
		{1, 1},
	}

	answer := 0

	for i := range h {
		for j := range w {
			if diagram[i][j] != '@' {
				continue
			}

			count := 0
			for _, d := range dirs {
				ni := i + d[0]
				nj := j + d[1]

				if ni < 0 || ni >= h || nj < 0 || nj >= w {
					continue
				}
				if diagram[ni][nj] == '@' {
					count++
					if count > 3 {
						break
					}
				}
			}

			if count <= 3 {
				answer++
			}
		}
	}

	fmt.Println(answer)
}
