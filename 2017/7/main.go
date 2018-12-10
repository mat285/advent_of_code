package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	lines := read()
	part1(lines)
}

func part1(lines []string) {
	entries := parseEntries(lines)
	n := buildTree(entries)
	fmt.Println(n.name)
}

func part2(lines []string) {
	// entries := parseEntries(lines)
	// n := buildTree(entries)

}

type node struct {
	name     string
	weight   int
	children []*node
}

type entry struct {
	name     string
	weight   int
	children []string
}

func (n *node) value() int {
	sum := n.weight
	for _, c := range n.children {
		sum += c.value()
	}
	return sum
}

func (n *node) unbalanced() *node {
	values := map[int][]*node{}
	for _, c := range n.children {
		v := c.value()
		values[v] = append(values[v], c)
	}
	return nil
}

func buildTree(entries []*entry) *node {
	m := map[string]*node{}
	free := map[string]*node{}
	for _, e := range entries {
		var n *node
		var ok bool
		if n, ok = m[e.name]; ok {
			n.weight = e.weight
		} else {
			n = &node{
				name:   e.name,
				weight: e.weight,
			}
			m[n.name] = n
			free[n.name] = n
		}
		for _, c := range e.children {
			if child, ok := m[c]; ok {
				n.children = append(n.children, child)
				delete(free, child.name)
			} else {
				child = &node{name: c}
				m[child.name] = child
				n.children = append(n.children, child)
			}
		}
	}

	if len(free) != 1 {
		panic("ahhh")
	}
	for _, n := range free {
		return n
	}
	return nil
}

func parseEntries(lines []string) []*entry {
	entries := []*entry{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		w, err := strconv.Atoi(parts[1][1 : len(parts[1])-1])
		if err != nil {
			panic(err)
		}
		cs := []string{}
		if len(parts) > 2 {
			for _, p := range parts[3:] {
				cs = append(cs, strings.Trim(p, ","))
			}
		}
		entries = append(entries, &entry{
			name:     parts[0],
			weight:   w,
			children: cs,
		})
	}
	return entries
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
