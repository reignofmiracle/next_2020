package intset_test

import (
	"fmt"
	"testing"

	"../intset"
)

func TestAdd(t *testing.T) {
	var testObj intset.IntSet
	testObj.Add(10)
	fmt.Println(testObj.String())
}

func TestUnionWith(t *testing.T) {
	var testObj1 intset.IntSet
	testObj1.Add(10)
	var testObj2 intset.IntSet
	testObj2.Add(33)
	testObj1.UnionWith(&testObj2)
	fmt.Println(testObj1.String())
}

func TestHas(t *testing.T) {
	var testObj intset.IntSet
	testObj.Add(10)
	if !testObj.Has(10) {
		t.Error()
	}
	fmt.Println(testObj.String())
}
