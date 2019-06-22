package main

func Contains(slice []string, query string) bool {
	for _, pointer := range slice {
		if query == pointer {
			return true
		}
	}
	return false
}
