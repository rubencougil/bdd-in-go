package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	ID string
	Title string
	Author string
}

type Store struct {
	books []*Book
}

func (s *Store) AddBook(book *Book) (err error) {
	s.books = append(s.books, book)
	return
}

func (s *Store) GetBook(id string) (book *Book, err error) {
	for _, book := range s.books {
		if book.ID == id {
			return book, nil
		}
 	}
 	return nil, errors.New(fmt.Sprintf("Book with id %s not found", id))
}
