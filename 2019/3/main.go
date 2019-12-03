package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	w1 := wire(lines[0])
	w2 := wire(lines[1])
	ins := intersect(w1, w2)

	part1(w1, w2, ins)
	part2(w1, w2, ins)
}

func part1(w1, w2, ins []point) {
	min := math.MaxInt32
	for _, i := range ins {
		d := dist(i, point{0, 0})
		if d < min {
			min = d
		}
	}
	fmt.Println(min)
}

func part2(w1, w2, ins []point) {
	min := math.MaxInt32
	for _, i := range ins {
		v := stepsTo(i, w1) + stepsTo(i, w2)
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
}

type point struct{ x, y int }

func stepsTo(p point, w []point) int {
	for i, wp := range w {
		if eq(wp, p) {
			return i + 1
		}
	}
	panic("not here")
}

func dist(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func eq(p1, p2 point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func hash(p point) string {
	return fmt.Sprintf("x:%d,y:%d", p.x, p.y)
}

func intersect(w1, w2 []point) []point {
	ret := []point{}
	m := map[string]bool{}
	for _, p1 := range w1 {
		m[hash(p1)] = true
	}
	for _, p2 := range w2 {
		if has, ok := m[hash(p2)]; ok && has {
			ret = append(ret, p2)
		}
	}
	return ret
}

func wire(line string) []point {
	ret := []point{}
	parse := strings.Split(line, ",")

	cur := point{0, 0}
	for _, ins := range parse {
		dir := string(ins[0])
		dist, err := strconv.Atoi(ins[1:])
		if err != nil {
			panic(err)
		}

		var n point
		for i := 1; i <= dist; i++ {
			switch dir {
			case "R":
				n = point{cur.x + i, cur.y}
			case "L":
				n = point{cur.x - i, cur.y}
			case "U":
				n = point{cur.x, cur.y + i}
			case "D":
				n = point{cur.x, cur.y - i}
			default:
				panic("Unknown dir")
			}
			ret = append(ret, n)
		}
		cur = n
	}
	return ret
}
