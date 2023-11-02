package id

import (
	"golang.org/x/exp/constraints"
)

type Identifier int32

func (i Identifier) Exists() bool   { return i >= 0 }
func (i Identifier) Empty() bool    { return i <= 0 }
func (i Identifier) String() string { return getString(i) }

const (
	pagePower = 12
	pageSize  = 1 << pagePower
	pageHash  = pageSize - 1
)

var (
	mapping     = make(map[string]int32, 256)
	storageInfo = make([]string, 0, 256)
	storageData = make([]string, 0)
	storageEnd  = 0
)

func Get(s string) Identifier {
	if id, exists := mapping[s]; exists {
		return Identifier(id)
	} else {
		length := len(s)
		// nextEnd := storageEnd + length
		// pageEnd := len(storageData) * pageSize
		// if nextEnd > pageEnd {
		// 	storageEnd = pageEnd
		// 	actualPageSize := pageSize
		// 	if length > actualPageSize {
		// 		actualPageSize = length
		// 	}
		// 	storageData = append(storageData, strings.Repeat(" ", actualPageSize))
		// }
		nextId := int32(len(storageInfo))

		// pageIndex := storageEnd << pagePower
		// pageOffset := storageEnd & pageHash
		// page := storageData[pageIndex]
		// begin := pageOffset
		// end := begin + length
		// for i, c := range []byte(s) {
		// page[begin+i] = c
		// }
		// storageInfo = append(storageInfo, page[begin:end])

		storageInfo = append(storageInfo, s)
		storageEnd += length
		mapping[s] = nextId
		return Identifier(nextId)
	}
}

func init() {
	Get("")
}

func Maybe(s string) Identifier {
	id, exists := mapping[s]
	if exists {
		return Identifier(id)
	}
	return Identifier(-1)
}

func getString(i Identifier) string {
	if i >= 0 && int(i) < len(storageData) {
		return storageInfo[i]
	}
	return ""
}

type Option[O any] func(optionable O)

func processOptions[O any](optionable O, options []Option[O]) {
	for _, option := range options {
		option(optionable)
	}
}

type HasResizeBuffer interface{ setResizeBuffer(int) }

func WithResizeBuffer(resizeBuffer int) Option[HasResizeBuffer] {
	return func(optionable HasResizeBuffer) {
		optionable.setResizeBuffer(resizeBuffer)
	}
}

type HasCapacity interface{ setCapacity(int) }

func WithCapacity(capacity int) Option[HasCapacity] {
	return func(optionable HasCapacity) {
		optionable.setCapacity(capacity)
	}
}

type HasArea[A any] interface{ setArea(*A) }

func WithArea[A any](area *A) Option[HasArea[A]] {
	return func(optionable HasArea[A]) {
		optionable.setArea(area)
	}
}

type Area[From constraints.Unsigned, To constraints.Unsigned] struct {
	tos          []To
	next         To
	resizeBuffer int
}

var _ HasCapacity = &Area[uint, uint]{}
var _ HasResizeBuffer = &Area[uint, uint]{}

func NewArea[From constraints.Unsigned, To constraints.Unsigned](opts ...Option[*Area[From, To]]) *Area[From, To] {
	a := &Area[From, To]{
		tos: make([]To, 0),
	}
	processOptions(a, opts)
	return a
}

func (m *Area[From, To]) setResizeBuffer(resizeBuffer int) {
	m.resizeBuffer = resizeBuffer
}

func (m *Area[From, To]) setCapacity(capacity int) {
	m.tos = make([]To, 0, capacity)
}

func (a *Area[From, To]) Translate(from From) To {
	if a == nil {
		return To(from)
	}
	if int(from) >= len(a.tos) {
		nextSize := int(from) + a.resizeBuffer + 1
		a.tos = resize(a.tos, nextSize)
	}
	if a.tos[from] == 0 {
		a.next++
		a.tos[from] = a.next
	}
	return a.tos[from] - 1
}

func (a Area[From, To]) Has(from From) bool {
	return int(from) < len(a.tos) && a.tos[from] != 0
}

func (a *Area[From, To]) Peek(from From) int {
	if a != nil && int(from) < len(a.tos) {
		return int(a.tos[from]) - 1
	}
	return -1
}

func (a *Area[From, To]) Remove(from From, maintainOrder bool) int {
	if int(from) >= len(a.tos) || a.tos[from] == 0 {
		return -1
	}
	a.next--
	removedTo := a.tos[from]
	a.tos[from] = 0

	if removedTo == a.next {
		return int(removedTo) - 1
	}

	if maintainOrder {
		for i, to := range a.tos {
			if to > removedTo {
				a.tos[i] = to - 1
			}
		}
	} else {
		for i := len(a.tos) - 1; i >= 0; i-- {
			if a.tos[i] == a.next {
				a.tos[i] = removedTo
				break
			}
		}
	}
	return int(removedTo) - 1
}

