package main

import (
	"fmt"
	"strconv"
)

var lower = 248345
var upper = 746315

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Println(count(lower, upper, ge))
}

func part2() {
	fmt.Println(count(lower, upper, eq))
}

func count(lower, upper int, cmp func(int, int) bool) int {
	count := 0
	for i := lower; i <= upper; i++ {
		if satisfies(i, cmp) {
			count++
		}
	}
	return count
}

func eq(x, y int) bool {
	return x == y
}

func ge(x, y int) bool {
	return x >= y
}

func satisfies(i int, cmp func(int, int) bool) bool {
	chars := strconv.Itoa(i)
	digits := make([]int, len(chars))
	for i, c := range chars {
		d, _ := strconv.Atoi(string(c))
		digits[i] = d
	}
	last := -1
	count := map[int]int{}
	for _, d := range digits {
		if d < last {
			return false
		}
		last = d
		count[d]++
	}
	for _, v := range count {
		if cmp(v, 2) {
			return true
		}
	}
	return false
}
