package examples

import (
	"testing"

	"github.com/peter-mueller/rel"
)

func TestStringInsert(t *testing.T) {
	p := Person{Name: "Hans", Age: 33}
	stmt, params := Insert(p)
	expected := "INSERT INTO person (name, age)\nVALUES (?,?)\n"
	if stmt != expected {
		t.Fatalf("stmt was\n\n%s\nbut should be\n\n%s", stmt, expected)
	}
	expectedParams := []any{"Hans", 33}
	for i := range expectedParams {
		if expectedParams[i] != params[i] {
			t.Fatalf("params were %v but should be%v", params, expectedParams)
		}
	}
}

func TestStringInsertGeneric(t *testing.T) {
	p := Person{Name: "Hans", Age: 33}
	v := rel.NewVar[Person]()
	stmt, params := InsertGeneric(v, "person", p)
	expected := "INSERT INTO person (name, age)\nVALUES (?,?)\n"
	if stmt != expected {
		t.Fatalf("stmt was\n\n%s\nbut should be\n\n%s", stmt, expected)
	}
	expectedParams := []any{"Hans", 33}
	for i := range expectedParams {
		if expectedParams[i] != params[i] {
			t.Fatalf("params were %v but should be%v", params, expectedParams)
		}
	}
}
