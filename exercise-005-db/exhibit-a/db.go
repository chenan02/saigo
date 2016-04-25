package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Person struct {
	ID   int    `db:"person_id"`
	Name string `db:"name"`
	SSN  int    `db:"ssn"`
}

func main() {
	db, err := sqlx.Open("postgres", "dbname=test sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.QueryRow("INSERT INTO people(name, ssn) VALUES($1, $2)", "Bruce Leroy", 111223333)
	db.QueryRow("INSERT INTO people(name, ssn) VALUES($1, $2)", "Sho 'Nuff", 444556666)

	people := []Person{}
	db.Select(&people, "SELECT person_id, name, ssn FROM people")

	for _, p := range people {
		fmt.Printf("Person %5d %-15s %9d\n", p.ID, p.Name, p.SSN)
	}
}
