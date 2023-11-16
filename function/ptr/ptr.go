package ptr

import (
	"github.com/non1996/util4go/function/value"
)

// IsNil 判断是否是nil指针
func IsNil[T any](p *T) bool {
	return p == nil
}

// NonNil 判断是否是nil指针
func NonNil[T any](p *T) bool {
	return p != nil
}

// Ref 取地址，返回指针
func Ref[T any](v T) *T {
	return &v
}

// Indirect 函数接受一个指针，返回指针指向的值（解引用），对于空指针回返回指针指向类型的默认值
func Indirect[T any](p *T) T {
	if p != nil {
		return *p
	}
	return value.Zero[T]()
}

// IndirectOr 和 Indirect 类似，但允许用户自定义遇到空指针时返回的默认值。
func IndirectOr[T any](p *T, d T) T {
	if p != nil {
		return *p
	}
	return d
}

// Copy 只拷贝对象本身，不拷贝其指针成员指向的其他变量
func Copy[T any](p *T) *T {
	if p == nil {
		return nil
	}

	c := *p
	return &c
}
