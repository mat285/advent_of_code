package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	twoCount := 0
	threeCount := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		two, three := count(line)
		twoCount += two
		threeCount += three
	}
	fmt.Println(twoCount * threeCount)
}

func part2(lines []string) {

	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			res, single := oneDiff(lines[i], lines[j])
			if single {
				fmt.Println(lines[i])
				fmt.Println(lines[j])
				fmt.Println(string(res))
				return
			}
		}
	}

	fmt.Println("error")
}

func count(str string) (int, int) {
	m := map[rune]int{}
	for _, r := range str {
		if val, ok := m[r]; !ok {
			m[r] = 1
		} else {
			m[r] = val + 1
		}
	}
	tw := 0
	th := 0
	for _, c := range m {
		if c == 2 {
			tw = 1
		}
		if c == 3 {
			th = 1
		}
	}
	return tw, th
}

func oneDiff(s1, s2 string) ([]byte, bool) {
	b1 := []byte(s1)
	b2 := []byte(s2)
	res := []byte{}
	if len(b1) != len(b2) {
		return nil, false
	}
	diff := false
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			if diff {
				return nil, false
			}
			diff = true
		} else {
			res = append(res, b1[i])
		}
	}
	return res, diff
}
