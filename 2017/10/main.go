package main

import "fmt"

var lengths = []int{106, 118, 236, 1, 130, 0, 235, 254, 59, 205, 2, 87, 129, 25, 255, 118}

func main() {
	part1(duplicate(lengths))
}

func part1(input []int) {
	r := newRing(iter(256))
	for _, i := range input {
		r.act(i)
	}
	fmt.Println(int(r.buf[0]) * int(r.buf[1]))
}

type ring struct {
	buf      []byte
	pos      int
	skipSize int
}

func newRing(buf []byte) *ring {
	return &ring{
		buf:      buf,
		pos:      0,
		skipSize: 0,
	}
}

func (r *ring) skip(i int) {
	r.pos = (r.pos + i) % len(r.buf)
}

func (r *ring) reverse(l int) {
	for i := 0; i < l/2; i++ {
		start := (r.pos + i) % len(r.buf)
		end := (r.pos + l - i - 1) % len(r.buf)
		v1 := r.buf[start]
		r.buf[start] = r.buf[end]
		r.buf[end] = v1
	}
}

func (r *ring) act(l int) {
	r.reverse(l)
	r.skip(l + r.skipSize)
	r.skipSize++
}

func iter(l int) []byte {
	ret := make([]byte, l)
	for i := 0; i < l; i++ {
		ret[i] = byte(i)
	}
	return ret
}

func duplicate(i []int) []int {
	dst := make([]int, len(i))
	copy(dst, i)
	return dst
}
