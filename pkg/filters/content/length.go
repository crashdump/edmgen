package content

func LineLength(min int, max int) func(content []string) (result []string) {
	return func(content []string) (result []string) {
		for _, line := range content {
			if len(line) >= min && len(line) <= max {
				result = append(result, line)
			}
		}
		return result
	}
}
