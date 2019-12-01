package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	lines := read()
	part1(lines)
}

func part1(lines []string) {
	pts := parsePoints(lines)
	groups := map[point][]*point{}

	for _, p := range pts {
		groups[*p] = group(p, pts)
	}

	cons := merge(groups)
	fmt.Println(len(cons))
}

type point struct{ x, y, z, t int }

type constellation map[point]bool

func merge(groups map[point][]*point) []constellation {
	ret := []constellation{}

	for len(groups) > 0 {
		p := selectKey(groups)
		c := constellation{}
		include := dedup(getAll(p, groups, map[point]bool{}))
		for _, o := range include {
			c[*o] = true
			delete(groups, *o)
		}
		ret = append(ret, c)
	}

	return ret
}

func getAll(p *point, groups map[point][]*point, seen map[point]bool) []*point {
	ret := []*point{p}
	seen[*p] = true
	g := groups[*p]
	for _, o := range g {
		all := []*point{}
		if !seen[*o] {
			all = getAll(o, groups, seen)
		}
		ret = append(ret, all...)
	}
	return ret
}

func dedup(pts []*point) []*point {
	m := map[point]bool{}
	ret := []*point{}
	for _, p := range pts {
		if _, ok := m[*p]; !ok {
			ret = append(ret, p)
			m[*p] = true
		}
	}
	return ret
}

func selectKey(m map[point][]*point) *point {
	for k := range m {
		return &k
	}
	return nil
}

func group(p *point, pts []*point) []*point {
	g := []*point{}
	for _, o := range pts {
		if distance(p, o) <= 3 {
			g = append(g, o)
		}
	}
	return g
}

func distance(p1, p2 *point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z) + abs(p1.t-p2.t)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func parsePoints(lines []string) []*point {
	pts := []*point{}
	for _, line := range lines {
		p := &point{}
		_, err := fmt.Sscanf(line, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.t)
		if err != nil {
			panic(err)
		}
		pts = append(pts, p)
	}
	return pts
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split((string(data[:len(data)-1])), "\n")
}
