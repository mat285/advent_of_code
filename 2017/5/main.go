package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	lines := read()
	jumps := translate(lines)
	part1(jumps)
	jumps = translate(lines)
	part2(jumps)
}

func part1(jumps []int) {
	i := 0
	count := 0
	for i >= 0 && i < len(jumps) {
		j := jumps[i]
		jumps[i]++
		i += j
		count++
	}
	fmt.Println(count)
}

func part2(jumps []int) {
	i := 0
	count := 0
	for i >= 0 && i < len(jumps) {
		j := jumps[i]
		if j < 3 {
			jumps[i]++
		} else {
			jumps[i]--
		}
		i += j
		count++
	}
	fmt.Println(count)
}

func translate(lines []string) []int {
	jumps := []int{}
	for _, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		jumps = append(jumps, v)
	}
	return jumps
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
