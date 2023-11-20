package rtype2gtype

import (
	"fmt"
	"go/types"
	"log"
	"reflect"
	"testing"

	example "github.com/ChaimHong/ReflectType2GoType/example"
	"github.com/ChaimHong/gobuf/parser"
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
	Other example.A
}

type A struct {
	a4 CINT
}

func TestConver(t *testing.T) {
	v, _ := NewConver().Conver(reflect.TypeOf(A{}))
	fmt.Printf("%s", debug.Dump(debug.DumpStyle{Format: true, Indent: " "}, "TestConver", v))
}

func TestConst(t *testing.T) {
	fmt.Printf("%v", reflect.ValueOf(CINT_B).Int())

	v := ConstConver(reflect.ValueOf(CINT_B))
	fmt.Printf("const %v", v)

	doc, err := parser.ParseData("main", []*types.Const{v}, nil, nil)
	if err != nil {
		panic(err)
	}

	log.Printf("doc dump %s", debug.Dump(debug.DumpStyle{Format: true, Indent: " "}, doc))
	return
}
