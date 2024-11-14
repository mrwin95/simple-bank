package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mrwin95/simple_bank/api"
	db "github.com/mrwin95/simple_bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:Thang123@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8811"
)

func main() {
	// This is the main function
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
