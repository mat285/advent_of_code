package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	lines := read()
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	bots := parseNanobots(lines)
	s := strongest(bots)
	in := inRangeBot(s, bots)
	fmt.Println(len(in))
}

func part2(lines []string) {
	bots := parseNanobots(lines)
	pos := point{0, 0, 0}
	curr := inRangePoint(&pos, bots)
	scale := 1
	checks := 0
	pts := []point{
		point{0, 0, 1},
		point{0, 1, 0},
		point{1, 0, 0},
		point{0, 0, -1},
		point{0, -1, 0},
		point{-1, 0, 0},
	}
	dir := 1
	i := 0
	for i < 100000 {
		next := pos.add(pts[dir].scale(scale))
		vn := inRangePoint(&next, bots)
		if vn < curr {
			if scale == 1 {
				dir = (dir + 1) % len(pts)
				checks++
				if checks > len(pts) {
					break
				}
			} else {
				scale /= 2
				checks = 0
			}
		} else {
			pos = next
			curr = vn
			scale *= 2
			checks = 0
		}
		i++
	}
	fmt.Println(distance(pos, point{0, 0, 0}))

}

type point struct{ x, y, z int }
type nanobot struct {
	point
	r int
}

func (p point) scale(f int) point {
	p.x *= f
	p.y *= f
	p.z *= f
	return p
}

func (p point) add(o point) point {
	p.x += o.x
	p.y += o.y
	p.z += o.z
	return p
}

func inRangePoint(p *point, bots []*nanobot) int {
	count := 0
	for _, b := range bots {
		if distance(*p, b.point) <= b.r {
			count++
		}
	}
	return count
}

func inRangeBot(b *nanobot, bots []*nanobot) []*nanobot {
	ret := []*nanobot{}
	for _, n := range bots {
		if distance(b.point, n.point) <= b.r {
			ret = append(ret, n)
		}
	}
	return ret
}

func strongest(bots []*nanobot) *nanobot {
	var max *nanobot
	for _, b := range bots {
		if max == nil || b.r > max.r {
			max = b
		}
	}
	return max
}

func distance(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func parseNanobots(lines []string) []*nanobot {
	bots := []*nanobot{}
	for _, line := range lines {
		n := &nanobot{}
		_, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &n.x, &n.y, &n.z, &n.r)
		if err != nil {
			panic(err)
		}
		bots = append(bots, n)
	}
	return bots
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split((string(data[:len(data)-1])), "\n")
}
