package content

import "strings"

func IgnoreLine(str string) func(content []string) (result []string) {
	return func(content []string) (result []string) {
		for _, line := range content {
			if strings.Contains(line, str) {
				continue
			}
			result = append(result, line)
		}
		return result
	}
}
