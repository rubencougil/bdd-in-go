Feature: Alumno de OpenWebinars puntúa un curso.

  Scenario: Total de puntuaciones recibidas se actualizan.
    Given el curso "BDD en Go" existe
     When el curso "BDD en Go" recibe una puntuación de "9" del alumno "carlos"
     Then el curso "BDD en Go" ha recibido un total de "1" puntuaciones