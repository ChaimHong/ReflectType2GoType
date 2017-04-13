package rtype2gtype

import (
	"fmt"
	"github.com/funny/debug"
	"reflect"
	"testing"
)

type B struct {
	f1 int
	f2 bool
	f3 int32
}

type A struct {
	a1 *B
	a2 []int
	a3 []byte
	a4 string
	a5 byte
}

func TestConver(t *testing.T) {
	v := Conver(reflect.TypeOf(A{}))
	fmt.Printf("%s", debug.Dump(debug.DumpStyle{Format: true, Indent: " "}, "TestConver", v))
}
