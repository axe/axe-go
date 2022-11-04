package ecs

import "github.com/axe/axe-go/pkg/ds"

type World struct {
	entities              *ds.SparseList[Entity]
	entityComponentSize   int
	components            []BaseComponent
	componentInstanceSize int
	componentFreeSize     int
}

type WorldOptions struct {
	EntityCount           int
	EntityFreeSize        int
	ComponentCount        int
	ComponentInstanceSize int
	ComponentFreeSize     int
	EntityComponentSize   int
}

const (
	DEFAULT_ENTITY_COUNT            = 128
	DEFAULT_ENTITY_FREE_SIZE        = 32
	DEFAULT_COMPONENT_COUNT         = 8
	DEFAULT_ENTITY_COMPONENT_SIZE   = 4
	DEFAULT_COMPONENT_INSTANCE_SIZE = 128
	DEFAULT_COMPONENT_FREE_SIZE     = 32
)

func determineValue(userValue int, defaultValue int) int {
	if userValue == 0 {
		return defaultValue
	}
	return userValue
}

func NewWorld(options WorldOptions) *World {
	entityCount := determineValue(options.EntityCount, DEFAULT_ENTITY_COUNT)
	entityFree := determineValue(options.EntityFreeSize, DEFAULT_ENTITY_FREE_SIZE)

	return &World{
		entities:              ds.NewSparseList[Entity](entityCount, entityFree),
		entityComponentSize:   determineValue(options.EntityComponentSize, DEFAULT_ENTITY_COMPONENT_SIZE),
		components:            make([]BaseComponent, 0, determineValue(options.ComponentCount, DEFAULT_COMPONENT_COUNT)),
		componentInstanceSize: determineValue(options.ComponentInstanceSize, DEFAULT_COMPONENT_INSTANCE_SIZE),
		componentFreeSize:     determineValue(options.ComponentFreeSize, DEFAULT_COMPONENT_FREE_SIZE),
	}
}

func (this *World) CreateWith(create EntityCreate) *Entity {
	e := this.Create()
	e.Flags = create.Flags

	for _, componentId := range create.Components {
		this.components[componentId].add(e)
	}

	return e
}

func (this *World) Create() *Entity {
	entity, id := this.entities.Take()
	entity.id = uint32(id)
	entity.has = 0
	entity.Flags = 0
	entity.components = make([]int, this.entityComponentSize)

	return entity
}

func (this *World) Destroy(entity *Entity) {
	for componentId, componentIndex := range entity.components {
		if (entity.has & (1 << componentId)) != 0 {
			this.components[componentId].free(componentIndex)
		}
	}
	entity.has = 0
	entity.Flags = 0
}

func (this *World) Search(search EntitySearch, handle func(entity *Entity)) {
	var componentBits uint64 = 0
	for _, componentId := range search.Components {
		componentBits = componentBits | (1 << componentId)
	}

	this.entities.Iterate(func(entity *Entity, _ int, _ int) {
		if search.FlagMatch != nil && !search.FlagMatch(search.Flags, entity.Flags) {
			return
		}
		if (entity.Flags & search.Flags) != search.Flags {
			return
		}
		if (entity.has & componentBits) != componentBits {
			return
		}

		handle(entity)
	})
}
