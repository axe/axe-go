package ds

import "sort"

type Sortable[T any] interface {
	Less(b T) bool
}

type SortableList[T Sortable[T]] struct {
	List[T]
}

func (l *SortableList[T]) Sort() {
	sort.Sort(l)
}

func (l *SortableList[T]) Less(i, j int) bool {
	return l.Items[i].Less(l.Items[j])
}
