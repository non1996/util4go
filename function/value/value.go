package value

import (
	"unsafe"

	"github.com/non1996/util4go/constraint"
)

type xface struct {
	x    uintptr
	data unsafe.Pointer
}

func IsNil(v any) bool {
	return (*xface)(unsafe.Pointer(&v)).data == nil
}

func NonNil(v any) bool {
	return !IsNil(v)
}

func Noop[T any](t T) T {
	return t
}

func Void[T any](_ T) {

}

// Zero 函数返回一个类型的零值
func Zero[T any]() T {
	return *new(T)
}

func IsZero[T comparable](v T) bool {
	return v == Zero[T]()
}

func IfZero[T comparable](v T, d T) T {
	if IsZero(v) {
		return d
	}
	return v
}

func CastToAny[T any](v T) any {
	return v
}

func CastFromAny[T any](v any) T {
	return v.(T)
}

func Sum[T constraint.Addable](v1, v2 T) T {
	return v1 + v2
}

func Sub[T constraint.Number | constraint.Complex](v1, v2 T) T {
	return v1 - v2
}

// Equal 判断两个值是否相等
func Equal[T comparable](v1, v2 T) bool {
	return v1 == v2
}

// Greater 判断v1是否大于v2
func Greater[T constraint.Orderable](v1, v2 T) bool {
	return v1 > v2
}

// GreaterOrEqual 判断v1是否大于等于v2
func GreaterOrEqual[T constraint.Orderable](v1, v2 T) bool {
	return v1 >= v2
}

// Less 判断v1是否小于v2
func Less[T constraint.Orderable](v1, v2 T) bool {
	return v1 < v2
}

// LessOrEqual 判断v1是否小于等于v2
func LessOrEqual[T constraint.Orderable](v1, v2 T) bool {
	return v1 <= v2
}

// Max 返回最大值
// Deprecated 已废弃，直接使用内置函数
func Max[T constraint.Orderable](v1 T, v2 ...T) T {
	return max(v1, v2...)
}

// Min 返回最小值
// Deprecated 已废弃，直接使用内置函数
func Min[T constraint.Orderable](v1 T, v2 ...T) T {
	return min(v1, v2...)
}
