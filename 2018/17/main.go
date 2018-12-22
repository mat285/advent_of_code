package main

import (
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	tileValueClay          = 0
	tileValueSand          = 1
	tileValueWaterStagnant = 2
	tileValueWaterFalling  = 3
	tileValueSource        = 4
)

var source = point{500, 0}

func main() {
	lines := read()
	part1(lines)
}

func part1(lines []string) {
	pts := parsePoints(lines)
	grid := makeGrid(pts)
	// water := []*point{}
	for i := 0; i < 10000; i++ {
		// water = tick(grid, water)
		followDrop(grid)
	}
	dumpFile(grid)
}

type point struct{ x, y int }
type tile struct {
	val int
}

func (t *tile) isBlocked() bool {
	return t.val != tileValueSand
}

func (t *tile) isStuck() bool {
	return t.val == tileValueClay || t.val == tileValueWaterStagnant
}

func (t *tile) isWater() bool {
	return t.val == tileValueWaterStagnant || t.val == tileValueWaterFalling
}

func (t *tile) stagnate() {
	t.val = tileValueWaterStagnant
}

func (p *point) hash() float64 {
	return float64(p.x) / float64(p.y)
}

func followDrop(grid [][]*tile) {
	prev := []*point{}
	water := &point{source.x, source.y + 1}
	grid[water.y][water.x].val = tileValueWaterFalling
	// i := 0
	for water.y == len(grid) || grid[water.y][water.x].val != tileValueWaterStagnant {
		prev = append(prev, &point{water.x, water.y})
		// fmt.Println(water, grid[water.y][water.x].val)
		advancePath(grid, water)

		// i++
		// if i == 150 {
		// 	dumpFile(grid)
		// 	os.Exit(0)
		// }
	}

	for _, p := range prev {
		if !(p.x == water.x && p.y == water.y) {
			grid[p.y][p.x].val = tileValueSand
		}
	}
}

func tick(grid [][]*tile, water []*point) []*point {
	freeList := []*point{}
	for _, p := range water {
		n := advance(grid, p)
		if n != nil {
			freeList = append(freeList, n)
		}
	}
	pl := -1
	for len(freeList) != pl {
		pl = len(freeList)
		pass := []*point{}
		for _, p := range freeList {
			n := advance(grid, p)
			if n != nil {
				pass = append(pass, n)
			}
		}
		freeList = pass
	}
	if !grid[source.y+1][source.x].isWater() {
		grid[source.y+1][source.x].val = tileValueWaterFalling
		water = append(water, &point{source.x, source.y + 1})
	}
	return water
}

func advancePath(grid [][]*tile, p *point) {
	t := grid[p.y][p.x]
	left := grid[p.y][p.x-1]
	right := grid[p.y][p.x+1]
	if p.y+1 < len(grid) {
		below := grid[p.y+1][p.x]
		if below.isStuck() {
			if left.isStuck() && right.isBlocked() {
				t.stagnate()
				return
			} else if right.isStuck() && left.isBlocked() {
				t.stagnate()
				return
			} else if !right.isBlocked() {
				right.val = tileValueWaterFalling
				p.x++
				return
			} else if !left.isBlocked() {
				left.val = tileValueWaterFalling
				p.x--
				return
			}
		} else {
			// fall down one tile and update
			below.val = tileValueWaterFalling
			p.y++
			return
		}
	} else {
		return
	}
}

