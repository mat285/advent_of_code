package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var primes = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func main() {
	lines := read()
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		seen := map[string]bool{}
		words := strings.Split(line, " ")
		valid := true
		for _, word := range words {
			_, ok := seen[word]
			seen[word] = true
			valid = valid && !ok
		}
		if valid {
			sum++
		}
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	sum := 0
	for _, line := range lines {
		seen := map[int64]bool{}
		words := strings.Split(line, " ")
		valid := true
		for _, word := range words {
			h := hash(word)
			_, ok := seen[h]
			seen[h] = true
			valid = valid && !ok
		}
		if valid {
			sum++
		}
	}
	fmt.Println(sum)
}

func hash(s string) int64 {
	bs := []byte(s)
	prod := int64(1)

	for _, b := range bs {
		p := primes[int(b-byte('a'))]
		prod *= p
	}
	return prod
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
