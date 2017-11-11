package take

import "testing"

type Book struct {
	Name string
}

func TestString(t *testing.T) {
	books := []*Book{
		&Book{
			Name: "foo",
		},
		&Book{
			Name: "Bar",
		},
	}

	names := String(books, func(i int) string {
		return books[i].Name
	})

	if names[0] != "foo" || names[1] != "Bar" {
		t.Error("take string failed")
	}
}
