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

		a1, a2 := min_2(l, w, h)
		sum += a1 + a1 + a2 + a2
		sum += l * w * h
	}
	fmt.Println(sum)
	//return sum
}

func min_2(a int, b int, c int) (int, int) {
	if a <= b {
		if b <= c {
			return a , b
		}else{
			return a, c
		}
	}else{
		if a <= c{
			return a, b
		}else{
			return b, c
		}
	}

}