func advance(grid [][]*tile, p *point) *point {
	t := grid[p.y][p.x]
	left := grid[p.y][p.x-1]
	right := grid[p.y][p.x+1]
	if p.y+1 < len(grid) {
		below := grid[p.y+1][p.x]
		if below.isStuck() {
			if left.isStuck() && right.isStuck() {
				// trapped on all sides, can no longer move
				t.stagnate()
				return nil
			} else if !right.isStuck() {
				if left.isStuck() && right.isBlocked() {
					// trapped on left and bottom, right has falling water, stagnate
					t.stagnate()
					return nil
				} else if left.isStuck() && !right.isBlocked() {
					// trapped on left and bottom, move right
					right.val = tileValueWaterFalling
					t.val = tileValueSand
					p.x++
					return nil
				} else if !right.isBlocked() {
					// right free
					right.val = tileValueWaterFalling
					t.val = tileValueSand
					p.x++
					return nil
				} else {
					// right has falling water come back to it
					return p
				}
			} else if !left.isStuck() {
				if right.isStuck() && left.isBlocked() {
					// trapped on left and bottom, right has falling water, stagnate
					t.stagnate()
					return nil
				} else if right.isStuck() && !left.isBlocked() {
					// trapped on left and bottom, move right
					left.val = tileValueWaterFalling
					t.val = tileValueSand
					p.x--
					return nil
				} else if !left.isBlocked() {
					// left free
					left.val = tileValueWaterFalling
					t.val = tileValueSand
					p.x--
					return nil
				} else {
					// left has falling water come back to it
					return p
				}
			} else {
				panic("errorrorororororororororor")
			}
		} else if below.isBlocked() {
			// the tile below is falling water, stop moving
			return nil
		} else {
			// fall down one tile and update
			below.val = tileValueWaterFalling
			t.val = tileValueSand
			p.y++
			return nil
		}
	} else {
		// reached bottom of grid and falls off from here
		return nil
	}
}

func dumpFile(grid [][]*tile) {
	fo, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	for y := 0; y < len(grid); y++ {
		line := ""
		for x := 0; x < len(grid[y]); x++ {
			switch grid[y][x].val {
			case tileValueClay:
				line += "#"
			case tileValueSand:
				line += "."
			case tileValueWaterFalling:
				line += "|"
			case tileValueWaterStagnant:
				line += "~"
			case tileValueSource:
				line += "x"
			}
		}
		fo.Write([]byte(line + "\n"))
	}
	fo.Close()
}

func makeGrid(pts []*point) [][]*tile {
	_, max := bounds(pts)
	grid := make([][]*tile, max.y+1)
	for y := 0; y < max.y+1; y++ {
		grid[y] = make([]*tile, max.x+10)
		for x := 0; x < max.x+10; x++ {
			grid[y][x] = &tile{tileValueSand}
		}
	}
	grid[source.y][source.x].val = tileValueSource
	for _, p := range pts {
		grid[p.y][p.x].val = tileValueClay
	}
	return grid
}

func bounds(pts []*point) (point, point) {
	min := point{math.MaxInt32, math.MaxInt32}
	max := point{math.MinInt32, math.MinInt32}
	for _, p := range pts {
		if p.x < min.x {
			min.x = p.x
		}
		if p.x > max.x {
			max.x = p.x
		}
		if p.y < min.y {
			min.y = p.y
		}
		if p.y > max.y {
			max.y = p.y
		}
	}
	return min, max
}

func parsePoints(lines []string) []*point {
	pts := []*point{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		var isx, isy bool
		var x, y, rngMin, rngMax int
		var err error
		parts[0] = strings.TrimSuffix(parts[0], ",")
		if strings.HasPrefix(parts[0], "x=") {
			isx = true
			x, err = strconv.Atoi(parts[0][2:])
		} else if strings.HasPrefix(parts[0], "y=") {
			isy = true
			y, err = strconv.Atoi(parts[0][2:])
		} else {
			panic("unknown prefix")
		}
		if err != nil {
			panic(err)
		}
		temp := strings.Split(parts[1][2:], "..")
		rngMin, err = strconv.Atoi(temp[0])
		if err != nil {
			panic(err)
		}
		rngMax, err = strconv.Atoi(temp[1])
		if err != nil {
			panic(err)
		}
		for i := rngMin; i <= rngMax; i++ {
			p := &point{}
			if isx {
				p.x = x
				p.y = i
			} else if isy {
				p.y = y
				p.x = i
			} else {
				panic("uh oh")
			}
			pts = append(pts, p)
		}
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
