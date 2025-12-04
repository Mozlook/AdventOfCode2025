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

	const k = 12
	answer := 0

	scanner := bufio.NewScanner(rawInput)
	for scanner.Scan() {
		bank := scanner.Text()
		n := len(bank)

		chosen := make([]byte, 0, k)

		for i := range n {
			d := bank[i]
			remain := n - i

			for len(chosen) > 0 &&
				chosen[len(chosen)-1] < d &&
				len(chosen)-1+remain >= k {

				chosen = chosen[:len(chosen)-1]
			}

			if len(chosen) < k {
				chosen = append(chosen, d)
			}
		}

		value := 0
		for _, c := range chosen {
			value = value*10 + int(c-'0')
		}

		answer += value
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scan error: %v\n", err)
		return
	}

	fmt.Println(answer)
}
