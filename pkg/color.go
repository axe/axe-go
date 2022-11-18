package axe

type Colori = Color[int]
type Colorf = Color[float32]

type Color[C any] struct {
	R, G, B, A C
}
