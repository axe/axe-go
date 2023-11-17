package buf

type Position struct {
	count, dataCount, indexCount int
}

func (vbs *Buffers[D, B]) Position() Position {
	pos := Position{
		count:      vbs.count,
		dataCount:  0,
		indexCount: 0,
	}
	if vbs.count > 0 {
		current := vbs.buffers[vbs.count-1]
		pos.dataCount = current.DataCount()
		pos.indexCount = current.IndexCount()
	}
	return pos
}

func (vbs *Buffers[D, B]) Reset(pos Position) {
	// Reset all buffers between last and the one after the reset
	for i := pos.count; i < vbs.count; i++ {
		vbs.init(&vbs.buffers[i], vbs.capacity)
	}
	// Reset the last buffer to the previous data/index counts.
	vbs.count = pos.count
	if pos.count > 0 {
		vbs.reset(&vbs.buffers[pos.count-1], pos)
	}
}

func (vb *Buffer[D]) Reset(pos Position) {
	vb.dataCount = pos.dataCount
	vb.dataReserved = 0
	vb.indexCount = pos.indexCount
	vb.indexReserved = 0
}
