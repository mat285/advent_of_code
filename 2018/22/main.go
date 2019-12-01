package main

import (
	"fmt"
	"math"
)

var inputDepth = 3198
var inputTarget = &point{12, 757}

var mod = 20183
var gidxX = 16807
var gidxY = 48271

var changeCost = 7
var stepCost = 1

func main() {
	part1()
}

func part1() {
	depth := inputDepth   //510
	target := inputTarget //&point{10, 10}

	ero := erosionMatrix(depth, target)
	gm := geoMap(ero)
	printGrid(gm, 11, 12)
	fmt.Println(riskLevel(gm, target))
}

type point struct{ x, y int }

type regionType int

const (
	rocky  regionType = 0
	wet    regionType = 1
	narrow regionType = 2
)

type equipment int

const (
	torch    equipment = 0
	climbing equipment = 1
	neither  equipment = 2
)

var equipments = []equipment{torch, climbing, neither}

type triple struct {
	point
	e equipment
}

func printGrid(grid [][]regionType, mx, my int) {
	for y := 0; y < my; y++ {
		for x := 0; x < mx; x++ {
			switch grid[y][x] {
			case rocky:
				fmt.Print(".")
			case wet:
				fmt.Print("=")
			case narrow:
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

func search(grid [][]regionType, target *point) int {
	m := make([][][]int, len(grid))
	for y := 0; y < len(m); y++ {
		m[y] = make([][]int, len(grid[y]))
		for x := 0; x < len(m[y]); x++ {
			m[y][x] = make([]int, len(equipments))
			for _, e := range equipments {
				m[y][x][e] = math.MaxInt32
			}
		}
	}

	visited := map[triple]bool{}
	curr := triple{point{0, 0}, torch}
	m[curr.y][curr.x][curr.e] = 0

	for {
		adj := adjacent(curr.point)
		neighbors := []triple{}
		for _, a := range adj {
			for _, e := range equipments {
				neighbors = append(neighbors, triple{point: a, e: e})
			}
		}
		for _, n := range neighbors {
			if visited[n] {
				continue
			}
			if validFor(grid[n.y][n.x], curr.e) {

			}
		}
	}

	m[target.y][target.x][torch] = 0             // torch done
	m[target.y][target.x][climbing] = changeCost // climbing must change

	freeList := map[point]bool{}
	seen := map[point]bool{}

	seen[*target] = true

	for _, p := range adjacent(*target) {
		if inBounds(p, grid) {
			freeList[p] = true
		}
	}

	// for len(freeList) > 0 {
	// 	p := selectKey(freeList)

	// }

	return m[0][0][torch]
}

func selectKey(m map[point]bool) point {
	for k := range m {
		return k
	}
	return point{-1, -1}
}

func adjacent(p point) []point {
	return []point{
		point{p.x + 1, p.y},
		point{p.x - 1, p.y},
		point{p.x, p.y + 1},
		point{p.x, p.y - 1},
	}
}

func inBounds(p point, grid [][]regionType) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[p.y])
}

func validFor(r regionType, e equipment) bool {
	return contains(e, allowed(r))
}

func allowed(r regionType) []equipment {
	switch r {
	case rocky:
		return []equipment{torch, climbing}
	case wet:
		return []equipment{climbing, neither}
	case narrow:
		return []equipment{torch, neither}
	}
	return []equipment{}
}

func contains(e equipment, es []equipment) bool {
	for _, o := range es {
		if e == o {
			return true
		}
	}
	return false
}

func riskLevel(grid [][]regionType, target *point) int {
	sum := 0
	for y := 0; y <= target.y; y++ {
		for x := 0; x <= target.x; x++ {
			sum += int(grid[y][x])
		}
	}
	return sum
}

func geoMap(grid [][]int) [][]regionType {
	geo := make([][]regionType, len(grid))
	for y := 0; y < len(geo); y++ {
		geo[y] = make([]regionType, len(grid[y]))
		for x := 0; x < len(geo[y]); x++ {
			geo[y][x] = regionType(grid[y][x] % 3)
		}
	}
	return geo
}

func erosionMatrix(depth int, target *point) [][]int {
	grid := make([][]int, depth+10)
	for y := 0; y <= depth; y++ {
		grid[y] = make([]int, target.x+20)
		if y == 0 {
			for x := 0; x <= target.x+20; x++ {
				grid[y][x] = (x*gidxX + depth) % mod
			}
		}
		grid[y][0] = (y*gidxY + depth) % mod
	}

	for y := 1; y < len(grid); y++ {
		for x := 1; x < len(grid[y]); x++ {
			if x == target.x && y == target.y {
				grid[y][x] = depth % mod
			} else {
				grid[y][x] = (grid[y][x-1]*grid[y-1][x] + depth) % mod
			}
		}
	}
	return grid
}
