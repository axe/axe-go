package types

import (
	"fmt"

	"github.com/axe/axe-go/pkg/util"
)

// Using
// - create all your types with types.New("typeName")
// - define all your types after creation with TypeInt.Define(...)
// - at the beginning of game play create all the access you'll need to validate and ensure values that need to be set/get can be.

// Features
// - define your game types so your game data can be dynamically set/get
// - this allows runtime access for animation, scripting, debugging, and editting purposes
// - it's extenable allowing users to define their own types and add to existing types
// - values can be a:
//   - gettable/settable prop on a type
//   - dynamic value on a type
//   - dynamic value on a type that takes arguments
//   - collection of values keyed by another type, ie. lists and maps
// - convert an expression like the following to an accessible value:
//   - "game.entities.player.head.position"
//   - "game.entities.withFlag('bullets').expire"
//   - "game.ui.pauseButton.disabled"
// - the above expressions can be used in a scripting language, in a console prompt to inspect game state, or be stored in a animation file to point to what value it animates

var (
	typeIDs = NewIncrementor[uint16](0, 1)
	types   = NewNameMap(getTypeName, 128, true, true)
)

func getTypeName(t *Type) string { return t.Name }

// Returns all defined types.
func Types() NameMap[*Type] {
	return types
}

// Creates a new type with the given name.
func New(name string) *Type {
	util.Assert(!types.Has(name), "type %s already created", name)
	t := &Type{ID: typeIDs.Get(), Name: name}
	types.Add(t)
	return t
}

// A lightweight type with an ID (auto-incrementing starting at 0) and a name.
// A definition can be attached to it later.
type Type struct {
	ID   uint16
	Name string
	Def  *TypeDef
}

// Sets the definition for the type.
func (t *Type) Define(def TypeDef) { t.Def = &def }

// Returns a new instance of the given type.
func (t Type) New() Value {
	util.Assert(t.Def != nil && t.Def.Create != nil, "create is not defined for %s", t.Name)
	return t.Def.Create()
}

// Gets a prop with the given name.
func (t Type) Prop(name string) *Prop {
	p, _ := t.Def.Props.Get(name)
	return p
}

// A value for a type (void* in C/C++)
type Value = any

// A value for a type but its explicitly an addressable value.
type Ptr = any

// A value with given defined type.
type TypedValue struct {
	Value Value
	Type  *Type
}

// Something that can iterate or fetch values from a collection of key-values.
type Collection interface {
	// The number of items in the collection.
	Len() int
	// Gets the value of the item in the collection with the given key
	Get(key Value) (Value, bool)
	// Sets the value of the item in the collection with the given key
	Set(key Value, value Value) bool
	// Adds the value to the collection
	Add(key Value, value Value) bool
	// Returns true if a value exists in the collection with the given key
	Contains(key Value) bool
	// Returns an iterator for the collection.
	Iterator() Iterator
}

// Iterates over values and keys.
type Iterator interface {
	// Returns true if there is a next key & value
	HasNext() bool
	// Returns the next key and value
	Next() (value Value, key Value)
	// Removes the last key and value returned by Next and returns true.
	// If they cannot be removed or Next() was never called then false is returned.
	Remove() bool
	// Sets the value for the last key returned by Next and returns true.
	// If it cannot be set or Next() was never called then false is returned.
	Set(value Value) bool
	// Resets the iterator to the beginning of the key-values.
	Reset()
}

// A definition for a type which holds key-values
type CollectionDef struct {
	// The key type.
	Key *Type
	// The value type.
	Value *Type
	// Returns a collection given a value.
	Collect func(value Value) Collection
}

// A definition for a type.
type TypeDef struct {
	// Creates a value of this type if possible.
	Create func() Value
	// If the type supports it this value is convertible to a string.
	ToString func(Value) string
	// If the type supports it this value is convertible from a string.
	FromString func(x string) (Value, error)
	// The properties on this type.
	Props NameMap[*Prop]
	// A definition if this type holds a collection of key-values.
	Collection *CollectionDef
	// Metadata on the Type
	Meta map[Meta]Value
}

var typeMetaNames = make(map[Meta]string)

