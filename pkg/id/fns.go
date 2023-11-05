package id

func resize[V any](slice []V, size int) []V {
	if size < len(slice) {
		return slice[:size]
	} else {
		return append(slice, make([]V, size-len(slice))...)
	}
}

func removeAt[V any](slice []V, index int) []V {
	return append(slice[:index], slice[index+1:]...)
}

func moveEndTo[V any](slice []V, index int) []V {
	end := len(slice) - 1
	slice[index] = slice[end]
	return slice[:end]
}
