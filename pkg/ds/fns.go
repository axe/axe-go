package ds

func ToSlice[V any](value any) []V {
	if slice, ok := value.([]V); ok {
		return slice
	}
	if indexed, ok := value.(Indexed[V]); ok {
		values := make([]V, indexed.Len())
		for i := range values {
			values[i] = indexed.At(i)
		}
		return values
	}
	if iterable, ok := value.(Iterable[V]); ok {
		value = iterable.Iterator()
	}
	if iterator, ok := value.(Iterator[V]); ok {
		values := make([]V, 0, 32)
		for iterator.HasNext() {
			values = append(values, *iterator.Next())
		}
		return values
	}
	return nil
}
