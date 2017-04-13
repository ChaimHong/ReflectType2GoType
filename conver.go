package rtype2gtype

import (
	"go/token"
	"go/types"
	"reflect"
)

func Conver(rType reflect.Type) types.Type {
	switch rType.Kind() {
	case reflect.Bool:
		return types.Typ[types.Bool]
	case reflect.Int:
		return types.Typ[types.Int]
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
		return types.NewArray(Conver(rType.Elem()), int64(rType.Len()))
	case reflect.Slice:
		return types.NewSlice(Conver(rType.Elem()))
	case reflect.Struct:
		fields := []*types.Var{}
		for i := 0; i < rType.NumField(); i++ {
			fields = append(fields, types.NewField((token.Pos)(i), nil, rType.Field(i).Name, Conver(rType.Field(i).Type), false))
		}

		return types.NewStruct(fields, nil)
	case reflect.Ptr:
		return types.NewPointer(Conver(rType.Elem()))
	default:
		panic("do not support this reflect type conver")
	}

	return nil
}
