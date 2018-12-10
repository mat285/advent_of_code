package main

import (
	"fmt"
)

const (
	input          = "428 players; last marble is worth 72061 points"
	realNumPlayers = 428
	realLastMarble = 72061

	testNumPlayers = 10
	testLastMarble = 1618

	special   = 23
	stepsBack = 7
)

var (
	numPlayers = realNumPlayers
	lastMarble = realLastMarble
)

func main() {
	part1()
	part2()
}

func part1() {
	players := []*player{}
	for i := 0; i < numPlayers; i++ {
		players = append(players, newPlayer(i))
	}
	c := &circle{}
	curr := 0
	for m := 0; m <= lastMarble; m++ {
		player := players[curr]
		if m != 0 && m%special == 0 {
			player.add(m)
			c.moveCounterClockwise(stepsBack)
			v := c.removeCurrent()
			player.add(v)

		} else {
			c.moveClockwise(1)
			c.addClockwise(m)
		}
		curr = (curr + 1) % numPlayers
	}

	var maxP *player
	for _, p := range players {
		if maxP == nil || maxP.score < p.score {
			maxP = p
		}
	}
	fmt.Println(maxP.score)
}

func part2() {
	lastMarble *= 100
	part1()
}

type player struct {
	id    int
	score int64
}

type node struct {
	val  int
	prev *node // counter clockwise
	next *node // clockwise
}

type circle struct {
	current *node
}

func newPlayer(id int) *player {
	return &player{
		id: id,
	}
}

func (p *player) add(s int) {
	p.score += int64(s)
}

func (c *circle) moveClockwise(steps int) {
	if c.current == nil {
		return
	}
	for i := 0; i < steps; i++ {
		c.current = c.current.next
	}
}

func (c *circle) moveCounterClockwise(steps int) {
	if c.current == nil {
		return
	}
	for i := 0; i < steps; i++ {
		c.current = c.current.prev
	}
}

// removes current and steps clockwise
func (c *circle) removeCurrent() int {
	if c.current == nil {
		return -1
	}
	if c.current.next == c.current {
		v := c.current.val
		c.current = nil
		return v
	}
	n := c.current
	n.prev.next = n.next
	n.next.prev = n.prev
	c.current = n.next
	return n.val
}

// adds clockwise and sets current
func (c *circle) addClockwise(val int) {
	n := &node{val: val}
	if c.current == nil {
		c.current = n
		n.next = n
		n.prev = n
		return
	}
	n.prev = c.current
	n.next = c.current.next
	c.current.next.prev = n
	c.current.next = n
	c.current = n
}
