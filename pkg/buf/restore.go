package buf

type Position struct {
	current, dataCount, indexCount int
}

func (vbs *Buffers[D, B]) Position() Position {
	pos := Position{
		current:    vbs.current,
		dataCount:  0,
		indexCount: 0,
	}
	if vbs.current >= 0 {
		current := vbs.buffers[vbs.current]
		pos.dataCount = current.DataCount()
		pos.indexCount = current.IndexCount()
	}
	return pos
}

func (vbs *Buffers[D, B]) Reset(pos Position) {
	vbs.current = pos.current
	if vbs.current >= 0 {
		vbs.reset(&vbs.buffers[vbs.current], pos)
	}
}

func (vb *Buffer[D]) Reset(pos Position) {
	vb.dataCount = pos.dataCount
	vb.indexCount = pos.indexCount
}
