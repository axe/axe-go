package scripts

import (
	"fmt"
	"reflect"
	"strconv"
)

func refOf(value any) Ref {
	val := concrete(value)
	return Ref{
		Value: val,
		Type:  val.Type(),
		Path:  []string{},
	}
}

type Ref struct {
	Value  reflect.Value
	Type   reflect.Type
	Path   []string
	Parent *Ref
	Error  error
	Exists bool
}

func (ref Ref) Valid() bool {
	return ref.Path != nil && ref.Error == nil && (ref.Value.IsValid() || ref.Type == anyType)
}
func (ref Ref) Clone() Ref {
	copy := ref
	copy.Value = clone(ref.Value)
	return copy
}
func (ref *Ref) GetError() error {
	if ref == nil {
		return INVALID_REF
	}
	if !ref.Valid() {
		if ref.Error != nil {
			return ref.Error
		} else {
			return INVALID_REF
		}
	}
	return nil
}
func (ref *Ref) Ref(next any) *Ref {
	node, err := toString(next)
	nodeValue := reflect.ValueOf(nil)
	nodeType := reflect.TypeOf(nil)
	refType := ref.Type
	exists := false

	switch refType.Kind() {
	case reflect.Map:
		key, keyErr := fromString(node, refType.Key())
		if keyErr == nil {
			if !ref.Value.IsNil() {
				keyValue := ref.Value.MapIndex(key)
				if keyValue.IsValid() {
					nodeValue = initValue(keyValue)
					exists = true
				}
			}
			if !exists {
				nodeValue = toWritable(valueOf(refType.Elem()))
			}
			nodeType = refType.Elem()
		} else {
			err = keyErr
		}
	case reflect.Struct:
		fieldValue := ref.Value.FieldByName(node)
		if fieldValue.IsValid() {
			nodeValue = initValue(fieldValue)
			nodeType = nodeValue.Type()
			exists = true
		}
	case reflect.Slice:
	case reflect.Array:
		index, indexErr := strconv.Atoi(node)
		if indexErr == nil {
			if ref.Value.IsNil() || index >= ref.Value.Len() {
				nodeValue = valueOf(refType.Elem())
				exists = true
			} else {
				nodeValue = toWritable(ref.Value.Index(index))
			}
			nodeType = refType.Elem()
		} else {
			err = indexErr
		}
	}

	return &Ref{
		Value:  nodeValue,
		Type:   nodeType,
		Path:   append(ref.Path[:], node),
		Parent: ref,
		Error:  err,
		Exists: exists,
	}
}
func (ref *Ref) Set(value any) error {
	if !ref.Value.IsValid() {
		ref.Value = valueOf(ref.Type)
	}
	converted := toType(value, ref.Type)
	if !converted.IsValid() {
		return fmt.Errorf("Value of type %v is not compatible with %v", reflectValue(value).Type(), ref.Type)
	}
	if !ref.Value.CanSet() {
		return fmt.Errorf("The value at %v is not writable, try passing in references", ref.Path)
	}
	ref.Value.Set(converted)
	return ref.Save()
}
func (ref Ref) Concrete() {
	refType := ref.Type
	switch refType.Kind() {
	case reflect.Map:
	case reflect.Slice:
	case reflect.Array:
	case reflect.Pointer:
		if ref.Value.IsNil() {
			ref.Value.Set(valueOf(refType))
		}
	}
}
func (ref Ref) Save() error {
	if ref.Parent == nil {
		return nil
	}

	node := ref.Path[len(ref.Path)-1]
	refType := ref.Parent.Type

	switch refType.Kind() {
	case reflect.Map:
		key, keyErr := fromString(node, refType.Key())
		if keyErr != nil {
			return keyErr
		}
		ref.Parent.Concrete()
		ref.Parent.Value.SetMapIndex(key, toType(ref.Value, refType.Elem()))
	case reflect.Slice:
	case reflect.Array:
		index, indexErr := strconv.Atoi(node)
		if indexErr != nil {
			return indexErr
		}
		ref.Parent.Concrete()
		len := ref.Parent.Value.Len()
		if index >= len {
			if refType.Kind() == reflect.Array {
				return INDEX_OUTSIDE_ARRAY
			} else {
				for i := len; i <= index; i++ {
					ref.Parent.Value = reflect.AppendSlice(ref.Parent.Value, valueOf(refType.Elem()))
				}
			}
		}
		ref.Parent.Value.Index(index).Set(toType(ref.Value, refType.Elem()))
	}

	return ref.Parent.Save()
}
