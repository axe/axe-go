package id

type Option func(optionable any)

func processOptions[O any](optionable O, options []Option) {
	for _, option := range options {
		option(optionable)
	}
}

type HasResizeBuffer interface {
	setResizeBuffer(int)
}

func WithResizeBuffer(resizeBuffer int) Option {
	return func(optionable any) {
		if has, ok := optionable.(HasResizeBuffer); ok {
			has.setResizeBuffer(resizeBuffer)
		}
	}
}

type HasCapacity interface {
	setCapacity(int)
}

func WithCapacity(capacity int) Option {
	return func(optionable any) {
		if has, ok := optionable.(HasCapacity); ok {
			has.setCapacity(capacity)
		}
	}
}

type HasArea[A any] interface {
	setArea(*A)
}

func WithArea[A any](area *A) Option {
	return func(optionable any) {
		if has, ok := optionable.(HasArea[A]); ok {
			has.setArea(area)
		}
	}
}

type KeyValue[V any] struct {
	Key   Identifier
	Value V
}

type HasSet[V any] interface {
	Set(key Identifier, value V)
}

func WithKeyValues[V any](kv []KeyValue[V]) Option {
	return func(optionable any) {
		if has, ok := optionable.(HasSet[V]); ok {
			setKeyValues(has, kv)
		}
	}
}

func WithStringMap[V any](m map[string]V) Option {
	return func(optionable any) {
		if has, ok := optionable.(HasSet[V]); ok {
			setStringMap(has, m)
		}
	}
}

func WithMap[V any](m map[Identifier]V) Option {
	return func(optionable any) {
		if has, ok := optionable.(HasSet[V]); ok {
			setMap(has, m)
		}
	}
}

func setKeyValues[V any](target HasSet[V], kv []KeyValue[V]) {
	for _, pair := range kv {
		target.Set(pair.Key, pair.Value)
	}
}

func setMap[V any](target HasSet[V], m map[Identifier]V) {
	for k, v := range m {
		target.Set(k, v)
	}
}

func setStringMap[V any](target HasSet[V], m map[string]V) {
	for k, v := range m {
		target.Set(Get(k), v)
	}
}
