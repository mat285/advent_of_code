package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

const realInput = 3463
const size = 300

const testInput = 18

var input = realInput

func main() {
	part1()
	part2()
}

func part1() {
	max := math.MinInt32
	var mp point
	memo := map[string]int{}
	for x := 1; x <= size-3; x++ {
		for y := 1; y <= size-3; y++ {
			v := valueSize(x, y, 3, memo)
			if v > max {
				max = v
				mp.x = x
				mp.y = y
			}
		}
	}
	fmt.Println(mp)
}

func part2() {
	var mp point
	mp.v = math.MinInt32
	memo := &protected{}
	// for s := 1; s <= size; s++ {
	// 	for x := 1; x <= size-s; x++ {
	// 		for y := 1; y <= size-s; y++ {
	// 			fmt.Println(x, y, s)
	// 			v := valueSize(x, y, s, memo)
	// 			if v > max {
	// 				max = v
	// 				mp.x = x
	// 				mp.y = y
	// 			}
	// 		}
	// 	}
	// }
	memo.result = make(chan point, 20)

	prev := 0

	for i := 0; i < 10; i++ {
		var min, max point
		min.x = prev + 1
		min.y = 1
		max.x = prev + size/10
		max.y = size
		prev = max.x
		memo.wg.Add(1)
		go search(min, max, memo)
	}

	memo.wg.Wait()
	l := len(memo.result)
	for i := 0; i < l; i++ {
		p := <-memo.result
		if p.v > mp.v {
			mp = p
		}
	}

	fmt.Println(mp)
}

func search(min, max point, memo *protected) {
	defer memo.wg.Done()
	d := map[string]int{}
	var mp point
	mp.v = math.MinInt32
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			for s := 1; s <= size; s++ {
				// fmt.Println(x, y, s)
				v := valueSize(x, y, s, d)
				if v > mp.v {
					mp.v = v
					mp.x = x
					mp.y = y
				}
			}
		}
	}
	memo.result <- mp
}

type point struct{ x, y, v int }

type protected struct {
	wg     sync.WaitGroup
	result chan point
}

func valueSize(x, y, s int, memo map[string]int) int {
	if s == 0 {
		return 0
	}
	str := fmt.Sprintf("%d,%d,%d", x, y, s)
	if val, ok := memo[str]; ok {
		return val
	}

	vs := valueSize(x, y, s-1, memo)

	v := 0
	for i := 0; i < s; i++ {
		v += value(x+i, y+s-1)
		v += value(x+s-1, y+i)
	}
	v -= value(x+s-1, y+s-1)
	memo[str] = v + vs
	return v + vs
}

func value(x, y int) int {
	rID := x + 10
	pl := rID * y
	pl += input
	pl *= rID
	pstr := strconv.Itoa(pl)
	vpl := 0
	if len(pstr) >= 3 {
		vpl, _ = strconv.Atoi(string(pstr[len(pstr)-3]))
	}
	return vpl - 5
}
