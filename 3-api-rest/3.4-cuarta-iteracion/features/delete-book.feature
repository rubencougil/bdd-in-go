Feature: Usuario borra un nuevo libro

  Scenario: Usuario trata de borrar un libro, pero el libro no existe.
    When Usuario borra el libro con id "1"
    Then Se produce un error: "el libro con id 1 no existe"

  Scenario: Usuario borra un libro que existe.
    Given Usuario crea un libro con la siguiente informacion:
      | key    | value                    |
      | id     | 1                        |
      | title  | El Principito            |
      | author | Antoine de Saint-Exupery |
     And El libro con id "1" y titulo "El Principito" ha sido creado
    When Usuario borra el libro con id "1"
    Then El libro con id "1" y titulo "El Principito" ha sido borrado
