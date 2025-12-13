package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type State struct {
	Node string
	Fft  bool
	Dac  bool
}

func DFS(node string, seenFft, seenDac bool, graph map[string][]string, memo map[State]int) int {
	if node == "fft" {
		seenFft = true
	}
	if node == "dac" {
		seenDac = true
	}

	if node == "out" {
		if seenFft && seenDac {
			return 1
		}
		return 0
	}

	st := State{Node: node, Fft: seenFft, Dac: seenDac}
	if val, ok := memo[st]; ok {
		return val
	}

	paths := 0
	for _, nb := range graph[node] {
		paths += DFS(nb, seenFft, seenDac, graph, memo)
	}

	memo[st] = paths
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

	memo := make(map[State]int)
	fmt.Println(DFS("svr", false, false, nbhList, memo))
}
