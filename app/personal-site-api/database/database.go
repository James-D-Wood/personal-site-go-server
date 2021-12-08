package database

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jdwoo/personal-site-go-server/app/personal-site-api/cfg"
	"github.com/jdwoo/personal-site-go-server/internal/webserverutils"
)

func InitalizeDatabase(config *cfg.Config) (*pgx.Conn, error) {
	var err error
	db, err := pgx.Connect(context.Background(), config.Database.GetConnectionString())
	return db, err
}

func TranslateError(err error) error {
	if strings.Contains(err.Error(), "violates unique constraint") {
		return webserverutils.RequestError{Message: "Unique constraint violated."}
	} else {
		return err
	}
}

func TeardownDatabase(db *pgx.Conn) {
	db.Close(context.Background())
}
