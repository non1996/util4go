package set

import (
	"encoding/json"
)

type HashSet[E Elem] struct {
	m map[E]placeholder
}

func NewHashSet[E Elem](list ...E) *HashSet[E] {
	set := &HashSet[E]{m: map[E]placeholder{}}
	set.AddAll(list...)
	return set
}

func (s *HashSet[E]) Size() int64 {
	return int64(len(s.m))
}

func (s *HashSet[E]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *HashSet[E]) Contains(v E) bool {
	_, exist := s.m[v]
	return exist
}

func (s *HashSet[E]) Add(i E) {
	s.m[i] = placeholder{}
}

func (s *HashSet[E]) AddAll(es ...E) {
	if len(es) == 0 {
		return
	}
	for _, e := range es {
		s.Add(e)
	}
}

func (s *HashSet[E]) Remove(i E) {
	delete(s.m, i)
}

func (s *HashSet[E]) RemoveIf(cond func(E) bool) {
	for elem := range s.m {
		if cond(elem) {
			s.Remove(elem)
		}
	}
}

func (s *HashSet[E]) Clear() {
	clear(s.m)
}

func (s *HashSet[E]) ToSlice() []E {
	if s.Size() == 0 {
		return nil
	}
	list := make([]E, 0, s.Size())
	for elem := range s.m {
		list = append(list, elem)
	}
	return list
}

func (s *HashSet[E]) Foreach(action func(E)) {
	for elem := range s.m {
		action(elem)
	}
}

func (s *HashSet[E]) AllMatch(predicate func(E) bool) bool {
	if s.Size() == 0 {
		return true
	}

	for elem := range s.m {
		if !predicate(elem) {
			return false
		}
	}

	return true
}

func (s *HashSet[E]) AnyMatch(predicate func(E) bool) bool {
	if s.Size() == 0 {
		return false
	}

	for elem := range s.m {
		if predicate(elem) {
			return true
		}
	}

	return false
}

func (s *HashSet[E]) NoneMatch(predicate func(E) bool) bool {
	if s.Size() == 0 {
		return true
	}

	for elem := range s.m {
		if predicate(elem) {
			return false
		}
	}

	return true
}

func (s *HashSet[E]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToSlice())
}

func (s *HashSet[E]) UnmarshalJSON(b []byte) error {
	var res []E
	err := json.Unmarshal(b, &res)
	if err != nil {
		return err
	}
	*s = *NewHashSet[E](res...)
	return nil
}
