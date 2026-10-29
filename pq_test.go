package priorityqueue

import (
	"math/rand"
	"slices"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New(func(a, b int) bool { return a < b })
	if _, ok := q.Pop(); ok {
		t.Fatal("Pop on empty queue returned ok")
	}
	in := rand.Perm(100)
	for _, v := range in {
		q.Push(v)
	}
	if v, _ := q.Peek(); v != 0 {
		t.Fatalf("Peek = %d, want 0", v)
	}
	var out []int
	for q.Len() > 0 {
		v, _ := q.Pop()
		out = append(out, v)
	}
	if !slices.IsSorted(out) || len(out) != len(in) {
		t.Fatalf("pop order not sorted: %v", out)
	}
}
