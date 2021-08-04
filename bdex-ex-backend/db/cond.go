package db

import (
	"fmt"
	"strings"

	"github.com/go-xorm/builder"
)

type condIn struct {
	col  string
	vals []interface{}
}

var _ builder.Cond = condIn{}

// In generates IN condition that support BINARY field query
func (db *DB) CondIn(col string, values ...interface{}) builder.Cond {
	return condIn{
		col:  db.Quote(col),
		vals: values,
	}
}

func (condIn condIn) handleBlank(w builder.Writer) error {
	_, err := fmt.Fprint(w, "0=1")
	return err
}

func (condIn condIn) WriteTo(w builder.Writer) error {
	if len(condIn.vals) <= 0 {
		return condIn.handleBlank(w)
	}

	switch condIn.vals[0].(type) {
	case []string:
		vals := condIn.vals[0].([]string)
		if len(vals) <= 0 {
			return condIn.handleBlank(w)
		}
		questionMark := strings.Repeat("unhex(?),", len(vals))
		if _, err := fmt.Fprintf(w, "%s IN (%s)", condIn.col, questionMark[:len(questionMark)-1]); err != nil {
			return err
		}
		for _, val := range vals {
			w.Append(val)
		}
	case [][]byte:
		vals := condIn.vals[0].([][]byte)
		if len(vals) <= 0 {
			return condIn.handleBlank(w)
		}
		questionMark := strings.Repeat("?,", len(vals))
		if _, err := fmt.Fprintf(w, "%s IN (%s)", condIn.col, questionMark[:len(questionMark)-1]); err != nil {
			return err
		}
		for _, val := range vals {
			w.Append(string(val))
		}
	default:
		c := builder.In(condIn.col, condIn.vals...)
		if err := c.WriteTo(w); err != nil {
			return err
		}
	}
	return nil
}

func (condIn condIn) And(conds ...builder.Cond) builder.Cond {
	return builder.And(condIn, builder.And(conds...))
}

func (condIn condIn) Or(conds ...builder.Cond) builder.Cond {
	return builder.Or(condIn, builder.Or(conds...))
}

func (condIn condIn) IsValid() bool {
	return len(condIn.col) > 0 && len(condIn.vals) > 0
}
