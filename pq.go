// Package priorityqueue provides a generic heap-backed priority queue.
package priorityqueue

import "container/heap"

// Queue is a priority queue. The element for which less reports true
// against all others is popped first. The zero value is not usable; use New.
type Queue[T any] struct{ h *items[T] }

// New returns an empty queue ordered by less.
// Use func(a, b T) bool { return a < b } for a min-queue.
func New[T any](less func(a, b T) bool) *Queue[T] {
	return &Queue[T]{h: &items[T]{less: less}}
}

// Push adds v to the queue in O(log n).
func (q *Queue[T]) Push(v T) { heap.Push(q.h, v) }

// Pop removes and returns the highest-priority element in O(log n).
// ok is false if the queue is empty.
func (q *Queue[T]) Pop() (v T, ok bool) {
	if q.Len() == 0 {
		return v, false
	}
	return heap.Pop(q.h).(T), true
}

// Peek returns the highest-priority element without removing it.
// ok is false if the queue is empty.
func (q *Queue[T]) Peek() (v T, ok bool) {
	if q.Len() == 0 {
		return v, false
	}
	return q.h.s[0], true
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int { return len(q.h.s) }

type items[T any] struct {
	s    []T
	less func(a, b T) bool
}

func (h *items[T]) Len() int           { return len(h.s) }
func (h *items[T]) Less(i, j int) bool { return h.less(h.s[i], h.s[j]) }
func (h *items[T]) Swap(i, j int)      { h.s[i], h.s[j] = h.s[j], h.s[i] }
func (h *items[T]) Push(x any)         { h.s = append(h.s, x.(T)) }
func (h *items[T]) Pop() any {
	var zero T
	n := len(h.s) - 1
	v := h.s[n]
	h.s[n] = zero // drop reference for GC
	h.s = h.s[:n]
	return v
}
