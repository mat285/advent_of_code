package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func main() {

	part1()
}

func part1() {
	lines := read()
	lights := parseLights(lines)

	found := search(lights)

	lights = parseLights(lines)

	for i := 0; i < found-3; i++ {
		tick(lights)
	}
	for i := 0; i < 6; i++ {
		plotL(lights, i)
		tick(lights)
	}
}

func hackPart1() {
	lines := read()
	lights := parseLights(lines)
	// manual binary search to find the value #yolo
	for i := 0; i < 10054; i++ {
		tick(lights)
	}
	for i := 0; i < 1; i++ {
		plotL(lights, i)
		tick(lights)
	}
}

func search(ls []*light) int {
	i := 0
	prev := spread(ls)
	decreasing := true
	for decreasing {
		tick(ls)
		curr := spread(ls)
		if curr.x+curr.y > prev.x+prev.y {
			decreasing = false
		}
		prev = curr
		i++
	}
	return i
}

type point struct {
	x, y int
}

type light struct {
	position point
	velocity point
}

type lights []*light

func (l lights) Len() int {
	return len(l)
}

func (l lights) XY(i int) (float64, float64) {
	return float64(l[i].position.x), -float64(l[i].position.y)
}

func spread(ls []*light) point {
	max, min := bounds(ls)
	return point{max.x - min.x, max.y - min.y}
}

func tick(lights []*light) {
	for _, l := range lights {
		l.position.x += l.velocity.x
		l.position.y += l.velocity.y
	}
}

func plotL(ls []*light, t int) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	scatter, err := plotter.NewScatter(lights(ls))
	if err != nil {
		panic(err)
	}
	p.Add(scatter)
	// plot is reversed so go look in a mirror
	err = p.Save(800, 100, fmt.Sprintf("%d.png", t))
	if err != nil {
		panic(err)
	}
}

func normalize(lights []*light, mmx, mmy int) {
	for _, l := range lights {
		l.position.x -= mmx
		l.position.y -= mmy
	}
}

func bounds(lights []*light) (point, point) {
	var max, min point
	max.x = math.MinInt32
	max.y = math.MinInt32

	min.x = math.MaxInt32
	min.y = math.MaxInt32

	for _, l := range lights {
		if l.position.x > max.x {
			max.x = l.position.x
		}
		if l.position.y > max.y {
			max.y = l.position.y
		}
		if l.position.x < min.x {
			min.x = l.position.x
		}
		if l.position.y < min.y {
			min.y = l.position.y
		}
	}
	return max, min
}

func parseLights(lines []string) []*light {
	lights := []*light{}
	for _, line := range lines {
		line = strings.Replace(line, "< ", "<", -1)
		line = strings.Replace(line, "  ", " ", -1)
		l := &light{}
		_, err := fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &l.position.x, &l.position.y, &l.velocity.x, &l.velocity.y)
		if err != nil {
			panic(err)
		}
		lights = append(lights, l)
	}
	return lights
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
