package main

// Contains finds an specific string into an string array
// and then returns true or false depends on the result
func Contains(slice []string, query string) bool {
	for _, pointer := range slice {
		if query == pointer {
			return true
		}
	}
	return false
}
