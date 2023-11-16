package hashmap

import (
	"container/list"

	"github.com/non1996/util4go/function/value"
)

type OrderedMap[K comparable, V any] struct {
	l   *list.List
	idx map[K]*list.Element
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		l:   list.New(),
		idx: map[K]*list.Element{},
	}
}

func (o *OrderedMap[K, V]) Keys() []K {
	return Keys(o.idx)
}

func (o *OrderedMap[K, V]) Values() []V {
	if o.l.Len() == 0 {
		return nil
	}

	values := make([]V, 0, o.l.Len())
	for iter := o.l.Front(); iter != nil; iter = iter.Next() {
		values = append(values, o.getValue(iter))
	}

	return values
}

func (o *OrderedMap[K, V]) Entries() []Entry[K, V] {
	if o.l.Len() == 0 {
		return nil
	}

	entries := make([]Entry[K, V], 0, o.l.Len())
	for iter := o.l.Front(); iter != nil; iter = iter.Next() {
		entries = append(entries, o.getEntry(iter))
	}

	return entries
}

func (o *OrderedMap[K, V]) Size() int64 {
	return int64(o.l.Len())
}

func (o *OrderedMap[K, V]) IsEmpty() bool {
	return o.l.Len() == 0
}

func (o *OrderedMap[K, V]) Contains(key K) bool {
	_, exist := o.idx[key]
	return exist
}

func (o *OrderedMap[K, V]) Add(key K, val V) {
	if elem, exist := o.idx[key]; exist {
		elem.Value = Entry[K, V]{key: key, value: val}
	} else {
		o.idx[key] = o.l.PushBack(Entry[K, V]{key: key, value: val})
	}
}

func (o *OrderedMap[K, V]) AddAll(entires ...Entry[K, V]) {
	for _, e := range entires {
		o.Add(e.key, e.value)
	}
}

func (o *OrderedMap[K, V]) Remove(key K) {
	if elem, exist := o.idx[key]; exist {
		o.l.Remove(elem)
		delete(o.idx, key)
	}
}

func (o *OrderedMap[K, V]) RemoveIf(cond func(K, V) bool) {
	for key, node := range o.idx {
		if cond(key, o.getValue(node)) {
			o.l.Remove(node)
			delete(o.idx, key)
		}
	}
}

func (o *OrderedMap[K, V]) Clear() {
	o.l.Init()
	clear(o.idx)
}

func (o *OrderedMap[K, V]) Get(k K) V {
	if elem, exist := o.idx[k]; exist {
		return o.getValue(elem)
	}
	return value.Zero[V]()
}

func (o *OrderedMap[K, V]) GetAndExist(k K) (V, bool) {
	elem, exist := o.idx[k]
	if exist {
		return o.getValue(elem), true
	}
	return value.Zero[V](), false
}

func (o *OrderedMap[K, V]) GetOrDefault(k K, defaultValue V) V {
	if elem, exist := o.idx[k]; exist {
		return o.getValue(elem)
	}
	return defaultValue
}

func (o *OrderedMap[K, V]) Foreach(fn func(K, V)) {
	for k, node := range o.idx {
		fn(k, o.getValue(node))
	}
}

func (o *OrderedMap[K, V]) Replace(k K, v V) {
	if o.Contains(k) {
		o.Add(k, v)
	}
}

func (o *OrderedMap[K, V]) ReplaceAll(fn func(K, V) V) {
	for key, elem := range o.idx {
		elem.Value = fn(key, o.getValue(elem))
	}
}

func (o *OrderedMap[K, V]) PutIfAbsent(k K, v V) {
	if !o.Contains(k) {
		o.Add(k, v)
	}
}

func (o *OrderedMap[K, V]) ComputeIfAbsent(k K, mapping func(K) (V, bool)) {
	if !o.Contains(k) {
		newValue, toAdd := mapping(k)
		if toAdd {
			o.Add(k, newValue)
		}
	}
}

func (o *OrderedMap[K, V]) ComputeIfPresent(k K, mapping func(K) (V, bool)) {
	if o.Contains(k) {
		newValue, toAdd := mapping(k)
		if toAdd {
			o.Add(k, newValue)
		} else {
			o.Remove(k)
		}
	}
}

func (o *OrderedMap[K, V]) Compute(k K, mapping func(K, V, bool) (V, bool)) (V, bool) {
	oldValue := o.Get(k)
	exist := o.Contains(k)

	newValue, toAdd := mapping(k, oldValue, exist)

	if !toAdd {
		if exist {
			o.Remove(k)
			return newValue, false
		} else {
			return newValue, false
		}
	} else {
		o.Add(k, newValue)
		return newValue, true
	}
}

func (o *OrderedMap[K, V]) AllMatch(predicate func(K, V) bool) bool {
	for k, node := range o.idx {
		if !predicate(k, o.getValue(node)) {
			return false
		}
	}

	return true
}

func (o *OrderedMap[K, V]) AnyMatch(predicate func(K, V) bool) bool {
	for k, node := range o.idx {
		if predicate(k, o.getValue(node)) {
			return true
		}
	}

	return false
}

func (o *OrderedMap[K, V]) NoneMatch(predicate func(K, V) bool) bool {
	for k, node := range o.idx {
		if predicate(k, o.getValue(node)) {
			return false
		}
	}

	return true
}

func (o *OrderedMap[K, V]) getValue(node *list.Element) V {
	return node.Value.(Entry[K, V]).value
}

func (o *OrderedMap[K, V]) getEntry(node *list.Element) Entry[K, V] {
	return node.Value.(Entry[K, V])
}
