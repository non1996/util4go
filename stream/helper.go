package stream

import (
	"github.com/non1996/util4go/function"
)

func Slice[T any](s []T) Stream[T] {
	return newPipeline[T](&sliceIterator[T]{s: s})
}

func New[T any](iter Iterator[T]) Stream[T] {
	return newPipeline[T](iter)
}

func Map[T1, T2 any](src Stream[T1], mapper func(T1) T2) Stream[T2] {
	pipe := src.(*pipeline[T1])
	iter := &mappingIterator[T1, T2]{
		upstream: pipe.iter,
		stages: func(in T1) (out T2, nextAction bool, nextElem bool) {
			in, nextAction, nextElem = pipe.stages(in)
			if !nextAction {
				return
			}
			return mapper(in), true, nextElem
		},
	}
	return newPipeline[T2](iter)
}

func CollectToMap[T any, K comparable, V any](
	s Stream[T],
	keyMapper function.Function[T, K],
	valMapper function.Function[T, V],
) map[K]V {
	res := map[K]V{}
	s.Foreach(func(v T) {
		res[keyMapper(v)] = valMapper(v)
	})
	return res
}
