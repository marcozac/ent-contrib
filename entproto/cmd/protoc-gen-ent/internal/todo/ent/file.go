// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/marcozac/ent-contrib/entproto/cmd/protoc-gen-ent/internal/todo/ent/file"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Contents holds the value of the "contents" field.
	Contents     string `json:"contents,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldID, file.FieldContents:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case file.FieldContents:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contents", values[i])
			} else if value.Valid {
				f.Contents = value.String
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the File.
// This includes values selected through modifiers, order, etc.
func (f *File) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return NewFileClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("contents=")
	builder.WriteString(f.Contents)
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File