// Metadata stored on a type/prop
type Meta int

func (m Meta) GetName() string { return typeMetaNames[m] }
func (m Meta) String() string  { return typeMetaNames[m] }

// Creates a new Meta with the given name.
func NewMeta(name string) Meta {
	key := Meta(len(typeMetaNames))
	typeMetaNames[key] = name
	return key
}

// A property on a type
type Prop struct {
	// The name of the property
	Name string
	// The type of the prop
	Type *Type
	// A function to return a copy of a prop value on a source value.
	Get GetArgsFn
	// A function to return a pointer to a prop value from a source value.
	Ref RefFn
	// A function to set a prop value on a source value.
	Set SetFn
	// A virtual prop is one that can be get or set but should be transferred.
	Virtual bool
	// The default value for a property. Why is this here?
	Default any
	// Metadata on the Prop
	Meta map[Meta]Value
}

func getPropName(p *Prop) string { return p.Name }

func (p Prop) GetName() string  { return p.Name }
func (p Prop) IsReadOnly() bool { return p.Set == nil }

// Returns a NameMap given a list of props.
func Props(props []Prop) NameMap[*Prop] {
	m := NewNameMap(getPropName, len(props), true, true)
	for i := range props {
		m.Set(&props[i])
	}
	return m
}

// A get function that returns a copy of a value.
type GetFn = func(source Value) Value

// A get function that can accept args. Source is
type GetArgsFn = func(source Value, args []Value) Value

// A ref function. Source is always a ptr and so is the returned value.
type RefFn = func(source Value) Value

// A set function. Source is always a ptr and value is not.
type SetFn = func(source Value, value Value) bool

// A path node represents a path starting with a base value
// and ending with a desired value.
type PathNode struct {
	Token string
	Key   Value
	Args  []Value
}

func (n PathNode) toAccessNode(t *Type) (node *AccessNode, err error) {
	if n.Token != "" {
		prop := t.Prop(n.Token)
		if prop != nil {
			node = accessNodeForProp(prop, n.Args)
			return
		}
		if t.Def.Collection != nil {
			key := n.Key
			keyType := t.Def.Collection.Key
			if key == nil && keyType.Def.FromString != nil {
				key, err = keyType.Def.FromString(n.Token)
				if err != nil {
					return
				}
			}
			if key != nil {
				node = accessNodeForKey(key, t.Def.Collection)
				return
			}
		}
	} else if n.Key != nil && t.Def.Collection != nil {
		node = accessNodeForKey(n.Key, t.Def.Collection)
		return
	}
	err = fmt.Errorf("%s not found on type %s", n.Token, t.Name)
	return
}
func accessNodeForProp(prop *Prop, args []Value) *AccessNode {
	node := &AccessNode{
		Prop: prop,
		Type: prop.Type,
		Args: args,
		Ref:  prop.Ref,
		Set:  prop.Set,
	}
	if prop.Get != nil {
		node.Get = func(source Value) Value {
			return prop.Get(source, args)
		}
	}
	return node
}
func accessNodeForKey(key Value, def *CollectionDef) *AccessNode {
	return &AccessNode{
		Key:  key,
		Type: def.Value,
		Get: func(source Value) Value {
			v, _ := def.Collect(source).Get(key)
			return v
		},
		Set: func(source, value Value) bool {
			return def.Collect(source).Set(key, value)
		},
	}
}

// Creates a new simple path from a stream of tokens.
func NewPathFromTokens(tokens []string) []PathNode {
	nodes := make([]PathNode, len(tokens))
	for i, token := range tokens {
		nodes[i].Token = token
	}
	return nodes
}

