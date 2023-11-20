package rtype2gtype

import (
	"go/constant"
	"go/types"
	"reflect"
)

func ConstConver(rValue reflect.Value) *types.Const {
	switch rValue.Kind() {
	case reflect.Bool:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Bool], constant.MakeBool(rValue.Bool()))
	case reflect.Int:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Int], constant.MakeInt64(rValue.Int()))
	case reflect.Int16:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Int16], constant.MakeInt64(rValue.Int()))
	case reflect.Int32:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Int32], constant.MakeInt64(rValue.Int()))
	case reflect.Int64:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Int64], constant.MakeInt64(rValue.Int()))
	case reflect.Uint:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Uint], constant.MakeUint64(rValue.Uint()))
	case reflect.Uint8:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Uint8], constant.MakeUint64(rValue.Uint()))
	case reflect.Uint16:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Uint16], constant.MakeUint64(rValue.Uint()))
	case reflect.Uint32:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Uint32], constant.MakeUint64(rValue.Uint()))
	case reflect.Uint64:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Uint64], constant.MakeUint64(rValue.Uint()))
	case reflect.Float32:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Float32], constant.MakeFloat64(rValue.Float()))
	case reflect.Float64:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.Float64], constant.MakeFloat64(rValue.Float()))
	case reflect.String:
		return types.NewConst(0, nil, rValue.Type().Name(), types.Typ[types.String], constant.MakeString(rValue.String()))

	default:
		panic("do not support this reflect value kind")
	}
}
