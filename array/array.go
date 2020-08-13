package array

// ContainString determine whether the string element is in the array
func ContainString(array []string, item string) bool {
	if array == nil || len(array) == 0 {
		return false
	}
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}
