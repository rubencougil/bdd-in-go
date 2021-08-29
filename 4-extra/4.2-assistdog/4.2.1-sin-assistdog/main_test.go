package main_test

import (
	"errors"
	"github.com/cucumber/godog"
)

type Book struct {
	ID string
	Title string
}

var book *Book

func creoUnNuevoLibroConLaSiguienteInformacion(books *godog.Table) error {

	book = &Book{
		ID:    books.Rows[1].Cells[0].Value,
		Title: books.Rows[1].Cells[1].Value,
	}

	return nil
}

func elLibroConIdYTituloExiste(id, title string) error {
	if book.ID != id || book.Title != title {
		return errors.New("book was not created")
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Creo un nuevo libro con la siguiente informacion:$`, creoUnNuevoLibroConLaSiguienteInformacion)
	ctx.Step(`^El libro con id "([^"]*)" y titulo "([^"]*)" existe$`, elLibroConIdYTituloExiste)
}