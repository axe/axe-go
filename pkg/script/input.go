package script

type Rune struct {
	Value rune
	Position
}

func (r *Rune) Rune() rune {
	if r == nil {
		return 0
	}
	return r.Value
}

type RuneBuffer struct {
	buffer Buffer[rune]
	pos    Position
	max    Position
}

func NewRuneBuffer(data string) *RuneBuffer {
	r := &RuneBuffer{
		buffer: NewBuffer[rune](len(data)),
	}
	r.buffer.Set([]rune(data))
	return r
}

func (b *RuneBuffer) Set(data string) {
	b.buffer.Set([]rune(data))
	b.pos = Position{}
	b.max = Position{}
}

func (b RuneBuffer) Len() int {
	return b.buffer.n
}

func (b *RuneBuffer) Read() (r *Rune) {
	r = b.Peek()
	if r != nil {
		b.buffer.i++
		b.pos.Index = b.buffer.i
		if r.Value == '\n' {
			b.pos.Column = 0
			b.pos.Line++
		} else {
			b.pos.Column++
		}
		if b.pos.Index > b.max.Index {
			b.max = b.pos
		}
	}
	return
}

func (b RuneBuffer) Peek() (r *Rune) {
	p := b.buffer.Peek()
	if p != nil {
		r = &Rune{
			Value:    *p,
			Position: b.pos,
		}
	}
	return
}

func (b RuneBuffer) Pos() Position {
	return b.pos
}

func (b RuneBuffer) Max() Position {
	return b.max
}

func (b RuneBuffer) Get(start, end int) string {
	return string(b.GetRunes(start, end))
}

func (b RuneBuffer) GetRunes(start, end int) []rune {
	return b.buffer.data[start:end]
}

func (b *RuneBuffer) Reset(pos Position) {
	b.buffer.i = pos.Index
	b.pos = pos
}

func (b *RuneBuffer) Ended() bool {
	return b.buffer.Ended()
}

// The parsed output
type Input struct {
	Matchers []MatcherMatch
	Runes    *RuneBuffer
	cycles   Set[cycleState]
}

// A state that should not repeat itself, this is to avoid infinite loops
type cycleState struct {
	matcher    int
	reset      int
	greediness int
}
