package set

type Elem interface {
	comparable
}

type Set[E Elem] interface {
	Size() int64
	IsEmpty() bool
	Contains(E) bool
	Add(E)
	AddAll(...E)
	Remove(E)
	RemoveIf(func(E) bool)
	Clear()
	ToSlice() []E
	Foreach(func(E))
	AllMatch(func(E) bool) bool
	AnyMatch(func(E) bool) bool
	NoneMatch(func(E) bool) bool
}

type placeholder struct{}

var pl = placeholder{}
