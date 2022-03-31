package main

import (
	"fmt"
)

type Node struct {
	Value string
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func NewStack() *Stack {
	return &Stack{}
}

// Stack adalah tumpukan LIFO dasar yang diubah ukurannya sesuai kebutuhan.
type Stack struct {
	nodes []*Node
	count int
}

// Push menambahkan simpul ke tumpukan.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop menghapus dan mengembalikan simpul dari tumpukan di urutan terakhir ke urutan pertama.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// NewQueue mengembalikan antrian baru dengan ukuran awal yang diberikan.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Queue adalah antrean FIFO dasar berdasarkan daftar
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push menambahkan node ke antrian.
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop menghapus dan mengembalikan simpul dari antrian di urutan pertama hingga terakhir
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func main() {
	s := NewStack()
	s.Push(&Node{"Dewi"})
	s.Push(&Node{"Elsa"})
	s.Push(&Node{"fan"})
	fmt.Println("Tumpukan")

	fmt.Println(s.Pop())
	fmt.Println("----------")
	fmt.Println(s.Pop())
	fmt.Println("----------")
	fmt.Println(s.Pop())
	fmt.Println("----------")
	fmt.Println("")
	fmt.Println("Antrian")

	q := NewQueue(1)
	q.Push(&Node{"Mirda"})
	q.Push(&Node{"Oris"})
	q.Push(&Node{"Yola"})
	fmt.Println(q.Pop(), "|", q.Pop(), "|", q.Pop())
}
