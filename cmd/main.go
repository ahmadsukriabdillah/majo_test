package main

import (
	"database/sql"
	"log"
	"majo_test/cmd/bootrstap"
	"majo_test/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	tokerMaker, err := utils.NewJWTMaker(config.AccessToken)

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error: Failed init database", err.Error())
	}

	server, err := bootrstap.NewServer(config, db, tokerMaker)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
