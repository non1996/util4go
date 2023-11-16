package slice

import (
	"math/rand"
	"slices"

	"github.com/non1996/util4go/function"
	"github.com/non1996/util4go/function/value"
	"github.com/non1996/util4go/optional"
)

func NilToEmpty[T any](s []T) []T {
	if s != nil {
		return s
	}
	return make([]T, 0)
}

func Concat[T any](s1, s2 []T) []T {
	var res = make([]T, 0, len(s1)+len(s2))
	res = append(res, s1...)
	res = append(res, s2...)
	return res
}

func ConcatM[T any](s1 []T, ss ...[]T) []T {
	if len(ss) == 0 {
		return s1
	}

	var res []T
	res = append(res, s1...)
	for _, s2 := range ss {
		res = append(res, s2...)
	}
	return res
}

func GetFirst[T any](s []T) T {
	return s[0]
}

func GetFirstOr[T any](s []T, d T) T {
	if len(s) == 0 {
		return d
	}
	return s[0]
}

func GetLast[T any](s []T) T {
	return s[len(s)-1]
}

func GetLastOr[T any](s []T, d T) T {
	if len(s) == 0 {
		return d
	}
	return s[len(s)-1]
}

func Copy[T any](s []T) []T {
	return slices.Clone(s)
}

func Shuffle[T any](s []T) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func CollectToMap[T any, K comparable, V any](
	s []T,
	keyMapper function.Function[T, K],
	valMapper function.Function[T, V],
) map[K]V {
	res := map[K]V{}
	for idx := range s {
		res[keyMapper(s[idx])] = valMapper(s[idx])
	}
	return res
}

func Group[T any, K comparable, V any](
	s []T,
	keyMapper function.Function[T, K],
	valMapper function.Function[T, V],
) map[K][]V {
	return Aggregate(s, keyMapper, func(t T, v []V) []V { return append(v, valMapper(t)) })
}

func Reduce[T any](s []T, operation function.BiOperation[T]) optional.Optional[T] {
	return ReduceWithIdentity(value.Zero[T](), s, operation)
}

func ReduceWithIdentity[T any](identity T, s []T, operation function.BiOperation[T]) optional.Optional[T] {
	if len(s) == 0 {
		return optional.Empty[T]()
	}
	for idx := range s {
		identity = operation(identity, s[idx])
	}
	return optional.Of(identity)
}

func Map[T1 any, T2 any](s []T1, mapper function.Function[T1, T2]) []T2 {
	if len(s) == 0 {
		return nil
	}
	res := make([]T2, 0, len(s))
	for idx := range s {
		res = append(res, mapper(s[idx]))
	}
	return res
}

func MapWithError[T1, T2 any](list []T1, mapper func(T1) (T2, error)) (res []T2, err error) {
	if len(list) == 0 {
		return nil, nil
	}
	res = make([]T2, 0, len(list))
	for _, i := range list {
		r, err := mapper(i)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return
}

func FlatMap[T1 any, T2 any](s []T1, mapper function.Function[T1, []T2]) []T2 {
	if len(s) == 0 {
		return nil
	}
	res := make([]T2, 0)
	for idx := range s {
		res = append(res, mapper(s[idx])...)
	}
	return res
}

func Aggregate[T any, K comparable, V any, S ~[]T](s S, keyFn func(T) K, groupFn func(T, V) V) map[K]V {
	m := make(map[K]V)
	for i := range s {
		k := keyFn(s[i])
		m[k] = groupFn(s[i], m[k])
	}
	return m
}
