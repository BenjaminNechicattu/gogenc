package strings

func Reverse(original string) (revsered string) {
	for i := len(original); i > 0; i-- {
		revsered = revsered + string(original[i-1])
	}
	return
}
