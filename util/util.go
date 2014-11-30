package util

import (
	"errors"
)

func Contains(slice []string, obj string) (bool, error) {
	if slice == nil {
		return false, errors.New("nil input slice")
	}
	for _, element := range slice {
		if obj == element {
			return true, nil
		}
	}
	return false, nil
}
