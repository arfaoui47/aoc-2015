package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

func main() {
	file, err := os.Open("2015/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var arr [20000][20000]int
	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		fmt.Println(lineStr)
		lineStr = ">"
		i := 10000
		j := 10000
		arr[i][j] = 1
		sum = 1
		for i, rune := range lineStr {
			fmt.Println(i)
			if arr[i][j] > 0 {
				continue
			}
			if string(rune) == ">" {
				j += 1
				arr[i][j] = 1
				sum += 1
			} else if string(rune) == "^" {
				i += 1
				arr[i][j] = 1
				sum += 1
			} else if string(rune) == "<" {
				j -= 1
				arr[i][j] = 1
				sum += 1
			} else {
				i -= 1
				arr[i][j] = 1
				sum += 1
			}
		}
	}
	fmt.Println(sum)
}
