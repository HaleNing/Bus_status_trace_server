package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("book_name").Default("default_book"),
		field.String("author").Default("unknown").NotEmpty().MaxLen(10),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return nil
}