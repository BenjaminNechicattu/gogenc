package compare

func Max[T float32 | float64 | int | int8 | int16 | int32 | int64 | string](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T float32 | float64 | int | int8 | int16 | int32 | int64 | string](a, b T) T {
	if a > b {
		return a
	}
	return b
}
