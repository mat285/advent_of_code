package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const size = 1100

var testLines = strings.Split(`#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2`, "\n")

func main() {
	lines := read()
	claims := parseClaims(lines)
	part1(claims)
	part2(claims)
}

type entry struct {
	claimed []int
}

type point struct{ x, y int }

type claim struct {
	id     int
	corner point
	height int
	width  int
}

func part1(claims []*claim) {
	grid := buildGrid(claims)

	covered := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != nil {
				if len(grid[i][j].claimed) > 1 {
					covered++
				}
			}
		}
	}

	fmt.Println(covered)
}

func part2(claims []*claim) {
	grid := buildGrid(claims)

	m := map[int]bool{}

	for _, c := range claims {
		m[c.id] = false
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != nil {
				if len(grid[i][j].claimed) > 1 {
					for _, id := range grid[i][j].claimed {
						delete(m, id)
					}
				}
			}
		}
	}

	if len(m) > 1 {
		fmt.Println("error")
		return
	}
	for i := range m {
		fmt.Println(i)
	}
}

func buildGrid(claims []*claim) [][]*entry {
	grid := make([][]*entry, size)

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]*entry, size)
	}

	for _, c := range claims {
		for x := c.corner.x; x < c.corner.x+c.width; x++ {
			for y := c.corner.y; y < c.corner.y+c.height; y++ {
				if grid[y][x] == nil {
					grid[y][x] = &entry{}
				}
				grid[y][x].claimed = append(grid[y][x].claimed, c.id)
			}
		}
	}

	return grid
}

func parseClaims(lines []string) []*claim {
	claims := []*claim{}
	for _, line := range lines {
		c := &claim{}
		_, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.id, &c.corner.x, &c.corner.y, &c.width, &c.height)
		if err != nil {
			panic(err)
		}
		claims = append(claims, c)
	}
	return claims
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
