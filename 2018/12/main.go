package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const initial = "#.####...##..#....#####.##.......##.#..###.#####.###.##.###.###.#...#...##.#.##.#...#..#.##..##.#.##"

const alive = "#"
const dead = "."

func main() {
	lines := read()
	rs := parseRules(lines)
	part1(rs)
	part2(rs)
}

func part1(rs []*rule) {
	head := parsePots(initial)
	for i := 0; i < 20; i++ {
		head = addBufferN(head, 4)
		head = tick(head, rs)
	}
	fmt.Println(sum(head))
}

func part2(rs []*rule) {
	head := parsePots(initial)
	prev := sum(head)
	change := -1
	i := 0
	for ; change != 0; i++ {
		head = addBufferN(head, 4)
		head = contractToN(head, 4)
		head = tick(head, rs)
		s := sum(head)
		if s-prev == change {
			break
		}
		change = s - prev
		prev = s
	}
	curr := sum(head)
	curr += (50000000000 - i - 1) * change
	fmt.Println(curr)
}

type rule struct {
	left  string
	curr  string
	right string

	outcome string
}

type pot struct {
	id    int
	state string
	next  *pot
	prev  *pot

	collapsed bool
}

func sum(p *pot) int {
	p = toHead(p)
	sum := 0
	for p != nil {
		if p.state == alive {
			sum += p.id
		}
		p = p.next
	}
	return sum
}

func tick(p *pot, rs []*rule) *pot {
	for p.prev != nil {
		p = p.prev
	}

	var head *pot
	var temp *pot
	for p != nil {
		n := apply(p, rs)
		p = p.next
		if head == nil {
			head = n
			temp = n
			continue
		}
		temp.next = n
		n.prev = temp
		temp = n
	}
	return head
}

func apply(p *pot, rs []*rule) *pot {
	var left, right string
	if p.prev != nil {
		if p.prev.prev != nil {
			left = p.prev.prev.state
		} else {
			left = dead
		}
		left += p.prev.state
	} else {
		left = dead + dead
	}
	if p.next != nil {
		right += p.next.state
		if p.next.next != nil {
			right += p.next.next.state
		} else {
			right += dead
		}
	} else {
		right = dead + dead
	}
	curr := p.state

	next := nextGenSingle(left, curr, right, rs)
	return &pot{
		id:    p.id,
		state: next,
	}
}

func nextGenSingle(left, curr, right string, rs []*rule) string {
	for _, r := range rs {
		if doesApply(left, curr, right, r) {
			return r.outcome
		}
	}
	return curr
}

func doesApply(left, curr, right string, r *rule) bool {
	return r.left == left && r.curr == curr && r.right == right
}

func contractToN(p *pot, n int) *pot {
	p = toHead(p)
	t := p
	c := 0
	for t.next != nil && t.state == dead {
		c++
		t = t.next
	}

	head := p
	if c > n {
		head = head.next
		head.prev = nil
		c--
	}

	t = toTail(head)

	c = 0
	for t.prev != nil && t.state == dead {
		c++
		t = t.prev
	}

	t = toTail(t)

	if c > n {
		t = t.prev
		t.next = nil
		c--
	}

	return head
}

func addBufferN(p *pot, n int) *pot {
	t := toHead(p)

	for i := 0; i < n; i++ {
		t.prev = &pot{
			id:    t.id - 1,
			state: dead,
			next:  t,
		}
		t = t.prev
	}

	head := t

	t = toTail(t)

	for i := 0; i < n; i++ {
		t.next = &pot{
			id:    t.id + 1,
			state: dead,
			prev:  t,
		}
		t = t.next
	}
	return head
}

func toTail(p *pot) *pot {
	for p.next != nil {
		p = p.next
	}
	return p
}

func toHead(p *pot) *pot {
	for p.prev != nil {
		p = p.prev
	}
	return p
}

func parsePots(line string) *pot {
	var head *pot
	var temp *pot
	for i := 0; i < len(line); i++ {
		p := &pot{
			id:    i,
			state: string(line[i]),
		}
		if head == nil {
			head = p
			temp = p
			continue
		}
		temp.next = p
		p.prev = temp
		temp = p
	}
	return head
}

func parseRules(lines []string) []*rule {
	ret := []*rule{}
	for _, line := range lines {
		ret = append(ret, &rule{
			left:    line[0:2],
			curr:    string(line[2]),
			right:   line[3:5],
			outcome: string(line[len(line)-1]),
		})
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
