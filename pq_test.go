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

func TestEmpty(t *testing.T) {
	q := New(func(a, b int) bool { return a < b })
	if v, ok := q.Peek(); ok || v != 0 {
		t.Fatalf("Peek on empty = %d, %v; want 0, false", v, ok)
	}
	if q.Len() != 0 {
		t.Fatalf("Len = %d, want 0", q.Len())
	}
	q.Push(1)
	q.Pop()
	if _, ok := q.Pop(); ok {
		t.Fatal("Pop after draining returned ok")
	}
}

func TestPeekDoesNotRemove(t *testing.T) {
	q := New(func(a, b int) bool { return a < b })
	q.Push(5)
	q.Peek()
	q.Peek()
	if q.Len() != 1 {
		t.Fatalf("Len = %d, want 1", q.Len())
	}
}

func TestMaxQueueWithDuplicates(t *testing.T) {
	q := New(func(a, b int) bool { return a > b })
	for _, v := range []int{3, 1, 3, 2, 1, 3} {
		q.Push(v)
	}
	want := []int{3, 3, 3, 2, 1, 1}
	for i, w := range want {
		if v, ok := q.Pop(); !ok || v != w {
			t.Fatalf("Pop #%d = %d, %v; want %d, true", i, v, ok, w)
		}
	}
}

func TestStructPriority(t *testing.T) {
	type task struct {
		name string
		pri  int
	}
	q := New(func(a, b task) bool { return a.pri < b.pri })
	q.Push(task{"low", 9})
	q.Push(task{"urgent", 0})
	q.Push(task{"mid", 5})
	for _, want := range []string{"urgent", "mid", "low"} {
		if v, _ := q.Pop(); v.name != want {
			t.Fatalf("Pop = %q, want %q", v.name, want)
		}
	}
}

// TestInterleaved compares random Push/Pop sequences against a sorted slice.
func TestInterleaved(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	q := New(func(a, b int) bool { return a < b })
	var ref []int
	for i := 0; i < 10000; i++ {
		if r.Intn(3) > 0 {
			v := r.Intn(100)
			q.Push(v)
			ref = append(ref, v)
			slices.Sort(ref)
			continue
		}
		v, ok := q.Pop()
		if ok != (len(ref) > 0) {
			t.Fatalf("step %d: Pop ok = %v with %d refs", i, ok, len(ref))
		}
		if ok {
			if v != ref[0] {
				t.Fatalf("step %d: Pop = %d, want %d", i, v, ref[0])
			}
			ref = ref[1:]
		}
		if q.Len() != len(ref) {
			t.Fatalf("step %d: Len = %d, want %d", i, q.Len(), len(ref))
		}
	}
}
