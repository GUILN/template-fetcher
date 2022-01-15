package helpers

func SliceContains(slice []interface{}, filterFunction func(interface{}) bool) bool {
	for _, element := range slice {
		if filterFunction(element) {
			return true
		}
	}
	return false
}

// SliceReduce
// Params: slice []interface{}; func(previousElement, currentElement)
func SliceReduce(slice []interface{}, reducerFunc func(interface{}, interface{}) interface{}) interface{} {
	var curr, prev interface{}
	for i := range slice {
		curr = slice[i]
		prev = reducerFunc(prev, curr)
	}

	return prev
}
