package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2015/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		lineStr := scanner.Text()

		s := strings.Split(lineStr, "x")
		//fmt.Println( reflect.TypeOf(s))
		l, _ := strconv.Atoi(s[0])
		w, _ := strconv.Atoi(s[1])
		h, _ := strconv.Atoi(s[2])

		s1 :=  l * w
		s2 :=  l * h
		s3 :=  w * h
		mini := min(s1, s2, s3)
		fmt.Println(l, w, h)
		fmt.Println(s1, s2, s3, mini)
		current_sum := 2 * s1 + 2 * s2 + 2 * s3 + mini
		sum += current_sum
		fmt.Println("sum", sum)
	}
	fmt.Println(sum)
	//return sum
}

func min(a int, b int, c int) int {
	if a <= b {
		if a <= c {
			return a
		} else {
			return c
		}
	} else {
		if b <= c {
			return b
		} else {
			return c
		}
	}
}
