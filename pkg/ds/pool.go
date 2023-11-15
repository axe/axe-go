package ds

import "github.com/axe/axe-go/pkg/util"

type PoolCreate[V any] func() V

type PoolPrepare[V any] func(V) V

type Pool[V any] struct {
	create    PoolCreate[V]
	prepare   PoolPrepare[V]
	free      []V
	freeCount int
}

func NewPool[V any](inital int, create PoolCreate[V], prepare PoolPrepare[V]) Pool[V] {
	pool := Pool[V]{
		free:      make([]V, inital),
		freeCount: inital,
		create:    create,
		prepare:   prepare,
	}
	for i := range pool.free {
		pool.free[i] = create()
	}
	return pool
}

func (p *Pool[V]) Get() V {
	if p.freeCount == 0 {
		return p.create()
	} else {
		p.freeCount--
		return p.prepare(p.free[p.freeCount])
	}
}

func (p *Pool[V]) Free(value V) {
	p.free = util.SliceEnsureSize(p.free, p.freeCount+1)
	p.free[p.freeCount] = value
	p.freeCount++
}
