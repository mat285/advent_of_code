package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	ds := read()
	tree, _ := buildTree(ds, 0)
	part1(tree)
	part2(tree)
}

func part1(tree *node) {
	fmt.Println(tree.sum())
}

func part2(tree *node) {
	fmt.Println(tree.value())
}

type header struct {
	children int
	entries  int
}

type node struct {
	header   *header
	children []*node
	entries  []int
}

func (n *node) sum() int {
	sum := 0
	for _, m := range n.entries {
		sum += m
	}
	for _, c := range n.children {
		sum += c.sum()
	}
	return sum
}

func (n *node) value() int {
	if len(n.children) > 0 {
		sum := 0
		for _, e := range n.entries {
			if e == 0 || e-1 >= len(n.children) {
				continue
			}
			sum += n.children[e-1].value()
		}
		return sum

	}
	return n.sum()
}

func buildTree(ds []string, start int) (*node, int) {
	i := start
	n := &node{}
	nc, err := strconv.Atoi(ds[i])
	if err != nil {
		panic(err)
	}
	i++
	ne, err := strconv.Atoi(ds[i])
	if err != nil {
		panic(err)
	}
	i++
	n.header = &header{nc, ne}

	for j := 0; j < n.header.children; j++ {
		child, end := buildTree(ds, i)
		i = end
		n.children = append(n.children, child)
	}

	if len(ds)-i < n.header.entries {
		panic("uh oh spaghetti-os")
	}

	for j := 0; j < n.header.entries; j++ {
		m, err := strconv.Atoi(ds[i])
		if err != nil {
			panic(err)
		}
		i++
		n.entries = append(n.entries, m)
	}

	return n, i
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), " ")
}
