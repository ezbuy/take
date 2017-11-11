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

func TestStringUnique(t *testing.T) {
	books := []*Book{
		&Book{
			Name: "foo",
		},
		&Book{
			Name: "Bar",
		},
		&Book{
			Name: "foo",
		},
		&Book{
			Name: "Bar",
		},
	}

	names := StringUnique(books, func(i int) string {
		return books[i].Name
	})

	if names[0] != "foo" || names[1] != "Bar" {
		t.Error("take string failed")
	}
}

type (
	Item struct {
		ID int64
		S  string
	}
)

func testItems() []Item {
	items := make([]Item, 1000)
	for i := range items {
		items[i] = Item{ID: int64(i), S: "test string"}
	}
	return items
}

func BenchmarkIndex(b *testing.B) {
	items := testItems()
	b.ResetTimer()
	for c := 0; c < b.N; c++ {
		ids := make([]int64, len(items))
		for i := range ids {
			ids[i] = items[i].ID
		}
	}
}

func BenchmarkAppend(b *testing.B) {
	items := testItems()
	b.ResetTimer()
	for c := 0; c < b.N; c++ {
		ids := make([]int64, 0, len(items))
		for i := range items {
			ids = append(ids, items[i].ID)
		}
	}
}

func BenchmarkTake(b *testing.B) {
	items := testItems()
	b.ResetTimer()
	for c := 0; c < b.N; c++ {
		ids := Int64(items, func(i int) int64 {
			return items[i].ID
		})
		_ = ids
	}
}

func BenchmarkTakeUniqueFor(b *testing.B) {
	items := testItems()
	b.ResetTimer()
	for c := 0; c < b.N; c++ {
		ids := int64UniqueFor(len(items), func(i int) int64 {
			return items[i].ID
		})
		_ = ids
	}
}

func BenchmarkTakeUniqueMap(b *testing.B) {
	items := testItems()
	b.ResetTimer()
	for c := 0; c < b.N; c++ {
		ids := int64UniqueMap(len(items), func(i int) int64 {
			return items[i].ID
		})
		_ = ids
	}
}
