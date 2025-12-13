package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DFS(start string, graph map[string][]string, visited map[string]bool, hasSpecial bool) int {
	if start == "fft" || start == "dac" {
		hasSpecial = true
	}

	if visited[start] {
		return 0
	}

	if start == "out" {
		if hasSpecial {
			return 1
		}
		return 0
	}

	visited[start] = true
	defer func() { visited[start] = false }()

	count := 0
	for _, nb := range graph[start] {
		count += DFS(nb, graph, visited, hasSpecial)
	}
	return count
}

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}

	defer rawInput.Close()
	nbhList := make(map[string][]string)
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		key := fields[0][:3]
		nbhList[key] = fields[1:]
	}
	fmt.Println(DFS("svr", nbhList, make(map[string]bool, 0), false))
}
