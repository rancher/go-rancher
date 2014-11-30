package util

func Contains(slice []string, obj string) bool {
	for _, element := range slice {
		if obj == element {
			return true
		}
	}
	return false
}
