package ecs

type World struct {
	entities              *data.SparseList[Entity]
	entityComponentSize   uint32
	components            []BaseComponent
	componentInstanceSize uint32
	componentFreeSize     uint32
}

type WorldOptions struct {
	entityCount           uint32
	entityFreeSize        uint32
	componentCount        uint32
	componentInstanceSize uint32
	componentFreeSize     uint32
	entityComponentSize   uint32
}

const (
	DEFAULT_ENTITY_COUNT            = 128
	DEFAULT_ENTITY_FREE_SIZE        = 32
	DEFAULT_COMPONENT_COUNT         = 8
	DEFAULT_ENTITY_COMPONENT_SIZE   = 4
	DEFAULT_COMPONENT_INSTANCE_SIZE = 128
	DEFAULT_COMPONENT_FREE_SIZE     = 32
)

func determineValue(userValue uint32, defaultValue uint32) uint32 {
	if userValue == 0 {
		return defaultValue
	}
	return userValue
}

func NewWorld(options WorldOptions) *World {
	entityCount := determineValue(options.entityCount, DEFAULT_ENTITY_COUNT)
	entityFree := determineValue(options.entityFreeSize, DEFAULT_ENTITY_FREE_SIZE)

	return &World{
		entities:              data.NewSparseList[Entity](entityCount, entityFree),
		entityComponentSize:   determineValue(options.entityComponentSize, DEFAULT_ENTITY_COMPONENT_SIZE),
		components:            make([]BaseComponent, 0, determineValue(options.componentCount, DEFAULT_COMPONENT_COUNT)),
		componentInstanceSize: determineValue(options.componentInstanceSize, DEFAULT_COMPONENT_INSTANCE_SIZE),
		componentFreeSize:     determineValue(options.componentFreeSize, DEFAULT_COMPONENT_FREE_SIZE),
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
	entity.flags = 0
	entity.components = make([]uint32, this.entityComponentSize)

	return entity
}

func (this *World) Destroy(entity *Entity) {
	for componentId, componentIndex := range entity.Components {
		if (entity.Has & (1 << componentId)) != 0 {
			this.components[componentId].free(componentIndex)
		}
	}
	entity.Has = 0
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
		if (entity.Has & componentBits) != componentBits {
			return
		}

		handle(entity)
	})
}
