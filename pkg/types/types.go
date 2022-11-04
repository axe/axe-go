package types

import (
	"reflect"
	"strings"
)

type Named interface {
	GetName() string
}

type NameMap[N Named] struct {
	Values []N
	Map    map[string]N
}

func NewNameMap[V Named]() NameMap[V] {
	return NameMap[V]{
		Values: make([]V, 0),
		Map:    make(map[string]V, 0),
	}
}

func (m *NameMap[V]) Add(value V) {
	m.Values = append(m.Values, value)
	m.Map[value.GetName()] = value
}

func (m *NameMap[V]) Remove(value V) {
	delete(m.Map, value.GetName())
}

func (m NameMap[N]) Get(name string) N {
	return m.Map[name]
}

var GameTypes = make(map[string]*GameType)
var GameTypesReflect = make(map[reflect.Type]*GameType)

type GameType struct {
	Name     string
	Zero     any
	Props    NameMap[*GameTypeProperty]
	Collects NameMap[*GameTypeCollection]
}

func New[S any](name string, zero S) *GameType {
	t := &GameType{
		Name:     name,
		Zero:     zero,
		Props:    NewNameMap[*GameTypeProperty](),
		Collects: NewNameMap[*GameTypeCollection](),
	}
	GameTypes[name] = t
	GameTypesReflect[reflect.TypeOf(zero)] = t
	return t
}

func (t GameType) Instance() any {
	i := t.Zero
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Pointer {
		i = v.Elem()
	}
	return i
}

func (t GameType) Get(name string) *GameTypeProperty {
	return t.Props.Map[name]
}

func (t *GameType) Add(props []GameTypeProperty) *GameType {
	for i := range props {
		t.Props.Add(&props[i])
	}
	return t
}

func (t *GameType) Access(path []string) *GameTypeAccess {
	props := make([]*GameTypeProperty, 0)
	cache := make([]any, 0)

	currentType := t
	canRef := true
	canCopy := true
	canSet := true

	for _, node := range path {
		prop := currentType.Props.Get(node)
		if prop == nil || prop.Type == nil {
			return nil
		}
		props = append(props, prop)
		currentType = prop.Type
		if prop.Ref == nil {
			canRef = false
		}
		if prop.Ref == nil && prop.Copy == nil {
			canCopy = false
		}
		if prop.Ref == nil && prop.Set == nil {
			canSet = false
		}
	}

	if !canRef && !canSet && !canCopy {
		return nil
	}

	if !canRef {
		for _, p := range props {
			cache = append(cache, p.Type.Instance())
		}
	}

	return &GameTypeAccess{
		Base:    t,
		Path:    path,
		Props:   props,
		Cache:   cache,
		CanRef:  canRef,
		CanCopy: canCopy,
		CanSet:  canSet,
	}
}

type GameTypeAccess struct {
	Base    *GameType
	Path    []string
	Props   []*GameTypeProperty
	Cache   []any
	CanSet  bool
	CanCopy bool
	CanRef  bool
}

func (a GameTypeAccess) Last() *GameTypeProperty {
	return a.Props[len(a.Props)-1]
}

func (a GameTypeAccess) Type() *GameType {
	return a.Last().Type
}

func (a GameTypeAccess) Copy(source any) any {
	if !a.CanCopy {
		panic("Accessor cannot copy requested value")
	}
	currentValue := source

	last := len(a.Path) - 1
	for i, p := range a.Props {
		if p.Ref != nil && i != last {
			currentValue = p.Ref(currentValue)
		} else {
			currentValue = p.Copy(currentValue)
		}
	}

	return currentValue
}

func (a GameTypeAccess) Ref(source any) any {
	if !a.CanRef {
		panic("Accessor cannot ref requested value")
	}
	currentValue := source

	for _, p := range a.Props {
		currentValue = p.Ref(currentValue)
	}

	return currentValue
}

func (a GameTypeAccess) Set(source any, value any) bool {
	if !a.CanSet {
		panic("Accessor cannot set requested value")
	}

	if a.CanRef {
		ref := a.Ref(source)
		return copy(value, ref)
	} else {
		currentValue := source

		// last := len(a.Path) - 1
		for _, p := range a.Props {
			currentValue = p.Ref(currentValue)
		}
	}

	return false
}

type GameObject struct {
	Value any
	Type  *GameType
	Tags  []string
	Props map[GameObjectPath]*GameObjectProperty
}

type GameObjectProperty struct {
	Prop   *GameTypeProperty
	Cached any
	Stale  bool
	Props  map[string]*GameObjectProperty
}

