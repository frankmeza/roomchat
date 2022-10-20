package utils

// IsStringInSlice("apel", []string{"apel", "jeruk"}) -> true
// IsStringInSlice("pisang", []string{"apel", "jeruk"}) -> false

func IsStringInSlice(inputString string, slice []string) bool {
	for _, item := range slice {
		if item == inputString {
			return true
		}
	}

	return false
}
