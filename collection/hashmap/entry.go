package hashmap

type Entry[K comparable, V any] struct {
	key   K
	value V
}

func NewEntry[K comparable, V any](k K, v V) Entry[K, V] {
	return Entry[K, V]{key: k, value: v}
}

func (e Entry[K, V]) Key() K {
	return e.key
}

func (e Entry[K, V]) Value() V {
	return e.value
}
