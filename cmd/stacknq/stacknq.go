package main

import (
	"fmt"

	sq "github.com/h-mavrodiev/data-structures-and-algorithms/pkg/stacknq"
)

func main() {
	s := sq.Stack{}

	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("Stack after Push: %d\n", s)
	pop := s.Pop()
	fmt.Printf("Poped item: %d\n", pop)
	fmt.Printf("Stack after Pop: %d\n", s)

	q := sq.Queue{}
	fmt.Print("#################################\n")
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Printf("Queue after Enqueue: %d\n", q)
	deq := q.Dequeue()
	fmt.Printf("Dequed item: %d\n", deq)
	fmt.Printf("Queue after Dequeue: %d\n", q)
}
