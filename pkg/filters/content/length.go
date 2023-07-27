package content

import "strings"

func LineLength(min int, max int, ignoreWhitespaces bool) func(content []string) (result []string) {
	return func(content []string) (result []string) {
		for _, line := range content {
			length := len(line)
			if ignoreWhitespaces {
				length = len(strings.ReplaceAll(line, " ", ""))
			}
			if length >= min && length <= max {
				result = append(result, line)
			}
		}
		return result
	}
}
