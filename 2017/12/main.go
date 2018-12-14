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
	part2(lines)
}

func part1(lines []string) {
	m := parsePrograms(lines)
	fmt.Println(len(group(0, m)))
}

func part2(lines []string) {
	m := parsePrograms(lines)
	groups := []map[int]bool{}
	for len(m) > 0 {
		k := selectKey(m)
		g := group(k, m)
		groups = append(groups, g)
		for id := range g {
			delete(m, id)
		}
	}
	fmt.Println(len(groups))
}

type program struct {
	id    int
	pipes map[int]bool
}

func newProgram(id int) *program {
	return &program{
		id:    id,
		pipes: map[int]bool{},
	}
}

func group(id int, m map[int]*program) map[int]bool {
	ret := map[int]bool{}
	working := map[int]*program{}
	working[id] = m[id]

	for len(working) > 0 {
		pid := selectKey(working)
		p := m[pid]
		delete(working, pid)
		if _, ok := ret[pid]; ok {
			continue
		}
		ret[pid] = true
		for k := range p.pipes {
			working[k] = m[k]
		}
	}

	return ret
}

func selectKey(m map[int]*program) int {
	for i := range m {
		return i
	}
	return -1
}

func parsePrograms(lines []string) map[int]*program {
	m := map[int]*program{}
	for _, line := range lines {
		parts := strings.Split(line, " ")

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		if _, ok := m[id]; !ok {
			m[id] = newProgram(id)
		}

		for _, s := range parts[2:] {
			pid, err := strconv.Atoi(strings.Trim(s, ","))
			if err != nil {
				panic(err)
			}
			m[id].pipes[pid] = true

			if _, ok := m[pid]; !ok {
				m[pid] = newProgram(pid)
			}
			m[pid].pipes[id] = true
		}
	}
	return m
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split((string(data[:len(data)-1])), "\n")
}
