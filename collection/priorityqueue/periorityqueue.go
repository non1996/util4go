package priorityqueue

import (
	"container/heap"

	"github.com/non1996/util4go/function"
)

type PriorityQueue[V any] struct {
	impl[V]
}

func New[V any](comparator function.Comparator[V]) *PriorityQueue[V] {
	return &PriorityQueue[V]{
		impl: impl[V]{
			comparator: comparator,
		},
	}
}

func (p *PriorityQueue[V]) Clear() {
	p.impl.l = nil
}

func (p *PriorityQueue[V]) Add(v V) {
	heap.Push(&p.impl, v)
}

func (p *PriorityQueue[V]) Peek() V {
	return p.impl.l[0]
}

func (p *PriorityQueue[V]) Poll() V {
	return heap.Pop(&p.impl).(V)
}

func (p *PriorityQueue[V]) Size() int64 {
	return int64(len(p.impl.l))
}

func (p *PriorityQueue[V]) ToSlice() []V {
	return p.impl.l
}

type impl[V any] struct {
	l          []V
	comparator function.Comparator[V]
}

func (i *impl[V]) Len() int {
	return len(i.l)
}

func (i *impl[V]) Less(idx1, idx2 int) bool {
	return i.comparator(i.l[idx1], i.l[idx2])
}

func (i *impl[V]) Swap(idx1, idx2 int) {
	i.l[idx1], i.l[idx2] = i.l[idx2], i.l[idx1]
}

func (i *impl[V]) Push(v any) {
	i.l = append(i.l, v.(V))
}

func (i *impl[V]) Pop() any {
	v := i.l[len(i.l)-1]
	i.l = i.l[:len(i.l)-1]
	return v
}
