package optional

import (
	"github.com/non1996/util4go/function"
	"github.com/non1996/util4go/function/ternary"
)

func Of[T any](v T) Optional[T] {
	return Optional[T]{
		value: v,
		valid: true,
	}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

type Optional[T any] struct {
	value T
	valid bool
}

func (o Optional[T]) Get() T {
	return o.value
}

func (o Optional[T]) GetOr(d T) T {
	return ternary.Ternary(o.valid, o.value, d)
}

func (o Optional[T]) IsPresent() bool {
	return o.valid
}

func (o Optional[T]) IfPresent(onValid function.Consumer[T]) {
	if o.valid {
		onValid(o.value)
	}
}

func (o Optional[T]) IfPresentElse(onValid function.Consumer[T], onElse func()) {
	if o.valid {
		onValid(o.value)
	} else {
		onElse()
	}
}
