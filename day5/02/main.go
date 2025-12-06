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

type Range struct {
	start int
	end   int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}
	defer rawInput.Close()

	var ranges []Range
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		lineInput := strings.TrimSpace(scanner.Text())
		if lineInput == "" {
			break
		}

		rangesArray := strings.Split(lineInput, "-")
		start, err := strconv.Atoi(strings.TrimSpace(rangesArray[0]))
		if err != nil {
			fmt.Println("conversion to int error")
		}
		end, err := strconv.Atoi(strings.TrimSpace(rangesArray[1]))
		if err != nil {
			fmt.Println("conversion to int error")
		}
		ranges = append(ranges, Range{start: start, end: end})
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scan error: %v\n", err)
		return
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.start != b.start {
			return cmp.Compare(a.start, b.start)
		}
		return cmp.Compare(a.end, b.end)
	})

	if len(ranges) == 0 {
		fmt.Println(0)
		return
	}

	answer := 0

	cur_start := ranges[0].start
	cur_end := ranges[0].end

	for i := 1; i < len(ranges); i++ {
		rangePair := ranges[i]
		if rangePair.start <= cur_end+1 {
			cur_end = max(cur_end, rangePair.end)
		} else {
			answer += (cur_end - cur_start + 1)
			cur_start = rangePair.start
			cur_end = rangePair.end
		}
	}

	answer += (cur_end - cur_start + 1)

	fmt.Println(answer)
}
