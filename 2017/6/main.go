package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}

func main() {
	part1(duplicate(input))
}

func part1(banks []int) {
	m := map[string]int{}
	count := 0
	for !seen(m, banks) {
		m[stringify(banks, 0)] = count
		redestribute(banks, maxIdx(banks))
		count++
	}
	fmt.Println(count)
	fmt.Println(count - m[stringify(banks, 0)])
}

func redestribute(banks []int, idx int) {
	v := banks[idx]
	banks[idx] = 0
	for i := 0; i < v; i++ {
		banks[(idx+i+1)%len(banks)]++
	}
}

func maxIdx(banks []int) int {
	idx := 0
	for i := 0; i < len(banks); i++ {
		if banks[idx] < banks[i] {
			idx = i
		}
	}
	return idx
}

func seen(m map[string]int, arr []int) bool {
	for i := 0; i < 1; i++ {
		_, ok := m[stringify(arr, i)]
		if ok {
			return true
		}
	}
	return false
}

func stringify(banks []int, start int) string {
	s := []string{}
	for i := 0; i < len(banks); i++ {
		s = append(s, strconv.Itoa(banks[(start+i)%len(banks)]))
	}
	return strings.Join(s, " ")
}

func duplicate(arr []int) []int {
	ret := make([]int, len(arr))
	for i, v := range arr {
		ret[i] = v
	}
	return ret
}
