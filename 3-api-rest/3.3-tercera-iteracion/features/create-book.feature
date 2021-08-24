Feature: Usuario crea un nuevo libro

  Scenario: Usuario crea un nuevo libro que ya existe. Se produce un error.
    Given La libreria contiene los siguientes libros:
      | ID | Title                | Author                   |
      | 1  | El Principito        | Antoine de Saint-Exupery |
      | 2  | El Nombre de la Rosa | Umberto Eco              |
    When Usuario crea un libro con la siguiente informacion:
      | key    | value                    |
      | id     | 1                        |
      | title  | El Principito            |
      | author | Antoine de Saint-Exupery |
    Then Se produce un error: "el libro con id 1 ya existe"

  Scenario: Usuario crea un nuevo libro.
    When Usuario crea un libro con la siguiente informacion:
      | key    | value                    |
      | id     | 1                        |
      | title  | El Principito            |
      | author | Antoine de Saint-Exupery |
    Then El libro con id "1" y titulo "El Principito" ha sido creado
     And No se ha producido ningun error
