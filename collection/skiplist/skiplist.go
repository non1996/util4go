package skiplist

import (
	"math/rand"

	"github.com/non1996/util4go/constraint"
	"github.com/non1996/util4go/function/value"
)

const (
	maxLevel int     = 32
	pFactor  float64 = 0.25
)

type Node[K constraint.Orderable, V any] struct {
	key     K
	val     V
	Forward []*Node[K, V]
}

type SkipList[K constraint.Orderable, V any] struct {
	Head  *Node[K, V]
	level int
	size  int64
}

func (s *SkipList[K, V]) Keys() []K {
	if s.size == 0 {
		return nil
	}

	res := make([]K, 0, s.size)
	for curr := s.Head; curr != nil; curr = curr.Forward[0] {
		res = append(res, curr.key)
	}

	return res
}

func (s *SkipList[K, V]) Values() []V {
	if s.size == 0 {
		return nil
	}

	res := make([]V, 0, s.size)
	for curr := s.Head; curr != nil; curr = curr.Forward[0] {
		res = append(res, curr.val)
	}

	return res
}

func (s *SkipList[K, V]) Size() int64 {
	return s.size
}

func (s *SkipList[K, V]) IsEmpty() bool {
	return s.size == 0
}

func (s *SkipList[K, V]) Contains(k K) bool {
	curr := s.findNode(k, nil)
	return curr.Forward[0] != nil && curr.Forward[0].key == k
}

func (s *SkipList[K, V]) Add(k K, v V) {
	update := s.updateSlice()

	curr := s.findNode(k, update)

	if curr.Forward[0] != nil && curr.Forward[0].key == k {
		curr.Forward[0].val = v
		return
	}

	lv := s.randomLevel()

	s.level = max(s.level, lv)
	newNode := &Node[K, V]{k, v, make([]*Node[K, V], lv)}
	for i, node := range update[:lv] {
		newNode.Forward[i] = node.Forward[i]
		node.Forward[i] = newNode
	}

	s.size++
}

func (s *SkipList[K, V]) Remove(k K) {
	update := s.updateSlice()

	curr := s.findNode(k, update)

	curr = curr.Forward[0]
	if curr == nil || curr.key != k {
		return
	}

	for i := 0; i < s.level && update[i].Forward[i] == curr; i++ {
		update[i].Forward[i] = curr.Forward[i]
	}

	for s.level > 1 && s.Head.Forward[s.level-1] == nil {
		s.level--
	}

	s.size--
}

func (s *SkipList[K, V]) Clear() {
	s.Head = &Node[K, V]{Forward: make([]*Node[K, V], maxLevel)}
	s.level = 1
	s.size = 0
}

func (s *SkipList[K, V]) Get(k K) V {
	curr := s.findNode(k, nil)
	curr = curr.Forward[0]

	if curr != nil && curr.key == k {
		return curr.val
	}

	return value.Zero[V]()
}

func (s *SkipList[K, V]) GetAndExist(k K) (V, bool) {
	curr := s.findNode(k, nil)
	curr = curr.Forward[0]

	if curr != nil && curr.key == k {
		return curr.val, true
	}

	return value.Zero[V](), false
}

func (s *SkipList[K, V]) GetOrDefault(k K, d V) V {
	curr := s.findNode(k, nil)
	curr = curr.Forward[0]

	if curr != nil && curr.key == k {
		return curr.val
	}

	return d
}

func (s *SkipList[K, V]) Foreach(fn func(K, V)) {
	for curr := s.Head.Forward[0]; curr != nil; curr = curr.Forward[0] {
		fn(curr.key, curr.val)
	}
}

func (s *SkipList[K, V]) Replace(k K, v V) {
	curr := s.findNode(k, nil).Forward[0]
	if curr != nil && curr.key == k {
		curr.val = v
	}
}

func (s *SkipList[K, V]) ReplaceAll(fn func(K, V) V) {
	for curr := s.Head.Forward[0]; curr != nil; curr = curr.Forward[0] {
		curr.val = fn(curr.key, curr.val)
	}
}

func (s *SkipList[K, V]) PutIfAbsent(k K, v V) {
	// TODO 要查找两次，待优化
	if !s.Contains(k) {
		s.Add(k, v)
	}
}

func (s *SkipList[K, V]) ComputeIfAbsent(k K, mapping func(K) (V, bool)) {
	// TODO 要查找两次，待优化
	if !s.Contains(k) {
		newValue, toAdd := mapping(k)
		if toAdd {
			s.Add(k, newValue)
		}
	}
}

func (s *SkipList[K, V]) ComputeIfPresent(k K, mapping func(K) (V, bool)) {
	if s.Contains(k) {
		newValue, toAdd := mapping(k)
		if toAdd {
			s.Add(k, newValue)
		} else {
			s.Remove(k)
		}
	}
}

func (s *SkipList[K, V]) Compute(k K, mapping func(K, V, bool) (V, bool)) (V, bool) {
	oldValue, exist := s.GetAndExist(k)

	newValue, toAdd := mapping(k, oldValue, exist)

	if !toAdd {
		if exist {
			s.Remove(k)
			return newValue, false
		} else {
			return newValue, false
		}
	} else {
		s.Add(k, newValue)
		return newValue, true
	}
}

func (s *SkipList[K, V]) AllMatch(predicate func(K, V) bool) bool {
	for curr := s.Head.Forward[0]; curr != nil; curr = curr.Forward[0] {
		if !predicate(curr.key, curr.val) {
			return false
		}
	}

	return true
}

func (s *SkipList[K, V]) AnyMatch(predicate func(K, V) bool) bool {
	for curr := s.Head.Forward[0]; curr != nil; curr = curr.Forward[0] {
		if predicate(curr.key, curr.val) {
			return true
		}
	}

	return false
}

func (s *SkipList[K, V]) NoneMatch(predicate func(K, V) bool) bool {
	for curr := s.Head.Forward[0]; curr != nil; curr = curr.Forward[0] {
		if predicate(curr.key, curr.val) {
			return false
		}
	}

	return true
}

func (s *SkipList[K, V]) updateSlice() []*Node[K, V] {
	update := make([]*Node[K, V], maxLevel)
	for i := range update {
		update[i] = s.Head
	}
	return update
}

func (s *SkipList[K, V]) findNode(key K, update []*Node[K, V]) *Node[K, V] {
	curr := s.Head

	for l := s.level - 1; l >= 0; l-- {
		for curr.Forward[l] != nil && curr.Forward[l].key < key {
			curr = curr.Forward[l]
		}
		if update != nil {
			update[l] = curr
		}
	}

	curr = curr.Forward[0]

	if curr == nil || curr.key != key {
		return nil
	}

	return curr
}

func (s *SkipList[K, V]) randomLevel() int {
	var lv = 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}
