package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type TileLocation struct {
	x int
	y int
}

type PairArea struct {
	firstTile  int
	secondTile int
	area       int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calcArea(first, second TileLocation) int {
	dx := Abs(int(first.x-second.x) + 1)
	dy := Abs(int(first.y-second.y) + 1)
	return dx * dy
}

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}

	defer rawInput.Close()
	var tiles []TileLocation
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("parsing error")
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("parsing error")
		}

		tiles = append(tiles, TileLocation{x: x, y: y})
	}

	var pairArea []PairArea
	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			pairArea = append(pairArea, PairArea{i, j, calcArea(tiles[i], tiles[j])})
		}
	}

	slices.SortFunc(pairArea, func(a, b PairArea) int {
		return -cmp.Compare(a.area, b.area)
	})
	fmt.Println(pairArea[0])
}
