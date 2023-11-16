package set

// IsSubset 判断 src 是不是 dst 的 subset
func IsSubset[E Elem](dst, src Set[E]) bool {
	if src.IsEmpty() {
		return true
	}
	if dst.IsEmpty() {
		return false
	}

	if src.Size() > dst.Size() {
		return false
	}

	return src.AllMatch(dst.Contains)
}

// IsSuperSet 判断 src 是不是 dst 的 superset
func IsSuperSet[E Elem](dst, src Set[E]) bool {
	return IsSubset[E](src, dst)
}

// Equal 判断两个 set 是否相等
func Equal[E Elem](one, another Set[E]) bool {
	if one.IsEmpty() && another.IsEmpty() {
		return true
	}
	if one.IsEmpty() || another.IsEmpty() {
		return false
	}

	if one.Size() != another.Size() {
		return false
	}

	return one.AllMatch(another.Contains)
}

// Union 并集
func Union[E Elem, S1, S2, S3 Set[E]](one S1, another S2, unioned S3) S3 {
	one.Foreach(unioned.Add)
	another.Foreach(unioned.Add)
	return unioned
}

// Intersect 交集
func Intersect[E Elem, S1, S2, S3 Set[E]](one S1, another S2, intersected S3) S3 {
	if one.IsEmpty() || another.IsEmpty() {
		return intersected
	}

	// loop over smaller Set
	var small Set[E] = one
	var large Set[E] = another

	if one.Size() > another.Size() {
		small, large = large, small
	}

	small.Foreach(func(e E) {
		if large.Contains(e) {
			intersected.Add(e)
		}
	})

	return intersected
}

// Difference 计算 src 和 dst 的差集
func Difference[E Elem, S1, S2, S3 Set[E]](dst S1, src S2, diff S3) S3 {
	if src.IsEmpty() {
		return diff
	}

	src.Foreach(func(e E) {
		if !dst.Contains(e) {
			diff.Add(e)
		}
	})

	return diff
}
