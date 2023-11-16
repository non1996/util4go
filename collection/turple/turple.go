package turple

// T2 is a 2-ary tuple.
type T2[T0, T1 any] struct {
	First  T0
	Second T1
}

// Values returns all elements of tuple.
func (t T2[T0, T1]) Values() (T0, T1) {
	return t.First, t.Second
}

// Make2 creates a tuple of 2 elements.
func Make2[T0, T1 any](first T0, second T1) T2[T0, T1] {
	return T2[T0, T1]{first, second}
}

// T3 is a 3-ary tuple.
type T3[T0, T1, T2 any] struct {
	First  T0
	Second T1
	Third  T2
}

// Values returns all elements of tuple.
func (t T3[T0, T1, T2]) Values() (T0, T1, T2) {
	return t.First, t.Second, t.Third
}

// Make3 creates a tuple of 3 elements.
func Make3[T0, T1, T2 any](first T0, second T1, third T2) T3[T0, T1, T2] {
	return T3[T0, T1, T2]{first, second, third}
}

// T4 is a 4-ary tuple.
type T4[T0, T1, T2, T3 any] struct {
	First  T0
	Second T1
	Third  T2
	Fourth T3
}

// Values returns all elements of tuple.
func (t T4[T0, T1, T2, T3]) Values() (T0, T1, T2, T3) {
	return t.First, t.Second, t.Third, t.Fourth
}

// Make4 creates a tuple of 4 elements.
func Make4[T0, T1, T2, T3 any](first T0, second T1, third T2, fourth T3) T4[T0, T1, T2, T3] {
	return T4[T0, T1, T2, T3]{first, second, third, fourth}
}

// T5 is a 5-ary tuple.
type T5[T0, T1, T2, T3, T4 any] struct {
	First  T0
	Second T1
	Third  T2
	Fourth T3
	Fifth  T4
}

// Values returns all elements of tuple.
func (t T5[T0, T1, T2, T3, T4]) Values() (T0, T1, T2, T3, T4) {
	return t.First, t.Second, t.Third, t.Fourth, t.Fifth
}

// Make5 creates a tuple of 5 elements.
func Make5[T0, T1, T2, T3, T4 any](first T0, second T1, third T2, fourth T3, fifth T4) T5[T0, T1, T2, T3, T4] {
	return T5[T0, T1, T2, T3, T4]{first, second, third, fourth, fifth}
}

// T6 is a 6-ary tuple.
type T6[T0, T1, T2, T3, T4, T5 any] struct {
	First  T0
	Second T1
	Third  T2
	Fourth T3
	Fifth  T4
	Sixth  T5
}

// Values returns all elements of tuple.
func (t T6[T0, T1, T2, T3, T4, T5]) Values() (T0, T1, T2, T3, T4, T5) {
	return t.First, t.Second, t.Third, t.Fourth, t.Fifth, t.Sixth
}

// Make6 creates a tuple of 6 elements.
func Make6[T0, T1, T2, T3, T4, T5 any](first T0, second T1, third T2, fourth T3, fifth T4, sixth T5) T6[T0, T1, T2, T3, T4, T5] {
	return T6[T0, T1, T2, T3, T4, T5]{first, second, third, fourth, fifth, sixth}
}

// T7 is a 7-ary tuple.
type T7[T0, T1, T2, T3, T4, T5, T6 any] struct {
	First   T0
	Second  T1
	Third   T2
	Fourth  T3
	Fifth   T4
	Sixth   T5
	Seventh T6
}

// Values returns all elements of tuple.
func (t T7[T0, T1, T2, T3, T4, T5, T6]) Values() (T0, T1, T2, T3, T4, T5, T6) {
	return t.First, t.Second, t.Third, t.Fourth, t.Fifth, t.Sixth, t.Seventh
}

// Make7 creates a tuple of 7 elements.
func Make7[T0, T1, T2, T3, T4, T5, T6 any](first T0, second T1, third T2, fourth T3, fifth T4, sixth T5, seventh T6) T7[T0, T1, T2, T3, T4, T5, T6] {
	return T7[T0, T1, T2, T3, T4, T5, T6]{first, second, third, fourth, fifth, sixth, seventh}
}

// T8 is a 8-ary tuple.
type T8[T0, T1, T2, T3, T4, T5, T6, T7 any] struct {
	First   T0
	Second  T1
	Third   T2
	Fourth  T3
	Fifth   T4
	Sixth   T5
	Seventh T6
	Eighth  T7
}

// Values returns all elements of tuple.
func (t T8[T0, T1, T2, T3, T4, T5, T6, T7]) Values() (T0, T1, T2, T3, T4, T5, T6, T7) {
	return t.First, t.Second, t.Third, t.Fourth, t.Fifth, t.Sixth, t.Seventh, t.Eighth
}

// Make8 creates a tuple of 8 elements.
func Make8[T0, T1, T2, T3, T4, T5, T6, T7 any](first T0, second T1, third T2, fourth T3, fifth T4, sixth T5, seventh T6, eighth T7) T8[T0, T1, T2, T3, T4, T5, T6, T7] {
	return T8[T0, T1, T2, T3, T4, T5, T6, T7]{first, second, third, fourth, fifth, sixth, seventh, eighth}
}

// T9 is a 9-ary tuple.
type T9[T0, T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	First   T0
	Second  T1
	Third   T2
	Fourth  T3
	Fifth   T4
	Sixth   T5
	Seventh T6
	Eighth  T7
	Ninth   T8
}

// Values returns all elements of tuple.
func (t T9[T0, T1, T2, T3, T4, T5, T6, T7, T8]) Values() (T0, T1, T2, T3, T4, T5, T6, T7, T8) {
	return t.First, t.Second, t.Third, t.Fourth, t.Fifth, t.Sixth, t.Seventh, t.Eighth, t.Ninth
}

// Make9 creates a tuple of 9 elements.
func Make9[T0, T1, T2, T3, T4, T5, T6, T7, T8 any](first T0, second T1, third T2, fourth T3, fifth T4, sixth T5, seventh T6, eighth T7, ninth T8) T9[T0, T1, T2, T3, T4, T5, T6, T7, T8] {
	return T9[T0, T1, T2, T3, T4, T5, T6, T7, T8]{first, second, third, fourth, fifth, sixth, seventh, eighth, ninth}
}

// T10 is a 10-ary tuple.
type T10[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	First   T0
	Second  T1
	Third   T2
	Fourth  T3
	Fifth   T4
	Sixth   T5
	Seventh T6
	Eighth  T7
	Ninth   T8
	Tenth   T9
}

// Values returns all elements of tuple.
func (t T10[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9]) Values() (T0, T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return t.First, t.Second, t.Third, t.Fourth, t.Fifth, t.Sixth, t.Seventh, t.Eighth, t.Ninth, t.Tenth
}

// Make10 creates a tuple of 10 elements.
func Make10[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9 any](first T0, second T1, third T2, fourth T3, fifth T4, sixth T5, seventh T6, eighth T7, ninth T8, tenth T9) T10[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return T10[T0, T1, T2, T3, T4, T5, T6, T7, T8, T9]{first, second, third, fourth, fifth, sixth, seventh, eighth, ninth, tenth}
}
