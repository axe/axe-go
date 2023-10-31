package state

import (
	"encoding/binary"
	"math"
)

type UserDataPropertyBase interface {
	Bytes() int
	Write(data UserData)
	SetIndex(index int)
}

type UserDataPropertyTyped[V any] interface {
	UserDataPropertyBase
	Get(data UserData) V
	Set(data UserData, value V)
}

func Bool(defaultValue bool) UserDataPropertyTyped[bool] {
	return &userDataPropertyBool{defaultValue: defaultValue}
}

type userDataPropertyBool struct {
	index        int
	defaultValue bool
}

func (p userDataPropertyBool) Get(data UserData) bool {
	offset := data.offsets[p.index]
	return data.data[offset] != 0
}

func (p userDataPropertyBool) Set(data UserData, value bool) {
	offset := data.offsets[p.index]
	if value {
		data.data[offset] = 1.0
	} else {
		data.data[offset] = 0.0
	}
}

func (p userDataPropertyBool) Bytes() int {
	return 1
}

func (p userDataPropertyBool) Write(data UserData) {
	p.Set(data, p.defaultValue)
}

func (p *userDataPropertyBool) SetIndex(index int) {
	p.index = index
}

func Float(defaultValue float32) UserDataPropertyTyped[float32] {
	return &userDataPropertyFloat{defaultValue: defaultValue}
}

type userDataPropertyFloat struct {
	index        int
	defaultValue float32
}

func (p userDataPropertyFloat) Get(data UserData) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(data.at(p.index)))
}

func (p userDataPropertyFloat) Set(data UserData, value float32) {
	binary.BigEndian.PutUint32(data.at(p.index), math.Float32bits(value))
}

func (p userDataPropertyFloat) Bytes() int {
	return 4
}

func (p userDataPropertyFloat) Write(data UserData) {
	p.Set(data, p.defaultValue)
}

func (p *userDataPropertyFloat) SetIndex(index int) {
	p.index = index
}

type UserData struct {
	data    []byte
	offsets []int
}

func NewUserData(props ...UserDataPropertyBase) UserData {
	offset := 0
	userData := UserData{
		offsets: make([]int, len(props)),
	}
	for i, p := range props {
		size := p.Bytes()
		userData.offsets[i] = offset
		offset += size
	}
	userData.data = make([]byte, offset)
	for i, p := range props {
		p.SetIndex(i)
		p.Write(userData)
	}
	return userData
}

func (ud UserData) at(prop int) []byte {
	return ud.data[ud.offsets[prop]:]
}
