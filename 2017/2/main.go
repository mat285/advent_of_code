package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/blend/go-sdk/util"
)

func main() {
	lines := read()
	sheet := toInts(lines)
	part1(sheet)
	part2(sheet)
}

func part1(sheet [][]int) {
	sum := 0
	for _, row := range sheet {
		min, max := util.Math.MinAndMaxOfInt(row...)
		sum += max - min
	}
	fmt.Println(sum)
}

func part2(sheet [][]int) {
	sum := 0
	for _, row := range sheet {
		for i := 0; i < len(row); i++ {
			for j := 0; j < len(row); j++ {
				if j != i && row[i]%row[j] == 0 {
					sum += row[i] / row[j]
				}
			}
		}
	}
	fmt.Println(sum)
}

func toInts(lines []string) [][]int {
	sheet := make([][]int, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, "\t")
		sheet[i] = make([]int, len(parts))
		for j, p := range parts {
			v, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}
			sheet[i][j] = v
		}
	}
	return sheet
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
