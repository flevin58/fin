package tools

func RemoveAtIndex[T comparable](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func FindIndexOf[T comparable](slice []T, what T) (int, bool) {
	for index, elem := range slice {
		if elem == what {
			return index, true
		}
	}
	return 0, false
}
