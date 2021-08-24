package main

import (
	"bookstore/bookstore"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"net/http"
	"net/http/httptest"
)

var app = App{store: &bookstore.Store{}}
var e error = nil

func usuarioCreaUnLibroConLaSiguienteInformacion(table *godog.Table) (err error) {

	assist := assistdog.NewDefault()

	data, err := assist.ParseMap(table)
	if err != nil {
		return
	}

	if err != nil {
		return
	}
	book := &bookstore.Book{
		ID: data["id"],
		Title: data["title"],
		Author: data["author"],
	}

	bookJson, err := json.Marshal(book)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookJson))
	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	app.addBook(w, req)

	if w.Code != http.StatusAccepted {
		//return errors.New(fmt.Sprintf("got %d, want %d", w.Code, http.StatusAccepted))
		e = errors.New(w.Body.String())
	}

	return
}

func elLibroConIdYTituloHaSidoCreado(id, title string) (err error) {
	req, err := http.NewRequest(http.MethodGet, "/books/" + id, nil)
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	app.getBook(w, req)

	if w.Code != http.StatusOK {
		return errors.New(fmt.Sprintf("got %d, want %d", w.Code, http.StatusOK))
	}

	var book bookstore.Book
	err = json.NewDecoder(w.Body).Decode(&book)
	if err != nil {
		return
	}

	if book.Title != title {
		return errors.New(fmt.Sprintf("got %s, want %s", book.Title, title))
	}

	return
}

func laLibreriaContieneLosSiguientesLibros(table *godog.Table) (err error) {
	assist := assistdog.NewDefault()
	books, err := assist.CreateSlice(new(bookstore.Book), table)
	if err != nil {
		return
	}
	
	for _, book := range books.([]*bookstore.Book) {
		err := app.store.AddBook(book)
		if err != nil {
			return err
		}
	}
	
	return 
	
}

func seProduceUnError(msg string) error {
	if e != nil && e.Error() != msg {
		return errors.New(fmt.Sprintf("want %s, got %s", msg, e.Error()))
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		app.store = &bookstore.Store{}
		return ctx, nil
	})

	ctx.Step(`^El libro con id "([^"]*)" y titulo "([^"]*)" ha sido creado$`, elLibroConIdYTituloHaSidoCreado)
	ctx.Step(`^Usuario crea un libro con la siguiente informacion:$`, usuarioCreaUnLibroConLaSiguienteInformacion)
	ctx.Step(`^La libreria contiene los siguientes libros:$`, laLibreriaContieneLosSiguientesLibros)
	ctx.Step(`^Se produce un error: "([^"]*)"$`, seProduceUnError)
}
