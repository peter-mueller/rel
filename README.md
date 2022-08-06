# rel 

> :construction: no stable api, just a playground for go database client alternatives

Go database helper that maps structs to tuples borrowing method names from Tutorial D syntax

*Usage example:*
```go
import "github.com/peter-mueller/rel"

type Person struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

var person = rel.NewVar[Person]("person")

func Insert(p Person) (stmt string, params []any) {
	stmt += fmt.Sprintf("INSERT INTO person (%s)\n", strings.Join(person.AttributeRef(), ", "))
	stmt += fmt.Sprintf("VALUES (%s)\n", strings.Join(person.Params("?"), ","))
	params = person.Tuple(p)
	return stmt, params
}

func InsertGeneric[T any](v *rel.Var[T], value T) (stmt string, params []any) {
	stmt += fmt.Sprintf("INSERT INTO %s (%s)\n", v.Name(), strings.Join(v.AttributeRef(), ", "))
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

*sqlwriter example:*

```go
p := Person{
	Name: "Franz",
	Age:  23,
}

var w sqlwriter.PostgresSQLWriter
w.InsertInto(person.Name(), person.AttributeRef())
w.Values(person.Tuple(p))
res, err := db.ExecContext(ctx, w.Stmt(), w.Params...)
```

```go
var p Person

var w sqlwriter.PostgresSQLWriter
w.SelectFrom(person.AttributeRef(), person.Dests(&p), person.Name())
w.Where("age > ?", 3)
res, err := db.ExecContext(ctx, w.Stmt(), w.Params...)
// Scan(w.Dests...)
```