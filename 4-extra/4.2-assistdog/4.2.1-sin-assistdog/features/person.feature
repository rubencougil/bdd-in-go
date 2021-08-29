Feature: Crear un libro

  Scenario: Crear un nuevo libro
    When Creo un nuevo libro con la siguiente informacion:
      | ID   | Title         |
      | 1    | El Principito |
    Then El libro con id "1" y titulo "El Principito" existe