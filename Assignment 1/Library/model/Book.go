package model

type Book struct {
	ID         uint64
	Title      string
	Author     string
	IsBorrowed bool
}

var IncrementingID uint64 = 1

func CreateBook(title, author string) Book {
	b := Book{
		ID:     IncrementingID,
		Title:  title,
		Author: author,
	}

	IncrementingID++

	return b
}
