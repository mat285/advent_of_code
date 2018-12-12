package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	lines := read()
	ins := parseInstructions(lines)
	part1(ins)
	part2(ins)
}

func part1(ins []*instruction) {
	rs := newRegisters()
	rs.operate(ins...)
	fmt.Println(rs.max())
}

func part2(ins []*instruction) {
	rs := newRegisters()
	m := math.MinInt32
	for _, in := range ins {
		rs.operate(in)
		v := rs.max()
		if v > m {
			m = v
		}
	}
	fmt.Println(m)
}

type instruction struct {
	reg  string
	op   string
	v    int
	cond string
	comp string
	cv   int
}

type register struct {
	name  string
	value int
}

type registers struct {
	rs map[string]*register
}

func newRegisters() *registers {
	return &registers{
		rs: make(map[string]*register),
	}
}

func (r *registers) get(name string) *register {
	if _, ok := r.rs[name]; !ok {
		r.rs[name] = &register{name: name, value: 0}
	}
	return r.rs[name]
}

func (r *registers) add(name string, v int) {
	r.get(name).value += v
}

func (r *registers) max() int {
	m := math.MinInt32
	for _, r := range r.rs {
		if r.value > m {
			m = r.value
		}
	}
	return m
}

func (r *registers) should(in *instruction) bool {
	cond := r.get(in.cond)
	switch in.comp {
	case "==":
		return cond.value == in.cv
	case "!=":
		return cond.value != in.cv
	case ">":
		return cond.value > in.cv
	case "<":
		return cond.value < in.cv
	case "<=":
		return cond.value <= in.cv
	case ">=":
		return cond.value >= in.cv
	default:
		panic(in.comp)
	}
}

func (r *registers) operate(ins ...*instruction) {
	for _, in := range ins {
		if !r.should(in) {
			continue
		}
		switch in.op {
		case "inc":
			r.add(in.reg, in.v)
		case "dec":
			r.add(in.reg, -in.v)
		default:
			panic(in.op)
		}
	}
}

func parseInstructions(lines []string) []*instruction {
	ret := []*instruction{}

	for _, line := range lines {
		in := &instruction{}
		_, err := fmt.Sscanf(line, "%s %s %d if %s %s %d", &in.reg, &in.op, &in.v, &in.cond, &in.comp, &in.cv)
		if err != nil {
			panic(err)
		}
		ret = append(ret, in)
	}
	return ret
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
