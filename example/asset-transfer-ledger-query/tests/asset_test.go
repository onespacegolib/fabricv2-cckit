package tests

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/onespacegolib/fabricv2-cckit/flannel"
	"github.com/onespacegolib/fabricv2-cckit/helper"
	"github.com/onespacegolib/fabricv2-cckit/testcc"
	"testing"
)

const ASSET = `asset`

func TestCreateAsset(t *testing.T) {
	res = cc.Invoke(testcc.MakeFunction(ASSET, `create`), "asset1", "blue", "5", "tom", "35")
	helper.PlayRes(res)
	println(res.Message)
	//res = cc.Query(testcc.MakeFunction(ASSET, `get[id]`), "asset1")
	//helper.PlayRes(res)
	//res = cc.Invoke(testcc.MakeFunction(ASSET, `put`), "asset1")
	//helper.PlayRes(res)
	//res = cc.Query(testcc.MakeFunction(ASSET, `get[id]`), "asset1")
	//helper.PlayRes(res)
}

func TestAsset(t *testing.T) {
	cc.Invoke(testcc.MakeFunction(ASSET, ``))
}

type Server struct {
	Name           string `json:"name,omitempty"`
	ID             int
	Enabled        bool
	users          []string // not exported
	flannel.Schema          // embedded
}

func Test(t *testing.T) {
	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}
	s := structs.New(server)

	// Get the Field struct for the "Name" field

	name := s.Field("Name")
	obj := s.Field("Schema").Field(`ObjectType`)
	id := s.Field("ID")
	id.Set(123)
	obj.Set(`123`)

	// Get the underlying value,  value => "gopher"
	//value := name.Value().(string)
	//println(value)

	// Set the field's value
	name.Set("another gopher")
	value := name.Value().(string)
	println(value)

	// Get the field's kind, kind =>  "string"
	name.Kind()

	// Check if the field is exported or not
	if name.IsExported() {
		fmt.Println("Name field is exported")
	}

	// Check if the value is a zero value, such as "" for string, 0 for int
	if !name.IsZero() {
		fmt.Println("Name is initialized")
	}

	// Check if the field is an anonymous (embedded) field
	if !obj.IsEmbedded() {
		fmt.Println("Schema is not an embedded field")
	}

	// Get the Field's tag value for tag name "json", tag value => "name,omitempty"
	tagValue := name.Tag("json")
	println(tagValue)

	helper.PrintStruct(server)
}
