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

	lines := strings.Split(strings.TrimSpace(string(data)), ",")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	is := toInt(lines)
	is[1] = 12
	is[2] = 2
	fmt.Println(run(is)[0])
}

func part2(lines []string) {
	src := toInt(lines)
	for x := 0; x <= 99; x++ {
		for y := 0; y <= 99; y++ {
			is := make([]int, len(src))
			copy(is, src)
			is[1] = x
			is[2] = y
			o := run(is)[0]
			if o == 19690720 {
				fmt.Println((100 * x) + y)
				return
			}
		}
	}

}

func toInt(lines []string) []int {
	ret := []int{}

	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ret = append(ret, i)
	}

	return ret
}

func run(is []int) []int {
	counter := 0

	for {
		op := is[counter]
		if op == 99 {
			return is
		}
		x := is[counter+1]
		y := is[counter+2]
		o := is[counter+3]

		if op == 1 {
			is[o] = is[x] + is[y]
		} else if op == 2 {
			is[o] = is[x] * is[y]
		} else {
			panic("unknown")
		}
		counter += 4
	}
}
