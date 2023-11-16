package hashmap

import (
	"maps"

	"github.com/non1996/util4go/function"
)

func Keys[K comparable, V any](m map[K]V) (keys []K) {
	if len(m) == 0 {
		return nil
	}
	keys = make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) (values []V) {
	if len(m) == 0 {
		return nil
	}
	values = make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Entries[K comparable, V any](m map[K]V) (entries []Entry[K, V]) {
	if len(m) == 0 {
		return nil
	}
	entries = make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, NewEntry(k, v))
	}
	return entries
}

func Size[K comparable, V any](m map[K]V) int64 {
	return int64(len(m))
}

func IsEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) == 0
}

func Contains[K comparable, V any](m map[K]V, k K) bool {
	_, exist := m[k]
	return exist
}

func Add[K comparable, V any](m map[K]V, key K, value V) {
	m[key] = value
}

func AddAll[K comparable, V any](m, other map[K]V) {
	for k, v := range other {
		m[k] = v
	}
}

func Remove[K comparable, V any](m map[K]V, key K) {
	delete(m, key)
}

func RemoveIf[K comparable, V any](m map[K]V, cond func(K, V) bool) {
	for k, v := range m {
		if cond(k, v) {
			delete(m, k)
		}
	}
}

func Clear[K comparable, V any](m map[K]V) {
	clear(m)
}

func Get[K comparable, V any](m map[K]V, key K) V {
	return m[key]
}

func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	v, exist := m[key]
	if exist {
		return v
	}
	return defaultValue
}

func ForEach[K comparable, V any](m map[K]V, fn func(K, V)) {
	for k, v := range m {
		fn(k, v)
	}
}

func Replace[K comparable, V any](m map[K]V, k K, v V) {
	if Contains(m, k) {
		m[k] = v
	}
}

func ReplaceAll[K comparable, V any](m map[K]V, fn func(K, V) V) {
	for key := range m {
		m[key] = fn(key, m[key])
	}
}

func PutIfAbsent[K comparable, V any](m map[K]V, k K, v V) {
	if !Contains(m, k) {
		m[k] = v
	}
}

func ComputeIfAbsent[K comparable, V any](m map[K]V, k K, mapping func(K) (V, bool)) {
	if !Contains(m, k) {
		newValue, toAdd := mapping(k)
		if toAdd {
			m[k] = newValue
		}
	}
}

func ComputeIfPresent[K comparable, V any](m map[K]V, k K, mapping func(K) (V, bool)) {
	if Contains(m, k) {
		newValue, toAdd := mapping(k)
		if toAdd {
			m[k] = newValue
		} else {
			delete(m, k)
		}
	}
}

func Compute[K comparable, V any](m map[K]V, k K, mapping func(K, V, bool) (V, bool)) (V, bool) {
	oldValue, exist := m[k]
	newValue, toAdd := mapping(k, oldValue, exist)

	if !toAdd {
		if exist {
			delete(m, k)
			return newValue, false
		} else {
			return newValue, false
		}
	} else {
		m[k] = newValue
		return newValue, true
	}
}

func AllMatch[K comparable, V any](m map[K]V, predicate func(K, V) bool) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}

	return true
}

func AnyMatch[K comparable, V any](m map[K]V, predicate func(K, V) bool) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}

	return false
}

func NoneMatch[K comparable, V any](m map[K]V, predicate func(K, V) bool) bool {
	for k, v := range m {
		if predicate(k, v) {
			return false
		}
	}

	return true
}

func Copy[K comparable, V any](m map[K]V) map[K]V {
	return maps.Clone(m)
}

func Merge[K comparable, V any](m1, m2 map[K]V) map[K]V {
	newMap := make(map[K]V, len(m1)+len(m2))

	maps.Copy(newMap, m1)
	maps.Copy(newMap, m2)

	return newMap
}

func MapToSlice[K comparable, V any, T any](
	m map[K]V,
	mapper function.BiFunction[K, V, T],
) []T {
	res := make([]T, 0, len(m))
	for k, v := range m {
		res = append(res, mapper(k, v))
	}
	return res
}

func Mapping[K comparable, V any, K2 comparable, V2 any](
	m map[K]V,
	mapper func(K, V) (K2, V2),
) map[K2]V2 {
	if m == nil {
		return nil
	}

	newMap := make(map[K2]V2, len(m))
	for k, v := range m {
		k2, v2 := mapper(k, v)
		newMap[k2] = v2
	}

	return newMap
}
