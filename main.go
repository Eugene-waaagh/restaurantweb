package main

import (
	"database/sql"
	"github.com/eugene-waaagh/restaurantweb/db/api"
	db "github.com/eugene-waaagh/restaurantweb/db/sqlc"
	"github.com/eugene-waaagh/restaurantweb/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDrive, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
