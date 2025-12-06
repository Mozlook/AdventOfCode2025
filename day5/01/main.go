package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}
	defer rawInput.Close()

	var ranges []Range
	isRanges := true
	answer := 0
	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		lineInput := scanner.Text()
		if lineInput == "" {
			isRanges = false
			continue
		}

		if isRanges {
			rangesArray := strings.Split(lineInput, "-")
			start, err := strconv.Atoi(rangesArray[0])
			if err != nil {
				fmt.Println("conversion to int error")
			}
			end, err := strconv.Atoi(rangesArray[1])
			if err != nil {
				fmt.Println("conversion to int error")
			}
			ranges = append(ranges, Range{start: start, end: end})
		} else {
			ingredient, err := strconv.Atoi(lineInput)
			if err != nil {
				fmt.Println("conversion to int error")
			}
			for _, rangePair := range ranges {
				if ingredient >= rangePair.start && ingredient <= rangePair.end {
					answer++
					break
				}
			}
		}
	}
	fmt.Println(answer)
}
