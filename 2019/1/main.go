package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	fuel := 0

	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		fuel += required(i)
	}
	fmt.Println(fuel)
}

func part2(lines []string) {
	fuel := 0
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		req := required(i)
		for req != 0 {
			fuel += req
			req = required(req)
		}
	}
	fmt.Println(fuel)
}

func required(i int) int {
	i = i / 3
	i = i - 2
	if i < 0 {
		return 0
	}
	return i
}
