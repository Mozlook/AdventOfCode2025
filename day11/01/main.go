package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DFS(start string, graph map[string][]string) int {
	if start == "out" {
		return 1
	}
	paths := 0
	neighbors := graph[start]
	for i := range neighbors {
		paths += DFS(neighbors[i], graph)
	}
	return paths
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
	fmt.Println(DFS("you", nbhList))
}
