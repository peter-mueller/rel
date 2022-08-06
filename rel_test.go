package rel_test

import (
	"testing"

	"github.com/peter-mueller/rel"
)

type Person struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

var person = rel.NewVar[Person]("person")

func TestTuple(t *testing.T) {
	p := Person{
		Name: "Franz",
		Age:  21,
	}
	expected := []any{"Franz", 21}
	tuple := person.Tuple(p)
	if len(expected) != len(tuple) {
		t.Fail()
	}
	for i := range expected {
		if expected[i] != tuple[i] {
			t.Fail()
		}
	}
	if t.Failed() {
		t.Fatalf("tuple %v did not match %#v", tuple, p)
	}
}

func TestAttributeRef(t *testing.T) {
	expected := []any{"name", "age"}
	tuple := person.AttributeRef()
	if len(expected) != len(tuple) {
		t.Fail()
	}
	for i := range expected {
		if expected[i] != tuple[i] {
			t.Fail()
		}
	}
	if t.Failed() {
		t.Fatalf("attributeref %v did not match struct %#v", tuple, Person{})
	}
}

func TestDests(t *testing.T) {
	p := Person{}
	expected := []any{&p.Name, &p.Age}
	dests := person.Dests(&p)
	if len(expected) != len(dests) {
		t.Fail()
	}
	for i := range expected {
		if expected[i] != dests[i] {
			t.Fail()
		}
	}
	if t.Failed() {
		t.Fatalf("attributeref %v did not match struct %#v", dests, Person{})
	}
}

func TestParams(t *testing.T) {
	expected := []string{"?", "?"}
	params := person.Params("?")
	if len(expected) != len(params) {
		t.Fail()
	}
	for i := range expected {
		if expected[i] != params[i] {
			t.Fail()
		}
	}
	if t.Failed() {
		t.Fatalf("params %v did not match struct %#v", params, Person{})
	}
}
