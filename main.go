package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mrwin95/simple_bank/api"
	db "github.com/mrwin95/simple_bank/db/sqlc"
	"github.com/mrwin95/simple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	// This is the main function
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
