package connection

import (
	"context"
	"fmt"
	"study/database/handlers"
	"study/utils"

	"github.com/jackc/pgx"
	"log"
)

func DB() *pgx.Conn {
	cfg := utils.GetConfig()

	var config pgx.ConnConfig
	config.Host = cfg.Host
	config.Port = cfg.Port
	config.Database = cfg.DBName
	config.User = cfg.User
	config.Password = cfg.Password

	// open database
	db, err := pgx.Connect(config)
	handlers.CheckError(err)

	// check db
	err = db.Ping(context.Background())
	handlers.CheckError(err)

	fmt.Println("Connected!")
	return db
}

func GetTX(db *pgx.Conn) *pgx.Tx {

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	return tx
}

func CloseDB(db *pgx.Conn) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connection is closed!")
	}
}
