package main

// Assign function works similiar to Objectassign function in javascript
// this function join map[string]interface{} overwriting values from left to right
func Assign(args ...map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})
	for _, item := range args {
		for key, value := range item {
			output[key] = value
		}
	}

	return output
}
