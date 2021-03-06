// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/HaleNing/bustrack/src/Model/ent/book"
	"github.com/HaleNing/bustrack/src/Model/ent/schema"
	"github.com/HaleNing/bustrack/src/Model/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescBookName is the schema descriptor for book_name field.
	bookDescBookName := bookFields[0].Descriptor()
	// book.DefaultBookName holds the default value on creation for the book_name field.
	book.DefaultBookName = bookDescBookName.Default.(string)
	// bookDescAuthor is the schema descriptor for author field.
	bookDescAuthor := bookFields[1].Descriptor()
	// book.AuthorValidator is a validator for the "author" field. It is called by the builders before save.
	book.AuthorValidator = func() func(string) error {
		validators := bookDescAuthor.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(author string) error {
			for _, fn := range fns {
				if err := fn(author); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[0].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
}
