package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

const file = "input.txt"

func main() {
	entries := read()
	sortEntries(entries)
	guards := parseGuards(entries)
	part1(guards)
	part2(guards)
}

func part1(guards []*guard) {
	max := -1
	var maxG *guard
	for _, g := range guards {
		asleep := g.numMinAsleep()
		if asleep > max {
			max = asleep
			maxG = g
		}
	}
	minute := maxG.minuteAsleepMost()
	fmt.Println(maxG.id * minute)
}

func part2(guards []*guard) {
	max := -1
	var maxG *guard
	for _, g := range guards {
		min := g.minuteAsleepMost()
		times := g.numTimesAsleepFor(min)
		if times > max {
			max = times
			maxG = g
		}
	}
	fmt.Println(maxG.id * maxG.minuteAsleepMost())
}

type guard struct {
	id        int
	start     time.Time
	durations []duration
}

type duration struct {
	start time.Time
	end   time.Time
}

type entry struct {
	time  time.Time
	value string
}

func (d *duration) containsMinute(min int) bool {
	intervalLength := int(d.end.Sub(d.start) / time.Minute)
	for i := 0; i < intervalLength; i++ {
		t := (d.start.Minute() + i) % 60
		if min == t {
			return true
		}
	}
	return false
}

func (g *guard) minuteAsleepMost() int {
	maxTimes := -1
	maxMin := 0
	for i := 0; i < 59; i++ {
		times := g.numTimesAsleepFor(i)
		if times > maxTimes {
			maxTimes = times
			maxMin = i
		}
	}
	return maxMin
}

func (g *guard) numTimesAsleepFor(min int) int {
	times := 0
	for _, d := range g.durations {
		if d.containsMinute(min) {
			times++
		}
	}
	return times
}

func (g *guard) numMinAsleep() int {
	sum := 0
	for _, d := range g.durations {
		num := d.end.Sub(d.start) / time.Minute
		sum += int(num)
	}
	return sum
}

func parseGuards(entries []entry) []*guard {
	guards := []*guard{}
	current := &guard{}
	cd := duration{}

	for _, e := range entries {
		if strings.HasPrefix(e.value, "Guard #") {
			parts := strings.Split(e.value, " ")
			id, err := strconv.Atoi(parts[1][1:])
			if err != nil {
				panic(err)
			}
			if current.id > 0 {
				guards = append(guards, current)
				current = &guard{}
			}
			current.id = id
			current.start = e.time
		} else if e.value == "falls asleep" {
			cd.start = e.time
		} else if e.value == "wakes up" {
			cd.end = e.time
			current.durations = append(current.durations, cd)
			cd = duration{}
		} else {
			panic(e.value)
		}
	}

	if current.id > 0 {
		guards = append(guards, current)
	}
	return consolidateGuards(guards)
}

func consolidateGuards(gs []*guard) []*guard {
	m := map[int]*guard{}

	for _, g := range gs {
		if e, ok := m[g.id]; ok {
			e.durations = append(e.durations, g.durations...)
		} else {
			m[g.id] = g
		}
	}
	guards := []*guard{}
	for _, g := range m {
		guards = append(guards, g)
	}
	return guards
}

func sortEntries(entries []entry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].time.Before(entries[j].time)
	})
}

func read() []entry {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	entries := []entry{}

	for _, line := range lines {
		i := strings.Index(line, "]")
		if i < 0 {
			panic("ahh")
		}

		timeStr := line[1:i]

		t, err := time.Parse("2006-01-02 15:04", timeStr)
		if err != nil {
			panic(err)
		}
		entries = append(entries, entry{
			time:  t,
			value: line[i+2:],
		})
	}
	return entries
}
