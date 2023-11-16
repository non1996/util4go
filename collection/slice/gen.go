package slice

func Empty[T any]() []T {
	return make([]T, 0)
}

func Nil[T any]() []T {
	return ([]T)(nil)
}

func Singleton[T any](t T) []T {
	s := make([]T, 1)
	s[0] = t
	return s
}

func Repeat[T any](v T, count int) []T {
	if count <= 0 {
		return nil
	}

	res := make([]T, count)
	for i := 0; i < count; i++ {
		res[i] = v
	}
	return res
}

func Gen[T any](generator func() (T, bool)) (res []T) {
	for {
		v, hasMore := generator()
		if !hasMore {
			return
		}
		res = append(res, v)
	}
}
