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
	grid := parseTiles(lines)
	for i := 0; i < 10; i++ {
		grid = next(grid)
	}
	fmt.Println(resourceValue(grid))
}

func part2(lines []string) {
	grid := parseTiles(lines)
	seen := map[string]int{}
	minutes := 1000000000
	cycleSize := 1
	cycleStart := 0
	for i := 0; i < minutes; i++ {
		grid = next(grid)
		h := hash(grid)
		v := resourceValue(grid)
		if v == 214375 {
			fmt.Print("hi ")
		}
		fmt.Println(i, i%cycleSize, cycleSize, v)
		if c, ok := seen[h]; !ok {
			seen[h] = i
		} else {
			if cycleSize == 1 {
				cycleSize = i - c
				cycleStart = i
			}
		}
		if i == 454+cycleSize {
			break
		}
	}
	fmt.Println(cycleSize, cycleStart)

	vals := make([]int, cycleSize)
	for i := cycleStart; i < cycleSize+cycleStart; i++ {
		idx := i % cycleSize
		vals[idx] = resourceValue(grid)
		grid = next(grid)
	}

	fmt.Println(vals)

	idx := (minutes - 1) % cycleSize
	fmt.Println(vals[idx])
}

type point struct{ x, y int }

type tile string

const (
	open   tile = "."
	lumber tile = "#"
	trees  tile = "|"
)

func hash(grid [][]tile) string {
	bs := make([]byte, 0, len(grid)*len(grid[0]))
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			b := []byte(grid[y][x])
			bs = append(bs, b[0])
		}
	}
	return string(bs)
}

func next(grid [][]tile) [][]tile {
	ng := make([][]tile, len(grid))
	for y := 0; y < len(grid); y++ {
		ng[y] = make([]tile, len(grid[y]))
		for x := 0; x < len(grid[y]); x++ {
			adj := adjacent(point{x, y}, grid)
			ng[y][x] = evolve(grid[y][x], adj)
		}
	}
	return ng
}

func evolve(t tile, adj []tile) tile {
	switch t {
	case open:
		if count(trees, adj) >= 3 {
			return trees
		}
		return open
	case lumber:
		if count(lumber, adj) >= 1 && count(trees, adj) >= 1 {
			return lumber
		}
		return open
	case trees:
		if count(lumber, adj) >= 3 {
			return lumber
		}
		return trees
	default:
		panic(t)
	}
}

func resourceValue(grid [][]tile) int {
	countLumber := 0
	countTrees := 0
	for _, row := range grid {
		countLumber += count(lumber, row)
		countTrees += count(trees, row)
	}
	return countLumber * countTrees
}

func count(t tile, ts []tile) int {
	c := 0
	for _, o := range ts {
		if o == t {
			c++
		}
	}
	return c
}

func adjacent(p point, grid [][]tile) []tile {
	adj := []tile{}
	pts := []point{
		{p.x, p.y - 1},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x - 1, p.y - 1},
		{p.x + 1, p.y},
		{p.x + 1, p.y - 1},
		{p.x + 1, p.y + 1},
	}

	for _, o := range pts {
		if o.x >= 0 && o.y >= 0 && o.y < len(grid) && o.x < len(grid[o.y]) {
			adj = append(adj, grid[o.y][o.x])
		}
	}
	return adj
}

func parseTiles(lines []string) [][]tile {
	grid := make([][]tile, len(lines))
	for i, line := range lines {
		for _, c := range line {
			grid[i] = append(grid[i], tile(c))
		}
	}
	return grid
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split((string(data[:len(data)-1])), "\n")
}
