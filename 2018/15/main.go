package main

import (
	"io/ioutil"
	"strings"
)

const (
	power  = 3
	health = 200
)

func main() {}

type point struct{ x, y int }

type positionable interface {
	position() point
}

type elf struct {
	power  int
	health int
	pos    point
}

type goblin struct {
	power  int
	health int
	pos    point
}

type entity struct {
	elf    *elf
	goblin *goblin
	wall   bool
}

func newElf(p point) *elf {
	return &elf{
		power:  power,
		health: health,
		pos:    p,
	}
}

func newGoblin(p point) *goblin {
	return &goblin{
		power:  power,
		health: health,
		pos:    p,
	}
}

func sortPositionable(ps []positionable)

func less(c1, c2 positionable) bool {
	if c1.position().y < c2.position().y {
		return true
	} else if c1.position().y == c2.position().y {
		return c1.position().x < c2.position().x
	}
	return false
}

func parse(lines []string) [][]*entity {
	grid := make([][]*entity, len(lines))

	for i, line := range lines {
		bs := []byte(line)
		grid[i] = make([]*entity, len(bs))
		for j, b := range bs {
			p := point{j, i}
			e := &entity{}
			switch b {
			case 'G':
				e.goblin = newGoblin(p)
			case 'E':
				e.elf = newElf(p)
			case '#':
				e.wall = true
			default:
				panic("noooo")
			}
			grid[i][j] = e
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
