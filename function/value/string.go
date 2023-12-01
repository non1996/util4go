package value

func StringEmpty(s string) bool {
	return len(s) == 0
}

func StringNotEmpty(s string) bool {
	return len(s) > 0
}
