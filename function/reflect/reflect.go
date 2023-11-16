package reflect

import (
	"reflect"
	"unsafe"
)

// TypeOf 返回一个类型的 reflect.Type
func TypeOf[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

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
