Feature: Usuario crea un nuevo libro

  Scenario: Usuario crea un nuevo libro
    When Usuario crea un libro con la siguiente informacion:
      | key    | value                    |
      | id     | 1                        |
      | title  | El Principito            |
      | author | Antoine de Saint-Exupery |
    Then El libro con id "1" y titulo "El Principito" ha sido creado
