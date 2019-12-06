package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	g := graph(lines)
	sum := 0
	for _, v := range g {
		sum += countOrbits(v)
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	g := graph(lines)
	fmt.Println(distance(g[you], g[san]))
}

const (
	com = "COM"
	you = "YOU"
	san = "SAN"
)

type object struct {
	id     string
	parent *object
}

func distance(o1, o2 *object) int {
	parents := parents(o1)
	var nearest *object
	a := o2
	for a != nil {
		if _, has := parents[a.id]; has {
			nearest = a
			break
		}
		a = a.parent
	}
	return pathToParent(o1, nearest) + pathToParent(o2, nearest)
}

func pathToParent(o, p *object) int {
	count := 0
	for o.parent != nil && o.parent.id != p.id {
		count++
		o = o.parent
	}
	return count
}

func parents(o *object) map[string]*object {
	parents := map[string]*object{}
	for o.parent != nil {
		parents[o.parent.id] = o.parent
		o = o.parent
	}
	return parents
}

func countOrbits(o *object) int {
	count := 0
	for o.parent != nil {
		count++
		o = o.parent
	}
	return count
}

func graph(lines []string) map[string]*object {
	root := &object{id: com}
	g := map[string]*object{
		com: root,
	}
	for _, line := range lines {
		parts := strings.Split(line, ")")
		parent, child := parts[0], parts[1]

		if _, ok := g[parent]; !ok {
			g[parent] = &object{id: parent}
		}
		if _, ok := g[child]; !ok {
			g[child] = &object{id: child}
		}
		g[child].parent = g[parent]
	}
	return g
}
