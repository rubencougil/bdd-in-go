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
	if s.BookExists(book.ID) {
		return errors.New(fmt.Sprintf("el libro con id %s ya existe", book.ID))
	}
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

func (s *Store) BookExists(id string) bool {
	_ , err := s.GetBook(id)
	if err != nil {
		return false
	}
	return true
}