type GameObjectAccess struct {
	Access GameTypeAccess
}

type GameTypeMeta int

var typeMetaNames = make(map[GameTypeMeta]string)

func (m GameTypeMeta) GetName() string {
	return typeMetaNames[m]
}

func (m GameTypeMeta) String() string {
	return typeMetaNames[m]
}

func NewMeta(name string) GameTypeMeta {
	key := GameTypeMeta(len(typeMetaNames))
	typeMetaNames[key] = name
	return key
}

type GameTypeProperty struct {
	Name    string
	Type    *GameType
	Copy    func(source any) any
	Ref     func(source any) any
	Set     func(source any, value any)
	Virtual bool
	Default any
	Meta    map[GameTypeMeta]any
}

func (p GameTypeProperty) GetName() string {
	return p.Name
}

func (p GameTypeProperty) IsReadOnly() bool {
	return p.Set == nil
}

type GameTypeCollection struct {
	Name     string
	Type     *GameType
	Iterator func(source any) GameObjectIterator
}

func (c GameTypeCollection) GetName() string {
	return c.Name
}

type GameObjectIterator interface {
	Len() int
	HasNext() bool
	Next() (value any, key any)
	Remove()
	Reset()
}

type iteratorSlice struct {
	index             int
	slice             []any
	keepOrderOnRemove bool
}

func (i iteratorSlice) Len() int {
	return len(i.slice)
}
func (i iteratorSlice) HasNext() bool {
	return i.index < len(i.slice)
}
func (i *iteratorSlice) Next() (any, any) {
	nextIndex := i.index
	next := i.slice[nextIndex]
	i.index++
	return next, nextIndex
}
func (i *iteratorSlice) Remove() {
	k := i.index - 1
	if i.keepOrderOnRemove {
		i.slice = append(i.slice[:k], i.slice[k+1:]...)
	} else {
		last := len(i.slice) - 1
		i.slice[k] = i.slice[last]
		i.slice = i.slice[:last]
	}
	i.index--
}
func (i *iteratorSlice) Reset() {
	i.index = 0
}

func NewIteratorSlice(slice []any, keepOrderOnRemove bool) GameObjectIterator {
	return &iteratorSlice{0, slice, keepOrderOnRemove}
}

type iteratorMap[K comparable, V any] struct {
	index  int
	keys   []K
	theMap map[K]V
}

func (i iteratorMap[K, V]) Len() int {
	return len(i.keys)
}
func (i iteratorMap[K, V]) HasNext() bool {
	return i.index < len(i.keys)
}
func (i *iteratorMap[K, V]) Next() (any, any) {
	nextKey := i.keys[i.index]
	next := i.theMap[nextKey]
	i.index++
	return next, nextKey
}
func (i *iteratorMap[K, V]) Remove() {
	keyIndex := i.index - 1
	key := i.keys[keyIndex]
	delete(i.theMap, key)
	last := len(i.keys) - 1
	i.keys[keyIndex] = i.keys[last]
	i.keys = i.keys[:last]
	i.index--
}
func (i *iteratorMap[K, V]) Reset() {
	i.index = 0
}

func NewIteratorMap[K comparable, V any](m map[K]V) GameObjectIterator {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return &iteratorMap[K, V]{0, keys, m}
}

func Copy(source any, path []string) any {
	t := GameTypesReflect[typeOf(source)]

	return t.Access(path).Copy(source)
}

func Ref(source any, path []string) any {
	t := GameTypesReflect[typeOf(source)]

	return t.Access(path).Ref(source)
}

func Set(source any, path []string, value any) {
	t := GameTypesReflect[typeOf(source)]

	t.Access(path).Set(source, value)
}

func typeOf(value any) reflect.Type {
	r := reflect.TypeOf(value)
	if r != nil && r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return r
}

func copy(src any, dst any) bool {
	s := reflect.ValueOf(src)
	d := reflect.ValueOf(dst)
	copied := reflect.Copy(d, s) > 0
	return copied
}

type GameObjectPath int

var paths = make(map[GameObjectPath][]string)
var pathReverse = make(map[string]GameObjectPath)
var pathDelimiter = "."

func GetPath(props []string) GameObjectPath {
	full := strings.Join(props, pathDelimiter)
	if existing, exists := pathReverse[full]; exists {
		return existing
	}
	id := GameObjectPath(len(paths))
	paths[id] = props
	pathReverse[full] = id
	return id
}

func (p GameObjectPath) Props() []string {
	return paths[p]
}
