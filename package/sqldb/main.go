package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT
	)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err = DbConnection.Exec(cmd, "Mike", 20)
	if err != nil {
		log.Fatalln(err)
	}

	cmd = "UPDATE person SET age = ? WHERE name = ?"
	_, err = DbConnection.Exec(cmd, 25, "Mike")
	if err != nil {
		log.Fatalln(err)
	}

	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p)
	}
	var p2 Person
	cmd = "SELECT * FROM person WHERE age = ?"
	row := DbConnection.QueryRow(cmd, 20)
	err = row.Scan(&p2.Name, &p2.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalln("No row")
		} else {
			log.Fatalln(err)
		}
	}
	fmt.Println(p2)

	cmd = "DELETE FROM person WHERE name = ?"
	_, err = DbConnection.Exec(cmd, "Mike")
	if err != nil {
		log.Fatalln(err)
	}
}