func (a *Area[From, To]) Clear() {
	a.tos = make([]To, 0, cap(a.tos))
	a.next = 0
}

func (a Area[From, To]) Len() int {
	return len(a.tos)
}

func (a Area[From, To]) Empty() bool {
	return len(a.tos) == 0
}

type SparseMap[V any, SID constraints.Unsigned] struct {
	area         *Area[uint32, SID]
	values       []V
	resizeBuffer int
}

var _ HasCapacity = &SparseMap[int, uint]{}
var _ HasResizeBuffer = &SparseMap[int, uint]{}

func NewSparseMap[V any, SID constraints.Unsigned](opts ...Option[*SparseMap[V, SID]]) SparseMap[V, SID] {
	sm := SparseMap[V, SID]{
		values: make([]V, 0),
	}
	processOptions(&sm, opts)
	return sm
}

func (m *SparseMap[V, SID]) setResizeBuffer(resizeBuffer int) {
	m.resizeBuffer = resizeBuffer
}

func (m *SparseMap[V, SID]) setCapacity(capacity int) {
	m.values = make([]V, 0, capacity)
}

func (m *SparseMap[V, SID]) setArea(area *Area[uint32, SID]) {
	m.area = area
}

func (m SparseMap[V, SID]) Values() []V {
	return m.values
}

func (m *SparseMap[V, SID]) Set(key Identifier, value V) {
	*m.Take(key) = value
}

func (m *SparseMap[V, SID]) Get(key Identifier) V {
	p := m.Ptr(key)
	if p == nil {
		var empty V
		return empty
	}
	return *p
}

func (m *SparseMap[V, SID]) Ptr(key Identifier) *V {
	if key.Exists() {
		mapID := m.area.Peek(uint32(key))
		if mapID >= 0 && mapID < len(m.values) {
			return &m.values[mapID]
		}
	}
	return nil
}

func (m *SparseMap[V, SID]) Has(key Identifier) bool {
	if key.Exists() {
		mapID := m.area.Peek(uint32(key))
		if mapID >= 0 && mapID < len(m.values) {
			return true
		}
	}
	return false
}

func (m *SparseMap[V, SID]) Take(key Identifier) *V {
	if !key.Exists() {
		return nil
	}
	mapID := int(m.area.Translate(uint32(key)))
	if mapID >= len(m.values) {
		nextSize := mapID + m.resizeBuffer + 1
		m.values = resize(m.values, nextSize)
	}
	return &m.values[mapID]
}

func (m SparseMap[V, SID]) Empty() bool {
	return len(m.values) == 0
}

func (m SparseMap[V, SID]) Len() int {
	return len(m.values)
}

type DenseMap[V any, A constraints.Unsigned, L constraints.Unsigned] struct {
	area   *Area[uint32, A]
	local  Area[A, L]
	values []V
}

func NewDenseMap[V any, A constraints.Unsigned, L constraints.Unsigned](opts ...Option[*DenseMap[V, A, L]]) DenseMap[V, A, L] {
	sm := DenseMap[V, A, L]{
		values: make([]V, 0),
	}
	processOptions(&sm, opts)
	return sm
}

func (m *DenseMap[V, A, L]) setCapacity(capacity int) {
	m.values = make([]V, 0, capacity)
	m.local.setCapacity(capacity)
}

func (m *DenseMap[V, A, L]) setArea(area *Area[uint32, A]) {
	m.area = area
}

func (m *DenseMap[V, A, L]) indexOf(key Identifier) int {
	if !key.Exists() {
		return -1
	}
	areaID := m.area.Peek(uint32(key))
	if areaID < 0 {
		return -1
	}
	return m.local.Peek(A(areaID))
}

func (m DenseMap[V, A, L]) Values() []V {
	return m.values
}

func (m *DenseMap[V, A, L]) Set(key Identifier, value V) {
	*m.Take(key) = value
}

func (m *DenseMap[V, A, L]) Get(key Identifier) V {
	i := m.indexOf(key)
	if i == -1 {
		var empty V
		return empty
	}
	return m.values[i]
}

func (m *DenseMap[V, A, L]) Ptr(key Identifier) *V {
	i := m.indexOf(key)
	if i == -1 {
		return nil
	}
	return &m.values[i]
}

func (m *DenseMap[V, A, L]) Has(key Identifier) bool {
	i := m.indexOf(key)
	return i >= 0 && i < len(m.values)
}

