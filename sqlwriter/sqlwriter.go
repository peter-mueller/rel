package sqlwriter

import (
	"fmt"
	"strings"
)

type PostgresSQLWriter struct {
	Builder strings.Builder
	Params  []any
	Dests   []any
}

func (w *PostgresSQLWriter) Stmt() string {
	return w.Builder.String()
}

func (w *PostgresSQLWriter) Linef(format string, a ...any) {
	fmt.Fprintf(&w.Builder, format, a...)
	w.Builder.WriteString("\n")
}

func (w *PostgresSQLWriter) InsertInto(relname string, attributerefs []string) {
	w.Linef(
		"INSERT INTO %s (%s)",
		relname, strings.Join(attributerefs, ", "),
	)
}

func (w *PostgresSQLWriter) Values(tuple []any) {
	params := make([]string, len(tuple))
	for i := range tuple {
		params[i] = "?"
	}
	w.Linef(
		"VALUES (%s)",
		strings.Join(params, ","),
	)

	w.Params = append(w.Params, tuple...)
}

func (w *PostgresSQLWriter) Where(cond string, params ...any) {
	w.Linef(
		"WHERE %s\n",
		cond,
	)
	w.Params = append(w.Params, params...)
}

func (w *PostgresSQLWriter) SelectFrom(attributerefs []string, dests []any, relname string) {
	w.Linef(
		"SELECT %s FROM %s\n",
		strings.Join(attributerefs, ", "), relname,
	)
	w.Dests = append(w.Dests, dests...)
}
