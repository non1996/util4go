package set

import (
	"encoding/json"

	"github.com/non1996/util4go/collection/hashmap"
)

type OrderedSet[E Elem] struct {
	m *hashmap.OrderedMap[E, placeholder]
}

func NewOrderedSet[E Elem](list ...E) *OrderedSet[E] {
	set := &OrderedSet[E]{m: hashmap.NewOrderedMap[E, placeholder]()}
	set.AddAll(list...)
	return set
}

func (o *OrderedSet[E]) Size() int64 {
	return o.m.Size()
}

func (o *OrderedSet[E]) IsEmpty() bool {
	return o.m.IsEmpty()
}

func (o *OrderedSet[E]) Contains(e E) bool {
	return o.m.Contains(e)
}

func (o *OrderedSet[E]) Add(e E) {
	o.m.Add(e, pl)
}

func (o *OrderedSet[E]) AddAll(es ...E) {
	if len(es) == 0 {
		return
	}
	for _, e := range es {
		o.Add(e)
	}
}

func (o *OrderedSet[E]) Remove(e E) {
	o.m.Remove(e)
}

func (o *OrderedSet[E]) RemoveIf(cond func(E) bool) {
	o.m.RemoveIf(func(e E, _ placeholder) bool { return cond(e) })
}

func (o *OrderedSet[E]) Clear() {
	o.m.Clear()
}

func (o *OrderedSet[E]) ToSlice() []E {
	return o.m.Keys()
}

func (o *OrderedSet[E]) Foreach(f func(E)) {
	o.m.Foreach(func(e E, _ placeholder) { f(e) })
}

func (o *OrderedSet[E]) AllMatch(predicate func(E) bool) bool {
	return o.m.AllMatch(func(e E, _ placeholder) bool { return predicate(e) })
}

func (o *OrderedSet[E]) AnyMatch(predicate func(E) bool) bool {
	return o.m.AnyMatch(func(e E, _ placeholder) bool { return predicate(e) })
}

func (o *OrderedSet[E]) NoneMatch(predicate func(E) bool) bool {
	return o.m.NoneMatch(func(e E, _ placeholder) bool { return predicate(e) })
}

func (o *OrderedSet[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToSlice())
}

func (o *OrderedSet[E]) UnmarshalJSON(b []byte) error {
	var res []E
	err := json.Unmarshal(b, &res)
	if err != nil {
		return err
	}
	*o = *NewOrderedSet[E](res...)
	return nil
}
