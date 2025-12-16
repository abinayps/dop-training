package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserDetails struct {
	ID        int8
	FirstName string
	LastName  string
	Age       int
	City      string
	Email     string
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := DBConnection(ctx)
	if err != nil {
		fmt.Println("DB Connection Failed: ", err)
		return
	}
	defer db.Close()

	// ************** INSERT QUERY *******************
	_, err = db.Exec(ctx,
		"insert into user_details (first_name,last_name,age,city,email) values ($1,$2,$3,$4,$5)",
		"Hardik", "Pandya", 40, "Mumbai", "pandya@gmail.com",
	)
	if err != nil {
		fmt.Println("Error querying the database: ", err)
		return
	}
	fmt.Println("Data inserted successfully")

	// ************** SELECT Query with single record ***********
	var user UserDetails
	err = db.QueryRow(ctx, "select first_name,last_name,age,city,email from user_details where id=$1", 1).
		Scan(&user.FirstName, &user.LastName, &user.Age, &user.City, &user.Email)
	if err != nil {
		fmt.Println("Error selecting one record from the database: ", err)
		return
	}
	fmt.Println("User Details after querying: ", user)

	// ***************** SELECT Query to fetch all records ***********
	rows, err := db.Query(ctx, "select id,first_name,last_name,age,city,email from user_details")
	if err != nil {
		fmt.Println("Error querying the database: ", err)
		return
	}
	defer rows.Close()
	var users = make([]UserDetails, 0)
	for rows.Next() {
		var u UserDetails
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.City, &u.Email)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
			return
		}
		users = append(users, u)
	}
	fmt.Println("All users selected: ", users)

	// *********** Update Query for updating User Details *********
	commandTag, err := db.Exec(ctx,
		"update user_details set first_name=$1, last_name=$2, age=$3, city=$4, email=$5 where id=$6",
		"Sourav", "Ganguly", 42, "Kolkata", "sourav@gmail.com", 5)
	if err != nil {
		fmt.Println("Error updating the record: ", err)
		return
	}
	if commandTag.RowsAffected() == 0 {
		fmt.Println("No rows updated")
		return
	}
	fmt.Println("Updated the record successfully ")

	// ********** DELETE query **************
	commandTag, err = db.Exec(ctx, "delete from user_details where id=1")
	if err != nil {
		fmt.Println("Error deleting the record: ", err)
		return
	}
	if commandTag.RowsAffected() == 0 {
		fmt.Println("No rows deleted")
		return
	}
	fmt.Println("Record deleted successfully")
}

func DBConnection(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres"
	//postgres://username:password@host:port/database

	// Parse pool config
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
		return nil, err
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to create pool: %v\n", err)
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("Database ping failed: %v\n", err)
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL!")
	return db, nil
}
