package scripts

import (
	"encoding"
	"fmt"
	"reflect"
	"strconv"
)

func toBool(value any) (bool, error) {
	if value == nil {
		return false, nil
	}
	switch v := value.(type) {
	case float32:
		return v != 0, nil
	case float64:
		return v != 0, nil
	case bool:
		return v, nil
	case int:
		return v != 0, nil
	case int8:
		return v != 0, nil
	case int16:
		return v != 0, nil
	case int32:
		return v != 0, nil
	case int64:
		return v != 0, nil
	case uint:
		return v != 0, nil
	case uint8:
		return v != 0, nil
	case uint16:
		return v != 0, nil
	case uint32:
		return v != 0, nil
	case uint64:
		return v != 0, nil
	case uintptr:
		return v != 0, nil
	case string:
		return v != "", nil
	}
	r := reflect.ValueOf(value)
	switch r.Kind() {
	case reflect.Slice:
	case reflect.Map:
		return !r.IsNil() && r.Len() > 0, nil
	case reflect.Pointer:
		return toBool(r.Elem())
	}
	return false, INVALID_BOOL_TYPE
}

func toFloat(value any) (float64, error) {
	switch v := value.(type) {
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case bool:
		if v {
			return 1.0, nil
		} else {
			return 0.0, nil
		}
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case uintptr:
		return float64(v), nil
	}
	valueString, err := toString(value)
	if err != nil {
		valueParsed, err := strconv.ParseFloat(valueString, 64)
		if err != nil {
			return valueParsed, nil
		}
	}
	return 0, INVALID_COMPARE_TYPE
}

func toInt(value any) (int64, error) {
	switch v := value.(type) {
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case bool:
		if v {
			return 1, nil
		} else {
			return 0, nil
		}
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case uintptr:
		return int64(v), nil
	}
	valueString, err := toString(value)
	if err != nil {
		valueParsed, err := strconv.ParseInt(valueString, 10, 64)
		if err != nil {
			return valueParsed, nil
		}
	}
	return 0, INVALID_BINARY_TYPE
}

func toString(value any) (string, error) {
	if marshal, ok := value.(encoding.TextMarshaler); ok {
		text, err := marshal.MarshalText()
		if err != nil {
			return "", err
		}
		return string(text), nil
	}
	base := concrete(value)
	if !base.IsValid() {
		return "nil", nil
	}
	return fmt.Sprintf("%+v", base.Interface()), nil
}

func fromString(x string, target reflect.Type) (reflect.Value, error) {
	value := valueOf(target)
	if unmarshal, ok := value.Interface().(encoding.TextUnmarshaler); ok {
		return reflect.ValueOf(unmarshal), unmarshal.UnmarshalText([]byte(x))
	}
	if !value.CanSet() {
		return value, nil
	}
	parsed, err := parseString(x, target)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	switch v := parsed.(type) {
	case string:
		value.SetString(v)
	case float64:
		value.SetFloat(v)
	case int64:
		value.SetInt(v)
	case uint64:
		value.SetUint(v)
	case bool:
		value.SetBool(v)
	default:
		value.Set(reflect.ValueOf(parsed))
	}
	return value, nil
}

func parseString(x string, concreteType reflect.Type) (any, error) {
	switch concreteType.Kind() {
	case reflect.String:
		return x, nil
	case reflect.Float32:
		return strconv.ParseFloat(x, 32)
	case reflect.Float64:
		return strconv.ParseFloat(x, 64)
	case reflect.Bool:
		return strconv.ParseBool(x)
	case reflect.Int:
		return strconv.ParseInt(x, 10, 32)
	case reflect.Int8:
		return strconv.ParseInt(x, 10, 8)
	case reflect.Int16:
		return strconv.ParseInt(x, 10, 16)
	case reflect.Int32:
		return strconv.ParseInt(x, 10, 32)
	case reflect.Int64:
		return strconv.ParseInt(x, 10, 64)
	case reflect.Uint:
		return strconv.ParseUint(x, 10, 32)
	case reflect.Uint8:
		return strconv.ParseUint(x, 10, 8)
	case reflect.Uint16:
		return strconv.ParseUint(x, 10, 16)
	case reflect.Uint32:
		return strconv.ParseUint(x, 10, 32)
	case reflect.Uint64:
		return strconv.ParseUint(x, 10, 64)
	case reflect.Uintptr:
		return strconv.ParseUint(x, 10, 64)
	}
	return nil, nil
}

var anyType = reflect.TypeOf([]any{}).Elem()

func reflectValue(value any) reflect.Value {
	if rv, ok := value.(reflect.Value); ok {
		return rv
	}
	return reflect.ValueOf(value)
}

func valueOf(typ reflect.Type) reflect.Value {
	return initValue(reflect.New(typ).Elem())
}

func ptrTo(value any) reflect.Value {
	val := reflectValue(value)
	ptr := reflect.New(val.Type())
	ptr.Elem().Set(val)
	return ptr
}

func toWritable(value any) reflect.Value {
	val := reflectValue(value)
	if val.Kind() != reflect.Interface {
		val = ptrTo(val)
	}
	return val
}

func toType(value any, target reflect.Type) reflect.Value {
	val := concrete(value)
	concreteType := val.Type()
	for target.Kind() == reflect.Pointer {
		val = ptrTo(val)
		target = target.Elem()
	}
	if target.Kind() == reflect.Interface {
		return ptrTo(any(val.Interface()))
	}
	if target != concreteType {
		return reflect.Value{}
	}
	return val
}

func concrete(value any) reflect.Value {
	rv := reflectValue(value)
	for rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	return rv
}

func initValue(value any) reflect.Value {
	val := reflectValue(value)
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		val = val.Elem()
	}
	return val
}

func clone(value any) reflect.Value {
	val := reflectValue(value)
	cop := ptrTo(reflectValue(value)).Elem()
	switch val.Kind() {
	case reflect.Pointer:
		cop.Set(clone(val.Elem()))
	case reflect.Slice:
		cop = reflect.MakeSlice(val.Type().Elem(), val.Len(), val.Cap())
		for i := 0; i < val.Len(); i++ {
			cop.Index(i).Set(clone(val.Index(i)))
		}
	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			cop.Index(i).Set(clone(val.Index(i)))
		}
	case reflect.Map:
		cop = reflect.MakeMap(val.Type())
		iter := val.MapRange()
		for iter.Next() {
			cop.SetMapIndex(iter.Key(), clone(iter.Value()))
		}
	}
	return cop
}

func sliceMap[S any, D any](slice []S, mapper func(source S) D) []D {
	dest := make([]D, len(slice), cap(slice))
	if slice != nil {
		for i := range slice {
			dest[i] = mapper(slice[i])
		}
	}
	return dest
}

func mapMap[K comparable, S any, D any](source map[K]S, mapper func(source S) D) map[K]D {
	dest := make(map[K]D)
	if source != nil {
		for key := range source {
			dest[key] = mapper(source[key])
		}
	}
	return dest
}
