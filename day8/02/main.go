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

type JunctionBox struct {
	x int
	y int
	z int
}

type PairDistance struct {
	firstBox  int
	secondBox int
	distance  int64
}

func calcDistSquared(first, second JunctionBox) int64 {
	dx := int64(first.x - second.x)
	dy := int64(first.y - second.y)
	dz := int64(first.z - second.z)
	return dx*dx + dy*dy + dz*dz
}

func Find(i int, parent []int) int {
	if parent[i] == i {
		return i
	}
	parent[i] = Find(parent[i], parent)
	return parent[i]
}

func Union(i, j int, parent, size []int) int {
	ri := Find(i, parent)
	rj := Find(j, parent)

	if ri == rj {
		return 0
	}
	if size[ri] < size[rj] {
		ri, rj = rj, ri
	}
	parent[rj] = ri
	size[ri] += size[rj]
	return -1
}

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}

	defer rawInput.Close()
	var boxes []JunctionBox
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
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("parsing error")
		}
		boxes = append(boxes, JunctionBox{x: x, y: y, z: z})
	}

	var pairDist []PairDistance
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			pairDist = append(pairDist, PairDistance{i, j, calcDistSquared(boxes[i], boxes[j])})
		}
	}

	slices.SortFunc(pairDist, func(a, b PairDistance) int {
		return cmp.Compare(a.distance, b.distance)
	})
	n := len(boxes)
	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i
		size[i] = 1
	}
	components := n
	lastPair := PairDistance{}
	for k := range pairDist {
		p := pairDist[k]
		components += Union(p.firstBox, p.secondBox, parent, size)
		if components == 1 {
			lastPair = pairDist[k]
			break
		}
	}

	answer := boxes[lastPair.firstBox].x * boxes[lastPair.secondBox].x
	fmt.Println(answer)
}
