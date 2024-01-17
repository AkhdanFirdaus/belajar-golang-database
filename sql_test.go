package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer (id, name) VALUES('akhdan', 'Akhdan')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id,name FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}

	fmt.Println("Success Select customer")
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("============================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}

	fmt.Println("Success Select customer")
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username=? AND password=? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses Login")
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "user1"
	password := "password"

	query := "INSERT INTO user (username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "user1@gmail.com"
	comment := "Testing Komentar"

	query := "INSERT INTO comments (email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	inserId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id =", inserId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments (email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 2; i < 12; i++ {
		email := "user" + strconv.Itoa(i) + "@gmail.com"
		comment := "Testing Komentar ke-" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}
}

func TestDatabaseTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments (email, comment) VALUES(?, ?)"

	for i := 12; i <= 24; i++ {
		email := "user" + strconv.Itoa(i) + "@gmail.com"
		comment := "Testing Komentar ke-" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		panic(err)
	}
}
