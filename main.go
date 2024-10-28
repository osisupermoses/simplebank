package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // without this our code won't be able to talk to the database
	"github.com/osisupermoses/simplebank/api"
	db "github.com/osisupermoses/simplebank/db/sqlc"
	"github.com/osisupermoses/simplebank/util"
)

func main() {
	// initialize all dependencies
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	// start server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}