package util

func makeSnakeCase(str ...string) string {
	if len(str) == 0 {
		return ""
	}

	result := str[0]
	if len(str) == 1 {
		return result
	}
	for i := 1; i < len(str); i++ {
		result += "_" + str[i]
	}

	return result
}

func makeStringList(str ...string) []string {
	strings := make([]string, len(str))

	for i, s := range str {
		strings[i] = s
	}

	return strings
}
