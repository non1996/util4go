package priorityqueue

import (
	"github.com/non1996/util4go/constraint"
)

func Max[V constraint.Orderable](v1 V, v2 V) bool {
	return v1 > v2
}

func Min[V constraint.Orderable](v1 V, v2 V) bool {
	return v1 < v2
}
