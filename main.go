package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "godemo"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	insertSQLStatement := `
		INSERT INTO users (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
	`
	_, err = db.Exec(insertSQLStatement, 30, "john@mail.com", "John", "Doe")
	if err != nil {
		panic(err)
	}

	updateSQLStatement := `
		UPDATE users
		SET first_name = $2, last_name = $3
		WHERE id = $1;
	`
	res, err := db.Exec(updateSQLStatement, 2, "Max", "Williams")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	sqlStatement := `SELECT id, email FROM users WHERE id=$1;`
	var email string
	var id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	rowQuery1 := db.QueryRow(sqlStatement, 1)
	switch err := rowQuery1.Scan(&id, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}

	type User struct {
		ID        int
		Age       int
		FirstName string
		LastName  string
		Email     string
	}

	QuerySQLStatement := `SELECT * FROM users WHERE id=$1;`
	var user User
	rowQuery2 := db.QueryRow(QuerySQLStatement, 1)
	err = rowQuery2.Scan(&user.ID, &user.Age, &user.FirstName,
		&user.LastName, &user.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}
}
