package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	s := read()
	part1(s)
	part2(s)
}

func part1(s string) {
	fmt.Println(process(s, 1))
}

func part2(s string) {
	fmt.Println(process(s, len(s)/2))
}

func process(s string, step int) int {
	sum := 0
	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		next := (i + step) % len(bs)
		if bs[i] == bs[next] {
			v, err := strconv.Atoi(string(bs[i]))
			if err != nil {
				panic(err)
			}
			sum += v
		}
	}
	return sum
}

func read() string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}