func (m *DenseMap[V, A, L]) Take(key Identifier) *V {
	if !key.Exists() {
		return nil
	}
	areaID := m.area.Translate(uint32(key))
	index := m.local.Translate(areaID)
	if int(index) == len(m.values) {
		var empty V
		m.values = append(m.values, empty)
	}
	return &m.values[index]
}

func (m DenseMap[V, A, L]) Empty() bool {
	return len(m.values) == 0
}

func (m DenseMap[V, A, L]) Len() int {
	return len(m.values)
}

func (m *DenseMap[V, A, L]) Clear() {
	m.values = make([]V, 0, cap(m.values))
	m.local.Clear()
}

func (m *DenseMap[V, A, L]) Remove(key Identifier, maintainOrder bool) bool {
	if !key.Exists() {
		return false
	}
	areaID := m.area.Peek(uint32(key))
	if areaID == -1 {
		return false
	}
	index := m.local.Remove(A(areaID), maintainOrder)
	if index == -1 {
		return false
	}
	if maintainOrder {
		m.values = removeAt(m.values, index)
	} else {
		m.values = moveEndTo(m.values, index)
	}
	return true
}

type DenseKeyMap[V any, A constraints.Unsigned, L constraints.Unsigned] struct {
	area   *Area[uint32, A]
	local  Area[A, L]
	values []V
	keys   []Identifier
}

func NewDenseKeyMap[V any, A constraints.Unsigned, L constraints.Unsigned](opts ...Option[*DenseKeyMap[V, A, L]]) DenseKeyMap[V, A, L] {
	sm := DenseKeyMap[V, A, L]{
		values: make([]V, 0),
	}
	processOptions(&sm, opts)
	return sm
}

func (m *DenseKeyMap[V, A, L]) setCapacity(capacity int) {
	m.values = make([]V, 0, capacity)
	m.keys = make([]Identifier, 0, capacity)
	m.local.setCapacity(capacity)
}

func (m *DenseKeyMap[V, A, L]) setArea(area *Area[uint32, A]) {
	m.area = area
}

func (m *DenseKeyMap[V, A, L]) indexOf(key Identifier) int {
	if !key.Exists() {
		return -1
	}
	areaID := m.area.Peek(uint32(key))
	if areaID < 0 {
		return -1
	}
	return m.local.Peek(A(areaID))
}

func (m DenseKeyMap[V, A, L]) Values() []V {
	return m.values
}

func (m DenseKeyMap[V, A, L]) Keys() []Identifier {
	return m.keys
}

func (m *DenseKeyMap[V, A, L]) Set(key Identifier, value V) {
	*m.Take(key) = value
}

func (m *DenseKeyMap[V, A, L]) Get(key Identifier) V {
	i := m.indexOf(key)
	if i == -1 {
		var empty V
		return empty
	}
	return m.values[i]
}

func (m *DenseKeyMap[V, A, L]) Ptr(key Identifier) *V {
	i := m.indexOf(key)
	if i == -1 {
		return nil
	}
	return &m.values[i]
}

func (m *DenseKeyMap[V, A, L]) Has(key Identifier) bool {
	i := m.indexOf(key)
	return i >= 0 && i < len(m.values)
}

func (m *DenseKeyMap[V, A, L]) Take(key Identifier) *V {
	if !key.Exists() {
		return nil
	}
	areaID := m.area.Translate(uint32(key))
	index := m.local.Translate(areaID)
	if int(index) == len(m.values) {
		var empty V
		m.values = append(m.values, empty)
		m.keys = append(m.keys, key)
	}
	return &m.values[index]
}

func (m DenseKeyMap[V, A, L]) Empty() bool {
	return len(m.values) == 0
}

func (m DenseKeyMap[V, A, L]) Len() int {
	return len(m.values)
}

func (m *DenseKeyMap[V, A, L]) Clear() {
	m.values = make([]V, 0, cap(m.values))
	m.keys = make([]Identifier, 0, cap(m.keys))
	m.local.Clear()
}

func (m *DenseKeyMap[V, A, L]) Remove(key Identifier, maintainOrder bool) bool {
	if !key.Exists() {
		return false
	}
	areaID := m.area.Peek(uint32(key))
	if areaID == -1 {
		return false
	}
	index := m.local.Remove(A(areaID), maintainOrder)
	if index == -1 {
		return false
	}
	if maintainOrder {
		m.values = removeAt(m.values, index)
		m.keys = removeAt(m.keys, index)
	} else {
		m.values = moveEndTo(m.values, index)
		m.keys = moveEndTo(m.keys, index)
	}
	return true
}

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
