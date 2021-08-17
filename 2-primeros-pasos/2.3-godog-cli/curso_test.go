package curso_test

import (
	"fmt"
	"github.com/cucumber/godog"
	curso "openwebinars"
	"testing"
)

var cursos []*curso.Curso

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			// Add step definitions here.
			InitializeScenario(ctx)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func elCursoExiste(arg1 string) error {
	cursos = nil
	cursos = append(cursos, curso.NewCurso(arg1))
	return nil
}

func elCursoHaRecibidoUnTotalDePuntuaciones(arg1 string, arg2 int) error {
	for _, c := range cursos {
		if c.Name == arg1 {
			if c.Votes != arg2 {
				return fmt.Errorf("want %d, got %d", arg2, c.Votes)
			}
		}
	}
	return nil
}

func elCursoRecibeUnaPuntuacinDeDelAlumno(arg1, arg2, arg3 string) error {
	for _, c := range cursos {
		if c.Name == arg1 {
			c.AddScore()
		}
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^el curso "([^"]*)" existe$`, elCursoExiste)
	ctx.Step(`^el curso "([^"]*)" ha recibido un total de "([^"]*)" puntuaciones$`, elCursoHaRecibidoUnTotalDePuntuaciones)
	ctx.Step(`^el curso "([^"]*)" recibe una puntuación de "([^"]*)" del alumno "([^"]*)"$`, elCursoRecibeUnaPuntuacinDeDelAlumno)
}