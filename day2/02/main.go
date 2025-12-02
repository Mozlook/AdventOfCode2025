package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkValidID(id string) bool {
	n := len(id)
	if n == 0 {
		return false
	}

	for size := 1; size <= n/2; size++ {
		if n%size != 0 {
			continue
		}

		pattern := id[0:size]
		ok := true

		for j := size; j < n; j += size {
			if id[j:j+size] != pattern {
				ok = false
				break
			}
		}

		if ok {
			return true
		}
	}

	return false
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
