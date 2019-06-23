package main

import (
	"errors"
	"fmt"
)

func Find(slice []string, query string) (int, error) {
	for index, pointer := range slice {
		if query == pointer {
			return index, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("'%v' was not found", query))
}
