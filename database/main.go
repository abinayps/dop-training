package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserDetails struct {
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
	_, err = db.Exec(ctx,
		"insert into user_details (first_name,last_name,age,city,email) values ($1,$2,$3,$4,$5)",
		"Sourav", "Ganguly", 40, "Kolkata", "sourav@gmail.com",
	)
	if err != nil {
		fmt.Println("Error querying the database: ", err)
		return
	}
	fmt.Println("Data inserted successfully")
	// var user UserDetails
	// err = db.QueryRow(ctx, "select first_name,last_name,age,city,email from user_details where id=$1", 2).
	// 	Scan(&user.FirstName, &user.LastName, &user.Age, &user.City, &user.Email)
	// if err != nil {
	// 	fmt.Println("Error querying the database: ", err)
	// 	return
	// }
	// fmt.Println("User Details after querying: ", user)
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
