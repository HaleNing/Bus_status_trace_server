// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/HaleNing/bustrack/src/Model/ent/book"
)

// Book is the model entity for the Book schema.
type Book struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BookName holds the value of the "book_name" field.
	BookName string `json:"book_name,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Book) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case book.FieldID:
			values[i] = new(sql.NullInt64)
		case book.FieldBookName, book.FieldAuthor:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Book", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Book fields.
func (b *Book) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case book.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case book.FieldBookName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field book_name", values[i])
			} else if value.Valid {
				b.BookName = value.String
			}
		case book.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				b.Author = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Book.
// Note that you need to call Book.Unwrap() before calling this method if this Book
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Book) Update() *BookUpdateOne {
	return (&BookClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Book entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Book) Unwrap() *Book {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Book is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Book) String() string {
	var builder strings.Builder
	builder.WriteString("Book(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("book_name=")
	builder.WriteString(b.BookName)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(b.Author)
	builder.WriteByte(')')
	return builder.String()
}

// Books is a parsable slice of Book.
type Books []*Book

func (b Books) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
