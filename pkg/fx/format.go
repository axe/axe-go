package fx

import (
	"github.com/axe/axe-go/pkg/ease"
	"github.com/axe/axe-go/pkg/util"
)

type Format struct {
	Size       int
	Attributes []AttributeFormat
}

func NewFormat() *Format {
	return &Format{Size: 0, Attributes: make([]AttributeFormat, 16)}
}

func (pf *Format) Get(attr Attribute, particle []float32) []float32 {
	if access := pf.Access(attr); access != nil {
		return access.Get(particle, pf)
	}
	return nil
}

func (pf Format) Access(attr Attribute) Access {
	if attr.id < len(pf.Attributes) {
		if attrFormat := pf.Attributes[attr.id]; attrFormat.Access != nil {
			return attrFormat.Access
		}
	}
	return nil
}

func (pf *Format) Add(attr AttributeFormat) *Format {
	pf.Attributes = util.SliceEnsureSize(pf.Attributes, attr.Attribute.id+1)
	pf.Attributes[attr.Attribute.id] = attr
	return pf
}

func (pf *Format) HasData(attr Attribute) bool {
	access := pf.Access(attr)
	if access == nil {
		return false
	}
	_, ok := access.(AccessData)
	return ok
}

func (pf *Format) AddData(attr Attribute) *Format {
	offset := pf.Size
	pf.Size += attr.size
	return pf.Add(AttributeFormat{
		Attribute: attr,
		Access: AccessData{
			Offset: offset,
			Size:   attr.size,
		},
	})
}

func (pf *Format) AddConstant(attr Attribute, constant ...float32) *Format {
	return pf.Add(AttributeFormat{
		Attribute: attr,
		Access:    AccessConstant{Constant: constant},
	})
}

func (pf *Format) AddDynamic(attr Attribute, dynamic Dynamic) *Format {
	return pf.Add(AttributeFormat{
		Attribute: attr,
		Access:    AccessDynamic{Dynamic: dynamic},
	})
}

func (pf *Format) AddLerp(attr Attribute, data [][]float32, easing ease.Easing) *Format {
	return pf.Add(AttributeFormat{
		Attribute: attr,
		Access: AccessLerp{
			temp:   make([]float32, attr.size),
			Data:   data,
			Easing: easing,
		},
	})
}

type Data struct {
	Format *Format
	Data   []float32
}

func NewData(format *Format, maxParticles int) Data {
	return Data{
		Format: format,
		Data:   make([]float32, maxParticles*format.Size),
	}
}

func (pd Data) At(index int) []float32 {
	offset := pd.Location(index)
	return pd.Data[offset : offset+pd.Format.Size]
}

func (pd Data) Location(index int) int {
	return index * pd.Format.Size
}

func (pd Data) Get(index int, attr Attribute) []float32 {
	return pd.Format.Get(attr, pd.At(index))
}
