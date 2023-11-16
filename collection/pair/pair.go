package pair

type Pair[T1, T2 any] struct {
	first  T1
	second T2
}

func NewPair[T1, T2 any](v1 T1, v2 T2) Pair[T1, T2] {
	return Pair[T1, T2]{
		first:  v1,
		second: v2,
	}
}

func (p Pair[T1, T2]) First() T1 {
	return p.first
}

func (p Pair[T1, T2]) Second() T2 {
	return p.second
}

func (p Pair[T1, T2]) Copy() Pair[T1, T2] {
	return Pair[T1, T2]{first: p.first, second: p.second}
}
