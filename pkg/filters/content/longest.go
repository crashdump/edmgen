package content

func LongestLine(content []string) (result []string) {
	var lineLengh int
	for _, line := range content {
		if len(line) > lineLengh {
			lineLengh = len(line)
			result = []string{line}
		}
	}
	return result
}
