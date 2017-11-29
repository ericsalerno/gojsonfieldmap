package gojsonfieldmap

import (
	"fmt"
	"testing"
)

type simpleObject struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type complexRoot struct {
	SomeValue string       `json:"Thing"`
	Children  simpleObject `json:"Children"`
}

type slicedChild struct {
	StringSlice []string       `json:"stringSlice"`
	ObjectSlice []simpleObject `json:"objectSlice"`
	IntSlice    []int          `json:"intSlice"`
}

type untaggedChild struct {
	TaggedItem   string `json:"tagged"`
	UntaggedItem string
}

type emptyStruct struct {
}

func TestStandardObject(t *testing.T) {
	object := []simpleObject{}

	output := GetJSONObjectFieldMap(object)
	fmt.Println("Standard: " + output)

	if output != `{"Name":1,"Value":1}` {
		t.Fatal("Invalid mapping retured for simple object!")
	}
}

func TestComplexObject(t *testing.T) {
	object := complexRoot{}

	output := GetJSONObjectFieldMap(object)
	fmt.Println("Complex: " + output)

	if output != `{"Thing":1,"Children":{"Name":1,"Value":1}}` {
		t.Fatal("Invalid mapping retured for complex object!")
	}
}

func TestSlicedChildren(t *testing.T) {
	object := slicedChild{}

	output := GetJSONObjectFieldMap(object)
	fmt.Println("Slice Children: " + output)

	if output != `{"stringSlice":1,"objectSlice":{"Name":1,"Value":1},"intSlice":1}` {
		t.Fatal("Invalid mapping retured for slice children object!")
	}
}

func TestUntaggedChildren(t *testing.T) {
	object := untaggedChild{}

	output := GetJSONObjectFieldMap(object)
	fmt.Println("Untagged Children: " + output)

	if output != `{"tagged":1}` {
		t.Fatal("Invalid mapping retured for untagged children object!")
	}
}

func TestEmptyObject(t *testing.T) {
	object := emptyStruct{}

	output := GetJSONObjectFieldMap(object)
	fmt.Println("Untagged Children: " + output)

	if output != `{}` {
		t.Fatal("Invalid mapping retured for untagged children object!")
	}
}

func TestNil(t *testing.T) {
	output := GetJSONObjectFieldMap(nil)

	if output != `{}` {
		t.Fatal("Invalid nil response.")
	}
}
