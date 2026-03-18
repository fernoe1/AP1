package model

import "errors"

type Library struct {
	Books map[string]*Book
}

func (lib *Library) initialize() {
	lib.Books = make(map[string]*Book)
}

func (lib *Library) AddBook(book *Book) error {
	if lib.Books == nil {
		lib.initialize()
	}

	if lib.Books[book.Title] != nil {
		return errors.New("book already exists")
	}

	lib.Books[book.Title] = book
	return nil
}

func (lib *Library) BorrowBook(title string) (*Book, error) {
	if lib.Books == nil {
		lib.initialize()
	}

	b := lib.Books[title]
	if b == nil {
		return nil, errors.New("book not found")
	}

	b.IsBorrowed = true
	return b, nil
}

func (lib *Library) ReturnBook(title string) error {
	if lib.Books == nil {
		lib.initialize()
	}

	b := lib.Books[title]
	if b == nil {
		return errors.New("book not found")
	}
	if !b.IsBorrowed {
		return errors.New("book is not borrowed")
	}

	b.IsBorrowed = false
	return nil
}

func (lib *Library) ListAvailableBooks() ([]Book, error) {
	if lib.Books == nil {
		lib.initialize()
	}

	books := make([]Book, 0, len(lib.Books))
	for _, book := range lib.Books {
		if book.IsBorrowed == false {
			books = append(books, *book)
		}
	}

	return books, nil
}
