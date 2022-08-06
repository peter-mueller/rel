package sqlwriter_test

import (
	"testing"

	"github.com/peter-mueller/rel"
	"github.com/peter-mueller/rel/sqlwriter"
)

type Person struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

var person = rel.NewVar[Person]("person")

func TestInsert(t *testing.T) {
	p := Person{
		Name: "Franz",
		Age:  23,
	}

	var w sqlwriter.PostgresSQLWriter
	w.InsertInto(person.Name(), person.AttributeRef())
	w.Values(person.Tuple(p))
}

func TestSelect(t *testing.T) {
	var p Person

	var w sqlwriter.PostgresSQLWriter
	w.SelectFrom(person.AttributeRef(), person.Dests(&p), person.Name())
	w.Where("age > ?", 3)
}
