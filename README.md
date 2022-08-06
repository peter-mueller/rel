# rel
Go database helper that maps structs to tuples borrowing method names from Tutorial D syntax

*Usage example:*
```go
import "github.com/peter-mueller/rel"

type Person struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// TODO: maybe include name in relation: rel.NewVar[Person]("person")
var person = rel.NewVar[Person]()

func Insert(p Person) (stmt string, params []any) {
	stmt += fmt.Sprintf("INSERT INTO person (%s)\n", strings.Join(person.AttributeRef(), ", "))
	stmt += fmt.Sprintf("VALUES (%s)\n", strings.Join(person.Params("?"), ","))
	params = person.Tuple(p)
	return stmt, params
}
// OR generic variant
func Insert[T any](v *rel.Var[T], relName string,  value T) (stmt string, params []any) {
	stmt += fmt.Sprintf("INSERT INTO %s (%s)\n", relName, strings.Join(v.AttributeRef(), ", "))
	stmt += fmt.Sprintf("VALUES (%s)\n", strings.Join(v.Params("?"), ","))
	params = v.Tuple(value)
	return stmt, params
}


p := Person{Name: "Hans", Age: 33}
stmt, params := Insert(p)
// INSERT INTO person (name, age) 
// VALUES (?,?)
// ["Hans", 33]
res, err := db.ExecContext(ctx, stmt, params...)
```
