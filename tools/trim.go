package tools

func TrimString(s string, length int) string {
	if len(s) <= length-1 {
		return s
	}

	max := length + 1
	return "..." + s[max:]
}
