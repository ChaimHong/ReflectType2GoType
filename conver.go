package rtype2gtype

import (
	"go/token"
	"go/types"
	"reflect"
)

type Conver struct {
	Named map[string]types.Type
}

func NewConver() *Conver {
	return &Conver{
		Named: make(map[string]types.Type),
	}
}

func (c *Conver) free() {
	c.Named = make(map[string]types.Type)
}

func (c *Conver) Conver(rType reflect.Type) types.Type {
	defer func() {
		c.free()
	}()

	return c.conver(rType, false)
}

func (c *Conver) conver(rType reflect.Type, named bool) (ret types.Type) {
	rk := rType.Kind()
	switch rType.Kind() {
	case reflect.Bool:
		ret = types.Typ[types.Bool]
		if rType.String() != rk.String() {
			ret = c.addNamed(rType.Name(), ret)
		}
		return
	case reflect.Int:
		ret = types.Typ[types.Int]
		if rType.String() != rk.String() {
			ret = c.addNamed(rType.Name(), ret)
		}
		return

	case reflect.Int8:
		return types.Typ[types.Int8]
	case reflect.Int16:
		return types.Typ[types.Int16]
	case reflect.Int32:
		return types.Typ[types.Int32]
	case reflect.Int64:
		return types.Typ[types.Int64]
	case reflect.Uint:
		return types.Typ[types.Uint]
	case reflect.Uint8:
		return types.Typ[types.Uint8]
	case reflect.Uint16:
		return types.Typ[types.Uint16]
	case reflect.Uint32:
		return types.Typ[types.Uint32]
	case reflect.Uint64:
		return types.Typ[types.Uint64]
	case reflect.Uintptr:
		return types.Typ[types.Uintptr]
	case reflect.Float32:
		return types.Typ[types.Float32]
	case reflect.Float64:
		return types.Typ[types.Float64]
	case reflect.Complex64:
		return types.Typ[types.Complex64]
	case reflect.Complex128:
		return types.Typ[types.Complex128]
	case reflect.String:
		return types.Typ[types.String]
	case reflect.UnsafePointer:
		return types.Typ[types.UnsafePointer]

	case reflect.Array:
		if elem := rType.Elem(); elem.Kind() == reflect.Struct {
			if elemNamed, exist := c.Named[elem.Name()]; exist {
				return types.NewArray(elemNamed, int64(rType.Len()))
			}
		}

		return types.NewArray(c.conver(rType.Elem(), true), int64(rType.Len()))
	case reflect.Slice:
		if elem := rType.Elem(); elem.Kind() == reflect.Struct {
			if elemNamed, exist := c.Named[elem.Name()]; exist {
				return types.NewSlice(elemNamed)
			}
		}

		return types.NewSlice(c.conver(rType.Elem(), true))
	case reflect.Struct:
		fields := []*types.Var{}
		for i := 0; i < rType.NumField(); i++ {
			fields = append(fields, types.NewField((token.Pos)(i), nil, rType.Field(i).Name, c.conver(rType.Field(i).Type, true), false))
		}

		var ret types.Type
		ret = types.NewStruct(fields, nil)

		if named {
			ret = c.addNamed(rType.Name(), ret)
		}

		return ret
	case reflect.Ptr:
		if elem := rType.Elem(); elem.Kind() == reflect.Struct {
			if elemNamed, exist := c.Named[elem.Name()]; exist {
				return types.NewPointer(elemNamed)
			}
		}

		return types.NewPointer(c.conver(rType.Elem(), true))
	default:
		panic("do not support this reflect type conver")
	}

	return nil
}

func (c *Conver) addNamed(name string, t types.Type) types.Type {
	ret := types.NewNamed(types.NewTypeName(0, nil, name, t), t, nil)
	c.Named[name] = ret
	return ret
}
