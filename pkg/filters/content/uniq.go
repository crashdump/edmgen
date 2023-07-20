package content

func Uniq(content []string) (result []string) {
	keys := make(map[string]bool)
	for _, entry := range content {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}
