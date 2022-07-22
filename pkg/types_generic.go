package types

import "reflect"

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

var GameTypes = make(map[string]IGameType)
var GameTypesReflect = make(map[reflect.Type]IGameType)

type IGameType interface {
	GetName() string
	GetZero() any
	GetProps() NameMap[GameTypeProperty[any]]
	GetCollects() NameMap[GameTypeCollection[any]]
	Copy(source any, path []string) (any, IGameType)
	Ref(source *any, path []string) (any, IGameType)
	Set(source *any, path []string, value any) (any, IGameType)
}

type GameType[S any] struct {
	Name     string
	Zero     S
	Props    NameMap[GameTypeProperty[S]]
	Collects NameMap[GameTypeCollection[S]]
}

func New[S any](name string, zero S) *GameType[S] {
	t := &GameType[S]{
		Name:     name,
		Zero:     zero,
		Props:    NewNameMap[GameTypeProperty[S]](),
		Collects: NewNameMap[GameTypeCollection[S]](),
	}
	GameTypes[name] = t
	GameTypesReflect[reflect.TypeOf(zero)] = t
	return t
}

func (t GameType[S]) GetName() string {
	return t.Name
}

func (t GameType[S]) GetZero() any {
	return t.Zero
}

func (t GameType[S]) GetProps() NameMap[GameTypeProperty[any]] {
	return NameMap[GameTypeProperty[any]]{} // t.Props
}

func (t GameType[S]) GetCollects() NameMap[GameTypeCollection[any]] {
	return NameMap[GameTypeCollection[any]]{} // t.Collects
}

func (t GameType[S]) Get(name string) *GameTypeProperty[S] {
	if prop, ok := t.Props.Map[name]; ok {
		return &prop
	}
	return nil
}

func (t *GameType[S]) Add(props []GameTypeProperty[S]) *GameType[S] {
	for i := range props {
		t.Props.Add(props[i])
	}
	return t
}

func (t *GameType[S]) Copy(source any, path []string) (any, IGameType) {
	var ct IGameType = t
	c := any(source)

	for _, node := range path {
		p := t.GetProps().Get(node)
		if p.Copy == nil {
			return nil, nil
		}
		c = p.Copy(c)
		ct = p.Type
	}

	return c, ct
}

func (t *GameType[S]) Ref(source *any, path []string) (any, IGameType) {
	var ct IGameType = t
	var c *any = source

	for _, node := range path {
		p := t.GetProps().Get(node)
		if p.Ref == nil {
			return nil, nil
		}
		r := p.Ref(c)
		c = &r
		ct = p.Type
	}

	return c, ct
}

func (t *GameType[S]) Set(source *any, path []string, value any) (any, IGameType) {
	var ct IGameType = t
	var c *any = source
	// combination of ref & get
	for _, node := range path {
		p := t.GetProps().Get(node)
		if p.Copy == nil {
			return nil, nil
		}
		r := p.Ref(c)
		c = &r
		ct = p.Type
	}

	return c, ct
}

type GameObject[S any] struct {
	Value S
	Type  *GameType[S]
	Tags  []string
}

type GameTypeProperty[S any] struct {
	Name    string
	Type    IGameType
	Copy    func(source S) any
	Ref     func(source *S) any
	Set     func(source *S, value any)
	Virtual bool
}

func (p GameTypeProperty[S]) GetName() string {
	return p.Name
}

func (p GameTypeProperty[S]) IsReadOnly() bool {
	return p.Set == nil
}

type GameTypeCollection[S any] struct {
	Name     string
	Type     *IGameType
	Iterator func(source *S) GameObjectIterator
}

func (c GameTypeCollection[S]) GetName() string {
	return c.Name
}

type GameObjectIterator interface {
	Len() int
	HasNext() bool
	Next() any
	Remove()
}

func Copy[S any](source S, path []string) (any, IGameType) {
	t := GameTypesReflect[reflect.TypeOf(source)]

	return t.Copy(source, path)
}

func Ref[S any](source *S, path []string) (any, IGameType) {
	t := GameTypesReflect[reflect.TypeOf(source)]
	a := any(*source)

	return t.Ref(&a, path)
}

func Set[S any](source *S, path []string, value any) (any, IGameType) {
	t := GameTypesReflect[reflect.TypeOf(source)]
	a := any(*source)

	return t.Set(&a, path, value)
}