// Starting with a type and path of props/keys return something
// that will allow us to set & get a value.
func (t *Type) Access(path []PathNode) (Access, error) {
	nodes := make([]AccessNode, 0, len(path))
	cache := make([]Value, 0, len(path))
	cacheRef := make([]bool, len(path))

	currentType := t
	canRef := true
	canGet := true
	canSet := true

	for _, pathNode := range path {
		node, err := pathNode.toAccessNode(currentType)
		if err != nil || node == nil || node.Type == nil {
			if err == nil {
				err = fmt.Errorf("%s could not be converted to an access node", pathNode.Token)
			}
			return Access{}, err
		}
		nodes = append(nodes, *node)
		currentType = node.Type
		if node.Ref == nil {
			canRef = false
		}
		if node.Ref == nil && node.Get == nil {
			canGet = false
		}
		if node.Ref == nil && node.Set == nil {
			canSet = false
		}
	}

	if !canRef && !canSet && !canGet {
		return Access{}, fmt.Errorf("its not possible to get, set, or ref the given path")
	}

	if !canRef {
		for _, p := range nodes {
			cache = append(cache, p.Type.New())
		}
	}

	return Access{
		Base:     t,
		Path:     path,
		Nodes:    nodes,
		Cache:    cache,
		CacheRef: cacheRef,
		CanRef:   canRef,
		CanGet:   canGet,
		CanSet:   canSet,
	}, nil
}

// Access to value down a path starting from a given type.
type Access struct {
	Base     *Type
	Path     []PathNode
	Nodes    []AccessNode
	Cache    []Value
	CacheRef []bool
	CanSet   bool
	CanGet   bool
	CanRef   bool
}

// A node along an access path with the necessary values.
type AccessNode struct {
	Prop *Prop
	Key  Value
	Type *Type
	Args []Value
	Get  GetFn
	Ref  RefFn
	Set  SetFn
}

// The last node in the access path. If there is no path then nil
// is returned.
func (a Access) Last() *AccessNode {
	i := len(a.Nodes) - 1
	if i == -1 {
		return nil
	}
	return &a.Nodes[i]
}

// The return type for the access.
func (a Access) Type() *Type {
	last := a.Last()
	if last != nil {
		return last.Type
	}
	return a.Base
}

// Gets a copy of a value in the access path from the source.
func (a Access) Get(source Value) Value {
	if !a.CanGet {
		panic("Accessor cannot copy requested value")
	}
	currentValue := source
	last := len(a.Path) - 1
	for i, n := range a.Nodes {
		if n.Ref != nil && i != last {
			currentValue = n.Ref(currentValue)
		} else {
			currentValue = n.Get(currentValue)
		}
	}

	return currentValue
}

// Gets a reference to a value in the access path from the source.
// A reference points to an addressable value meaning changing
// the value directly is possible.
func (a Access) Ref(source Value) Value {
	if !a.CanRef {
		panic("Accessor cannot ref requested value")
	}
	currentValue := source
	for _, n := range a.Nodes {
		currentValue = n.Ref(currentValue)
	}

	return currentValue
}

// Sets a value to a given source in the access path.
func (a Access) Set(source Ptr, value Value) bool {
	if !a.CanSet {
		panic("Accessor cannot set requested value")
	}

	isRef := true
	currentValue := source
	last := len(a.Nodes) - 1
	for i := 0; i <= last; i++ {
		a.Cache[i] = currentValue
		a.CacheRef[i] = isRef

		if i == last {
			break
		}
		n := a.Nodes[i]
		if n.Ref != nil && isRef {
			currentValue = n.Ref(currentValue)
		} else {
			isRef = false
			currentValue = n.Get(currentValue)
		}
		if currentValue == nil {
			return false
		}
	}

	if isRef {
		a.Nodes[last].Set(a.Cache[last], value)
	} else {
		prev := value
		for i := last; i >= 0; i-- {
			a.Nodes[i].Set(a.Cache[i], prev)
			if a.CacheRef[i] {
				break
			}
			prev = a.Cache[i]
		}
	}

	return true
}

/*
type Object struct {
	Value any
	Type  *Type
	Tags  []string
	Props map[GameObjectPath]*GameObjectProperty
}

type GameObjectProperty struct {
	Prop   *Prop
	Cached any
	Stale  bool
	Props  map[string]*GameObjectProperty
}

type GameObjectAccess struct {
	Access Access
}

type GameTypeCollection struct {
	Name     string
	Type     *TypeDef
	Iterator func(source any) GameObjectIterator
}

func (c GameTypeCollection) GetName() string {
	return c.Name
}

type GameObjectIterator interface {
	Len() int
	HasNext() bool
	Next() (value Value, key Value)
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
*/
