# rel
Go database helper that maps structs to tuples borrowing method names from Tutorial D syntax

*Usage example:*
```go
import "github.com/peter-mueller/rel"

type Person struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

var person = rel.NewVar[Person]()

func Insert(p Person) (stmt string, params []any) {
	stmt += fmt.Sprintf("INSERT INTO person (%s)\n", strings.Join(person.AttributeRef(), ", "))
	stmt += fmt.Sprintf("VALUES (%s)\n", strings.Join(person.Params("?"), ","))
	params = person.Tuple(p)
	return stmt, params
}


p := Person{Name: "Hans", Age: 33}
stmt, params := Insert(p)
// INSERT INTO person (name, age) 
// VALUES (?,?)
// ["Hans", 33]
res, err := db.ExecContext(ctx, stmt, params...)
```
