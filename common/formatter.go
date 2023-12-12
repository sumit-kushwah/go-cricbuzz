package common

func FormatInNLength(s string, length int) string {
	output := []rune{}
	for _, ch := range s {
		output = append(output, ch)
	}
	if len(s) >= length {
		output = output[:length]
		output[length-1] = '.'
		output[length-2] = '.'
		return string(output)
	}
	for i := 0; i < length-len(s); i++ {
		output = append(output, ' ')
	}

	return string(output)
}
