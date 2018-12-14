package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = 330121

var inputArr = []int{3, 3, 0, 1, 2, 1}

func main() {
	part1()
	part2()
}

func part1() {
	num := input
	r := newRing(initialNodes())

	for i := 0; i < num+10; i++ {
		r.tick()
	}
	vals := next(10, r.get(num))
	str := ""
	for _, v := range vals {
		str += strconv.Itoa(v)
	}
	fmt.Println(str)
}

func part2() {
	arr := inputArr
	r := newRing(initialNodes())
	var idx *node
	var startNode *node
	stepSize := 10000
	initalSkip := 15500000 // hack

	for i := 0; i < initalSkip; i++ {
		r.tick()
	}

	if initalSkip > len(inputArr) {
		startNode = r.tail
		for range inputArr {
			startNode = startNode.prev
		}
	}

	for idx == nil {
		for i := 0; i < stepSize; i++ {
			r.tick()
		}
		idx = r.search(arr, startNode)

		startNode = r.tail
		for range inputArr {
			startNode = startNode.prev
		}
	}
	fmt.Println(r.countBefore(idx))
}

type node struct {
	val  int
	prev *node
	next *node
}

type ring struct {
	head     *node
	tail     *node
	pointers []*node
	size     int
}

func newRing(head, tail *node, pointers []*node) *ring {
	return &ring{
		head:     head,
		tail:     tail,
		pointers: pointers,
		size:     len(pointers),
	}
}

func initialNodes() (*node, *node, []*node) {
	n1 := &node{val: 3}
	n2 := &node{val: 7}
	n1.next = n2
	n1.prev = n2
	n2.next = n1
	n2.prev = n1 // circular
	return n1, n2, []*node{n1, n2}
}

func (r *ring) print() {
	temp := r.head
	for temp.next != r.head {
		if temp == r.pointers[0] {
			fmt.Printf("(%d) ", temp.val)
		} else if temp == r.pointers[1] {
			fmt.Printf("[%d] ", temp.val)
		} else {
			fmt.Print(temp.val, " ")
		}
		temp = temp.next
	}
	fmt.Print(temp.val)
	fmt.Println()
}

func (r *ring) String(optional *node) string {
	temp := optional
	if optional == nil {
		temp = r.head
	}
	str := ""
	for temp.next != r.head {
		str += strconv.Itoa(temp.val)
		temp = temp.next
	}
	str += strconv.Itoa(temp.val)
	return str
}

func (r *ring) search(vals []int, optional *node) *node {
	sub := ""
	for _, v := range vals {
		sub += strconv.Itoa(v)
	}
	str := r.String(optional)
	idx := strings.Index(str, sub)
	if idx < 0 {
		return nil
	}
	if optional == nil {
		return r.get(idx)
	}
	for i := 0; i < idx; i++ {
		optional = optional.next
	}
	return optional
}

func (r *ring) countBefore(n *node) int {
	count := 0
	for n != r.head {
		n = n.prev
		count++
	}
	return count
}

func (r *ring) tick() {
	sum := 0
	for _, n := range r.pointers {
		sum += n.val
	}
	str := strconv.Itoa(sum)
	vals := []int{}
	for _, char := range str {
		v, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		vals = append(vals, v)
	}
	r.insert(vals...)

	for p := range r.pointers {
		steps := r.pointers[p].val + 1
		r.move(p, steps)
	}
}

func (r *ring) get(idx int) *node {
	temp := r.head
	for i := 0; i < idx; i++ {
		temp = temp.next
	}
	return temp
}

func (r *ring) insert(vals ...int) {
	for _, v := range vals {
		n := &node{val: v}
		r.tail.next = n
		n.prev = r.tail
		n.next = r.head
		r.head.prev = n
		r.tail = n
		r.size++
	}
}

func (r *ring) move(p int, steps int) {
	for i := 0; i < steps; i++ {
		r.pointers[p] = r.pointers[p].next
	}
}

func next(n int, p *node) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = p.val
		p = p.next
	}
	return ret
}
