package rtype2gtype

import (
	"encoding/json"
	"fmt"
	"github.com/ChaimHong/ReflectType2GoType/example"
	"github.com/ChaimHong/gobuf/parser"
	"github.com/funny/debug"
	"go/types"
	"log"
	"reflect"
	"testing"
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
	v, _ := NewConver().Conver(reflect.TypeOf(LoginIn{})).(*types.Struct)

	return
	log.Printf("doc dump %s\n", debug.Dump(debug.DumpStyle{Format: true, Indent: " "}, v))

	doc, err := parser.ParseData("main", nil, map[string]*types.Struct{"A": v})
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

	doc, err := parser.ParseData("main", []*types.Const{v}, nil)
	if err != nil {
		panic(err)
	}

	log.Printf("doc dump %s", debug.Dump(debug.DumpStyle{Format: true, Indent: " "}, doc))
	return
}
