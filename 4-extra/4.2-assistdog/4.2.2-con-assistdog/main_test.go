package main

import (
	"errors"
	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
)

type Book struct {
	ID string
	Title string
}

var book *Book

func creoUnNuevoLibroConLaSiguienteInformacion(books *godog.Table) error {

	assist := assistdog.NewDefault()
	b, err := assist.CreateSlice(new(Book), books)
	if err != nil {
		return err
	}

	book = b.([]*Book)[0]

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