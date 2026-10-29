# priorityqueue

A generic, heap-backed priority queue for Go.

```sh
go get github.com/glenpi/priorityqueue
```

```go
q := priorityqueue.New(func(a, b int) bool { return a < b }) // min-queue
q.Push(3)
q.Push(1)
v, ok := q.Pop() // 1, true
```

Push and Pop run in O(log n). Peek and Len run in O(1). The queue is not safe for concurrent use.
