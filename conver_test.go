package rtype2gtype

import (
	"encoding/json"
	"fmt"
	"go/types"
	"log"
	"reflect"
	"testing"

	"github.com/ChaimHong/ReflectType2GoType/example"
	"github.com/ChaimHong/gobuf/parser"
	"github.com/ChaimHong/util"
	"github.com/funny/debug"
)

type B struct {
	f1 int
}

type CINT int

const (
	CINT_A CINT = 0
	CINT_B CINT = 1

	INT_A int = 0
	INT_B int = 1
)

type LoginGroup int

const (
	LOGIN_GROUP_D LoginGroup = 1 // 项目d
	LOGIN_GROUP_E LoginGroup = 2
	LOGIN_GROUP_F LoginGroup = 3
	LOGIN_GROUP_G LoginGroup = 4
)

// 玩家登录参数
type LoginIn struct {
	// User []byte // 平台帐号，最大长度100字符
	// Group LoginGroup // 组别
	Other example.A
}

type A struct {
	a4 CINT
}

func TestConver(t *testing.T) {
	v, _ := NewConver().Conver(reflect.TypeOf(A{}))
	// v1 := c.Conver(reflect.TypeOf(A{}), false)
	// v := c.GetTypes(v1, "A")
	fmt.Printf("%s", debug.Dump(debug.DumpStyle{Format: true, Indent: " "}, "TestConver", v))
}

func TestParser(t *testing.T) {
	// var a int = 1
	// rtype := reflect.TypeOf(a)
	// fmt.Printf("t %v %v", rtype.String(), rtype.Kind().String())
	// return
	typ, _ := NewConver().Conver(reflect.TypeOf(LoginIn{}))
	v, _ := typ.(*types.Struct)

	log.Printf("doc dump %s\n", debug.Dump(debug.DumpStyle{Format: true, Indent: " "}, v))

	doc, err := parser.ParseData("main", nil, map[string]*types.Struct{"A": v}, []string{})
	if err != nil {
		panic(err)
	}

	// log.Printf("doc dump %s", debug.Dump(debug.DumpStyle{Format: true, Indent: " "}, doc))

	jsonData, err2 := json.MarshalIndent(doc, "", " ")

	log.Printf("json %s %v", jsonData, err2)

}

func TestConst(t *testing.T) {
	fmt.Printf("%v", reflect.ValueOf(CINT_B).Int())

	v := ConstConver(reflect.ValueOf(CINT_B))
	fmt.Printf("const %v", v)

	doc, err := parser.ParseData("main", []*types.Const{v}, nil, []string{})
	if err != nil {
		panic(err)
	}

	log.Printf("doc dump %s", debug.Dump(debug.DumpStyle{Format: true, Indent: " "}, doc))
	return
}

type Raw []byte

type Named struct {
	A *Raw
}

func TestNamed(t *testing.T) {
	{
		var v []byte
		typ, _ := NewConver().Conver(reflect.TypeOf(v))

		util.DebugPrintf("t", typ.Underlying(), typ.String())

	}

	{
		v := Raw{}
		typ, _ := NewConver().Conver(reflect.TypeOf(v))

		util.DebugPrintf("t", typ.Underlying(), typ.String())
		return
	}
	{
		typ := types.NewTypeName(0, nil, "Raw", types.NewSlice(types.Typ[types.Byte]))
		util.DebugPrintf("t", typ.Name(), typ.IsAlias())

		typ2 := types.NewNamed(typ, types.NewSlice(types.Typ[types.Byte]), nil)
		util.DebugPrintf("typ2", typ2.String(), typ2.Underlying())
	}

	// {
	// 	typ, _ := NewConver().Conver(reflect.TypeOf(Named{}))
	// 	log.Printf("t %#v", typ)
	// }
}
