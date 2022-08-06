package examples

import "github.com/peter-mueller/rel"

type Person struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

var person = rel.NewVar[Person]()
