package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
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
	_, cords := buildGrid(lines)
	max := math.MinInt32
	maxP := point{-1, -1}
	for _, p := range cords {
		los := lineOfSight(p, cords)
		if len(los) > max {
			max = len(los)
			maxP = p
		}
	}
	fmt.Println(max, maxP)
}

func part2(lines []string) {
	station := point{11, 19}
	_, cords := buildGrid(lines)

	cm := map[string]point{}

	for _, c := range cords {
		cm[c.hash()] = c
	}
	delete(cm, station.hash())

	m := map[int]point{}
	i := 1
	for len(cm) > 0 {
		remaining := toSlice(cm)
		los := lineOfSight(station, remaining)
		vectors := vectorize(station, los)

		for _, v := range vectors {
			m[i] = v.e
			delete(cm, v.e.hash())
			i++
		}
	}
	fmt.Println(m[200])
}

func toSlice(m map[string]point) []point {
	ret := []point{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func vectorize(p point, ps []point) []vector {
	vectors := []vector{}
	origin := vector{s: p, e: point{p.x, p.y - 1}}
	for _, o := range ps {
		v := vector{s: p, e: o}
		vectors = append(vectors, v)
	}
	sort.Slice(vectors, func(i, j int) bool {
		return origin.angle(vectors[i]) < origin.angle(vectors[j])
	})
	return vectors
}

func (v vector) angle(u vector) float64 {
	x1 := float64(v.e.x - v.s.x)
	x2 := float64(u.e.x - u.s.x)
	y1 := float64(v.e.y - v.s.y)
	y2 := float64(u.e.y - u.s.y)
	theta := math.Atan2(x1*y2-y1*x2, x1*x2+y1*y2) // [-pi, pi]

	if theta >= 0 {
		return theta
	}
	return 2*math.Pi + theta
}

func lineOfSight(p point, ps []point) []point {
	ds := []datum{}
	for _, other := range ps {
		if !eq(p, other) {
			ds = append(ds, datum{
				p: other,
				m: slope(p, other),
				d: dist(p, other),
			})
		}
	}

	m := map[string]datum{}

	for _, d := range ds {
		if c, ok := m[d.m.hash()]; ok {
			if c.d > d.d {
				m[d.m.hash()] = d
			}
		} else {
			m[d.m.hash()] = d
		}
	}

	visible := []point{}
	for _, d := range m {
		visible = append(visible, d.p)
	}
	return visible
}

func (p point) hash() string {
	return fmt.Sprintf("x:%d,y:%d", p.x, p.y)
}

func (p point) String() string {
	return p.hash()
}

func slope(p, q point) point {
	rise := p.y - q.y
	run := p.x - q.x
	g := gcd(abs(rise), abs(run))
	return point{run / g, rise / g}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func eq(p, q point) bool {
	return dist(p, q) == 0
}

func dist(p, q point) int {
	return abs(p.x-q.x) + abs(p.y-q.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type datum struct {
	p point
	m point
	d int
}

type point struct{ x, y int }

type vector struct{ s, e point }

func buildGrid(lines []string) ([][]int, []point) {
	cords := []point{}
	grid := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		grid[i] = make([]int, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '#' {
				grid[i][j] = 1
				cords = append(cords, point{j, i})
			}
		}
	}
	return grid, cords
}
