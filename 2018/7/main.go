package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strings"
)

const adjust = 65
const numWorkers = 5
const baseline = 60

var testLines = strings.Split(`Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`, "\n")

func main() {
	lines := read()
	// lines = testLines
	graph, free := buildGraph(lines)
	fmt.Println(part1(graph, free))

	graph, free = buildGraph(lines)
	fmt.Println(part2(graph, free))
}

func part2(graph map[string]*node, free heap.Interface) int {
	workers := []*worker{}
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, newWorker(i))
	}

	path := []*node{}
	counter := -1
	finished := 0
	for finished != len(graph) {
		counter++

		// free all workers first so we have all free nodes in the heap
		for _, w := range workers {
			if w.isDone(counter) {
				path = append(path, w.finish(free))
				finished++
			}
		}

		for _, w := range workers {
			if w.isFree() && free.Len() > 0 {
				w.startTask(popNode(free), counter)
			}
		}
	}

	return counter
}

func part1(graph map[string]*node, free heap.Interface) string {
	path := []*node{}
	for free.Len() > 0 {
		n := popNode(free)

		if len(n.dependencies) > 0 {
			panic("something went wrong")
		}
		path = append(path, n)

		for _, d := range n.dependents {
			delete(d.dependencies, n.id)
			if len(d.dependencies) == 0 {
				heap.Push(free, d)
			}
		}
	}

	if len(path) != len(graph) {
		panic("completed with missing nodes")
	}

	str := ""
	for _, node := range path {
		str += node.id
	}

	return str
}

type node struct {
	id           string
	dependents   map[string]*node
	dependencies map[string]*node
}

type nodeHeap struct {
	nodes []*node
}

type worker struct {
	id    int
	task  *node
	start int
}

func newNode(id string) *node {
	return &node{
		id:           id,
		dependents:   make(map[string]*node),
		dependencies: make(map[string]*node),
	}
}

func newHeap(nodes []*node) heap.Interface {
	h := &nodeHeap{nodes}
	heap.Init(h)
	return h
}

func newWorker(id int) *worker {
	return &worker{
		id: id,
	}
}

func (h *nodeHeap) Len() int {
	return len(h.nodes)
}

func (h *nodeHeap) Less(i, j int) bool {
	return h.nodes[i].id < h.nodes[j].id
}

func (h *nodeHeap) Swap(i, j int) {
	t := h.nodes[i]
	h.nodes[i] = h.nodes[j]
	h.nodes[j] = t
}

func (h *nodeHeap) Push(n interface{}) {
	typed, ok := n.(*node)
	if !ok {
		panic("wrong type")
	}
	h.nodes = append(h.nodes, typed)
}

func (h *nodeHeap) Pop() interface{} {
	n := h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	return n
}

func (w *worker) startTask(task *node, start int) {
	w.task = task
	w.start = start
}

func (w *worker) isDone(time int) bool {
	if w.task == nil {
		return false
	}
	return time > w.start+taskTime(w.task)
}

func (w *worker) isFree() bool {
	return w.task == nil
}

func (w *worker) finish(pq heap.Interface) *node {
	w.start = 0
	for _, d := range w.task.dependents {
		delete(d.dependencies, w.task.id)
		if len(d.dependencies) == 0 {
			heap.Push(pq, d)
		}
	}
	r := w.task
	w.task = nil
	return r
}

func taskTime(n *node) int {
	if len(n.id) > 1 {
		panic("unsupported")
	}
	return int(byte(n.id[0])) - adjust + baseline
}

func popNode(pq heap.Interface) *node {
	ready := heap.Pop(pq)
	n, ok := ready.(*node)
	if !ok {
		panic("Wrong type")
	}
	return n
}

func buildGraph(lines []string) (graph map[string]*node, pq heap.Interface) {
	graph = map[string]*node{}
	free := map[string]*node{}
	for _, line := range lines {
		var parent, child *node
		var ok bool
		parentID, childID := parseLine(line)
		if parent, ok = graph[parentID]; !ok {
			// if the parent doesn't exist, add it to graph and free list
			parent = newNode(parentID)
			graph[parent.id] = parent
			free[parent.id] = parent
		}
		if child, ok = graph[childID]; !ok {
			child = newNode(childID)
			graph[child.id] = child
		} else if _, ok = free[childID]; ok {
			delete(free, childID) // if the child is in the free list, remove it
		}
		parent.dependents[child.id] = child
		child.dependencies[parent.id] = parent
	}
	pq = buildHeap(free)
	return
}

func buildHeap(free map[string]*node) heap.Interface {
	slice := []*node{}
	for _, n := range free {
		slice = append(slice, n)
	}
	return newHeap(slice)
}

func parseLine(line string) (string, string) {
	parts := strings.Split(line, " ")
	return parts[1], parts[7]
}

func read() []string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}
