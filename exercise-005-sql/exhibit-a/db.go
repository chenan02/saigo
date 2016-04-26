package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PanicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Bruce Leroy", 111223333)
	PanicOn(err)

	_, err = db.Exec("INSERT INTO people(name, ssn) VALUES ($1, $2)", "Sho 'Nuff", 444556666)
	PanicOn(err)

	rows, err := db.Query("SELECT person_id, name, ssn FROM people")
	PanicOn(err)

	for rows.Next() {
		p := Person{}
		err := rows.Scan(&p.ID, &p.Name, &p.SSN)
		PanicOn(err)

		fmt.Printf("Person %5d %-15s %9d\n", p.ID, p.Name, p.SSN)
	}
}
