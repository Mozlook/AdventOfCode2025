package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkValidID(id string) bool {
	if len(id)%2 != 0 {
		return false
	}
	if id[0:len(id)/2] != id[len(id)/2:] {
		return false
	}
	return true
}

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v", err)
		return
	}
	defer rawInput.Close()
	answer := 0
	scanner := bufio.NewScanner(rawInput)
	scanner.Scan()
	rangeList := strings.SplitSeq(scanner.Text(), ",")
	for list := range rangeList {
		ranges := strings.Split(list, "-")
		bottomRange, err := strconv.Atoi(ranges[0])
		if err != nil {
			fmt.Println("error during converting to number")
		}

		topRange, err := strconv.Atoi(ranges[1])
		if err != nil {
			fmt.Println("error during converting to number")
		}

		for i := bottomRange; i < topRange; i++ {
			if checkValidID(strconv.Itoa(i)) {
				answer = answer + i
			}
		}
	}
	fmt.Println(answer)
}
