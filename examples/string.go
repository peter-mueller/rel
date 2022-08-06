package examples

import (
	"fmt"
	"strings"
)

func Insert(p Person) (stmt string, params []any) {
	stmt += fmt.Sprintf("INSERT INTO person (%s)\n", strings.Join(person.AttributeRef(), ", "))
	stmt += fmt.Sprintf("VALUES (%s)\n", strings.Join(person.Params("?"), ","))
	params = person.Tuple(p)
	return stmt, params
}
