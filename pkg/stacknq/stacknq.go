package stacknq

type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toPop := s.items[l]
	s.items = s.items[:l]

	return toPop
}

type Queue struct {
	items []int
}

// Enqueue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue
func (q *Queue) Dequeue() int {
	toDequeue := q.items[0]
	q.items = q.items[1:]

	return toDequeue
}
