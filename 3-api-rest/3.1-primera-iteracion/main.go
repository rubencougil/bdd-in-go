package main

import (
	"bookstore/bookstore"
	"encoding/json"
	"net/http"
	"strings"
)

type App struct {
	store *bookstore.Store
}

func (a *App) addBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var book bookstore.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = a.store.AddBook(&book)
		w.WriteHeader(http.StatusAccepted)
	}
}

func (a *App) getBook(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	if r.Method == http.MethodGet {
		book, err := a.store.GetBook(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		bookJson, err := json.Marshal(book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		_, _ = w.Write(bookJson)
	}
}

func main() {

	app := App{store: &bookstore.Store{}}
	http.HandleFunc("/books", app.addBook)
	http.HandleFunc("/books/:id", app.getBook)
	_ = http.ListenAndServe(":8080", nil)
}
