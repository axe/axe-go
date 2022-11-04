package axe

import "github.com/axe/axe-go/pkg/ds"

type NamedComponent interface {
	ComponentName() string
}

type DefinedComponent[T any] struct {
	Name         string
	DefaultValue T
	Common       bool
}

var _ NamedComponent = &DefinedComponent[string]{}

func DefineComponent[T any](name string, defaultValue T, common bool) DefinedComponent[T] {
	return DefinedComponent[T]{
		Name:         name,
		DefaultValue: defaultValue,
		Common:       common,
	}
}

func (c *DefinedComponent[T]) ComponentName() string {
	return c.Name
}

func (c *DefinedComponent[T]) Enable(world *World) *ComponentStorage[T] {
	if storage, exists := world.storageMap[c.Name]; exists {
		return storage.(*ComponentStorage[T])
	}

	size := world.Options.EntityMax
	valueSize := size
	if !c.Common {
		size = world.Options.UncommonComponentCount
		valueSize = 0
	}

	storage := &ComponentStorage[T]{
		Index:        len(world.storage),
		Component:    c,
		Values:       make([]T, size),
		ValuesSize:   valueSize,
		StagedValues: make([]T, world.Options.StageSize),
	}

	world.storage = append(world.storage, storage)
	world.storageMap[c.Name] = storage

	return storage
}

func (c *DefinedComponent[T]) AddSystem(world *World, system System[T]) bool {
	storage := c.StorageFor(world)
	adding := storage != nil
	if adding {
		world.systems = append(world.systems, system(storage))
	}
	return adding
}

func (c *DefinedComponent[T]) StorageFor(world *World) *ComponentStorage[T] {
	if storage, exists := world.storageMap[c.Name]; exists {
		if componentStorage, matches := storage.(*ComponentStorage[T]); matches {
			return componentStorage
		}
	}
	return nil
}

func (c *DefinedComponent[T]) Get(world *World, entity *Entity) *T {
	storage := c.StorageFor(world)
	if storage != nil {
		return storage.Take(entity)
	}
	return nil
}

func (c *DefinedComponent[T]) Set(world *World, entity *Entity, value T) bool {
	ptr := c.Get(world, entity)
	canSet := ptr != nil
	if canSet {
		*ptr = value
	}
	return canSet
}

type BaseComponentStorage interface {
	NamedComponent() NamedComponent
	SetDefaultValue(e *Entity)
	Unstage(e *Entity, stageId int)
}

type ComponentStorage[T any] struct {
	Index        int
	Component    *DefinedComponent[T]
	Values       []T
	ValuesSize   int
	StagedValues []T
}

var _ BaseComponentStorage = &ComponentStorage[string]{}

func (s *ComponentStorage[T]) NamedComponent() NamedComponent {
	return s.Component
}

func (s *ComponentStorage[T]) SetDefaultValue(e *Entity) {
	s.Set(e, s.Component.DefaultValue)
}
func (s *ComponentStorage[T]) Set(entity *Entity, value T) {
	entity.has |= (1 << s.Index)
	ptr := s.Take(entity)
	*ptr = value
}
func (s *ComponentStorage[T]) Unstage(entity *Entity, stageId int) {
	ptr := s.Get(entity)
	if ptr != nil {
		*ptr = s.StagedValues[-stageId-1]
	}
}
func (s *ComponentStorage[T]) Get(entity *Entity) *T {
	if (1<<s.Index)&entity.has == 0 {
		return nil
	}
	return s.Take(entity)
}
func (s *ComponentStorage[T]) Take(entity *Entity) *T {
	if entity.id > 0 {
		if s.Component.Common {
			return &s.Values[entity.id]
		} else {
			if entity.uncommon == nil {
				entity.uncommon = make([]int, 0, s.Index+1)
			}
			for len(entity.uncommon) <= s.Index {
				entity.uncommon = append(entity.uncommon, -1)
			}
			if entity.uncommon[s.Index] == -1 {
				entity.uncommon[s.Index] = s.ValuesSize
				s.ValuesSize++
			}
			return &s.Values[entity.uncommon[s.Index]]
		}
	} else {
		return &s.StagedValues[-entity.id-1]
	}
}

type BaseSystem interface {
	GameSystem
}

type System[T any] func(storage *ComponentStorage[T]) BaseSystem

type World struct {
	Options    WorldOptions
	live       *ds.SparseList[Entity]
	staging    *ds.List[*Entity]
	storage    []BaseComponentStorage
	storageMap map[string]BaseComponentStorage
	systems    []BaseSystem
}

type WorldOptions struct {
	EntityMax              int
	DestroyMax             int
	UncommonComponentCount int
	StageSize              int
}

var _ GameSystem = &World{}

func NewWorld(options WorldOptions) *World {
	return &World{
		Options:    options,
		live:       ds.NewSparseList[Entity](options.EntityMax, options.DestroyMax),
		staging:    ds.NewList[*Entity](options.StageSize),
		storage:    make([]BaseComponentStorage, 0, 16),
		storageMap: make(map[string]BaseComponentStorage, 16),
		systems:    make([]BaseSystem, 0, 16),
	}
}

func (world World) Enabled(c NamedComponent) bool {
	return world.storageMap[c.ComponentName()] != nil
}

func (world *World) Init(game *Game) error {
	for _, system := range world.systems {
		err := system.Init(game)
		if err != nil {
			return err
		}
	}
	return nil
}

func (world *World) Unstage() {
	for i := 0; i < world.staging.Size; i++ {
		live, liveIndex := world.live.Take()
		staged := world.staging.Items[i]
		*live = *staged
		live.id = liveIndex
		for _, storage := range world.storage {
			storage.Unstage(live, staged.id)
		}
	}
	world.staging.Clear()
}

func (world *World) Update(game *Game) {
	world.Unstage()
	for _, system := range world.systems {
		system.Update(game)
	}
}

func (world *World) Destroy() {
	for _, system := range world.systems {
		system.Destroy()
	}
}

func (world *World) Free(entity *Entity) {
	// TODO
}

type Entity struct {
	Flags    uint
	Tag      string
	id       int
	has      uint
	uncommon []int
}

type EntityCreate struct {
	Parent     *Entity
	Flags      uint
	Components []NamedComponent
}

func (world *World) Create() *Entity {
	return world.CreateWith(EntityCreate{})
}

func (world *World) CreateWith(create EntityCreate) *Entity {
	id := -(world.staging.Size + 1)
	e := &Entity{
		id:    id,
		Flags: create.Flags,
	}
	if create.Components != nil {
		for _, c := range create.Components {
			storage := world.storageMap[c.ComponentName()]
			if storage != nil {
				storage.SetDefaultValue(e)
			}
		}
	}
	world.staging.Add(e)
	return e
}

/*
func Test() {
	world := NewWorld(WorldOptions{
		EntityMax: 128,
		StageSize: 16,
	})

	transformStore := TRANSFORM2.Enable(world)
	TRANSFORM2.AddSystem(world, NewTransformSystem[Vec2[float32]])

	e := world.Create()
	transformStore.Set(e, Transform2{Local: 23})

	world.Unstage()
}
*/
