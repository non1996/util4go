package ternary

import (
	"github.com/non1996/util4go/function"
)

// Ternary 三元表达式
func Ternary[T any](cond bool, v1, v2 T) T {
	if cond {
		return v1
	}
	return v2
}

// Lazy 使用函数闭包延迟求值的三元表达式
func Lazy[T any](cond bool, supplier1, supplier2 function.Supplier[T]) T {
	if cond {
		return supplier1()
	}
	return supplier2()
}
