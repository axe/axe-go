package fx

import "github.com/axe/axe-go/pkg/util"

var (
	attributeID = util.NewIncrementor[int](0, 1)
)

type Attribute struct {
	id     int
	size   int
	init   Init
	modify Modify
}

func (a Attribute) ID() int   { return a.id }
func (a Attribute) Size() int { return a.size }

func NewAttribute(size int) Attribute {
	return Attribute{id: attributeID.Get(), size: size}
}

func (a Attribute) GetInit() Init {
	return a.init
}

func (a Attribute) Init(init Init) Attribute {
	a.init = init
	return a
}

func (a Attribute) GetModify() Modify {
	return a.modify
}

func (a Attribute) Modify(modify Modify) Attribute {
	a.modify = modify
	return a
}

type AttributeFormat struct {
	Access    Access
	Attribute Attribute
}
