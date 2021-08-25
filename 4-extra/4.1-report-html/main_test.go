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

type TestContext struct {
	App App
	Error error
}

func usuarioCreaUnLibroConLaSiguienteInformacion(cxt context.Context, table *godog.Table) (err error) {

	t := cxt.Value("TestContext").(*TestContext)

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
	t.App.addBook(w, req)

	if w.Code != http.StatusAccepted {
		t.Error = errors.New(w.Body.String())
	}

	return
}

func elLibroConIdYTituloHaSidoCreado(cxt context.Context, id, title string) (err error) {

	t := cxt.Value("TestContext").(*TestContext)

	req, err := http.NewRequest(http.MethodGet, "/books/" + id, nil)
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	t.App.handleBook(w, req)

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

func laLibreriaContieneLosSiguientesLibros(cxt context.Context, table *godog.Table) (err error) {

	t := cxt.Value("TestContext").(*TestContext)

	assist := assistdog.NewDefault()
	books, err := assist.CreateSlice(new(bookstore.Book), table)
	if err != nil {
		return
	}
	
	for _, book := range books.([]*bookstore.Book) {
		err := t.App.store.AddBook(book)
		if err != nil {
			return err
		}
	}
	
	return
}

func seProduceUnError(cxt context.Context, msg string) error {

	t := cxt.Value("TestContext").(*TestContext)

	if t.Error != nil && t.Error.Error() != msg {
		return errors.New(fmt.Sprintf("want %s, got %s", msg, t.Error.Error()))
	}

	return nil
}

func noSeHaProducidoNingunError(cxt context.Context) error {

	t := cxt.Value("TestContext").(*TestContext)

	if t.Error != nil {
		return errors.New(fmt.Sprintf("unexpected error: %s", t.Error.Error()))
	}

	return nil
}

func usuarioBorraElLibroConId(cxt context.Context, id string) (err error) {

	t := cxt.Value("TestContext").(*TestContext)

	req, err := http.NewRequest(http.MethodDelete, "/books/" + id, nil)
	if err != nil {
		return
	}

	w := httptest.NewRecorder()
	t.App.handleBook(w, req)

	if w.Code != http.StatusOK {
		t.Error = errors.New(w.Body.String())
		return nil
	}

	return nil
}

func elLibroConIdYTituloHaSidoBorrado(cxt context.Context, id string) (err error) {

	t := cxt.Value("TestContext").(*TestContext)

	req, err := http.NewRequest(http.MethodGet, "/books/" + id, nil)
	if err != nil {
		return
	}
	w := httptest.NewRecorder()
	t.App.handleBook(w, req)

	if w.Code != http.StatusNotFound {
		return errors.New(fmt.Sprintf("got %d, want %d", w.Code, http.StatusNotFound))
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {

		testContext := &TestContext{
			App: App{store: &bookstore.Store{}},
			Error: nil,
		}

		return context.WithValue(ctx, "TestContext", testContext), nil
	})

	ctx.Step(`^El libro con id "([^"]*)" y titulo "([^"]*)" ha sido creado$`, elLibroConIdYTituloHaSidoCreado)
	ctx.Step(`^Usuario crea un libro con la siguiente informacion:$`, usuarioCreaUnLibroConLaSiguienteInformacion)
	ctx.Step(`^La libreria contiene los siguientes libros:$`, laLibreriaContieneLosSiguientesLibros)
	ctx.Step(`^Se produce un error: "([^"]*)"$`, seProduceUnError)
	ctx.Step(`^No se ha producido ningun error$`, noSeHaProducidoNingunError)
	ctx.Step(`^Usuario borra el libro con id "([^"]*)"$`, usuarioBorraElLibroConId)
	ctx.Step(`^El libro con id "([^"]*)" y titulo "([^"]*)" ha sido borrado$`, elLibroConIdYTituloHaSidoBorrado)
}
