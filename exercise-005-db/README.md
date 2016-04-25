## Description
Learn to work with SQL in Go!

## Prerequisites
This exercise assumes the learner is familiar with basic SQL. If you are new to SQL then you must first
learn the basics before proceeding. Saigo uses [PostgreSQL](http://www.tutorialspoint.com/postgresql/) for all database work.

Set up a local PostgreSQL server on your deveopment machine and create a database called `test`. Connect to the database
(`$ psql test`) and copy the code in `migration_up.sql` to create the `people` table in your `test` database.

## Comprehension Tasks

The standard Go distribution includes a `database/sql` package with implementations for various databases, including PostgreSQL.
A nice extension to `database/sql` is the [`sqlx`](github.com/jmoiron/sqlx) package that uses some reflection to create a slightly
prettier interface.

Check out the code in [exhibit-a]():

```
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
```
Make sure you followed the steps above to create the `people` table in your `test` database. Run this code on the console
and see it work. If you have issues seek the assistance of an instructor.

Explain what each line of code does in `exhibit-a`.

# Engineering Tasks

Create an `exhibit-b` folder and expand the code (from `exhibit-a`) to:

1. Delete a person
1. Update a person's SSN
1. Insert a person and immediately retrieve the person's ID assigned by the database
1. Handle errors that might be returned by each of the function calls made on `db`
