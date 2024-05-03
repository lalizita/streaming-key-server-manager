package db

import (
	"database/sql"
	"fmt"

	"github.com/lalizita/streaming-key-server-manager/config/env"

	_ "github.com/lib/pq"
)

func OpenConn(conf env.EnvConfig) (*sql.DB, error) {
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.PostgresHost, conf.PostgresPort, conf.PostgresUser, conf.PostgrePass, conf.PostgresDB)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	return db, err
}
