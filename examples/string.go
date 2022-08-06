package examples

import (
	"fmt"
	"strings"

	"github.com/peter-mueller/rel"
)

func Insert(p Person) (stmt string, params []any) {
	stmt += fmt.Sprintf("INSERT INTO person (%s)\n", strings.Join(person.AttributeRef(), ", "))
	stmt += fmt.Sprintf("VALUES (%s)\n", strings.Join(person.Params("?"), ","))
	params = person.Tuple(p)
	return stmt, params
}

func InsertGeneric[T any](v *rel.Var[T], relName string, value T) (stmt string, params []any) {
	stmt += fmt.Sprintf("INSERT INTO %s (%s)\n", relName, strings.Join(v.AttributeRef(), ", "))
	stmt += fmt.Sprintf("VALUES (%s)\n", strings.Join(v.Params("?"), ","))
	params = v.Tuple(value)
	return stmt, params
}
