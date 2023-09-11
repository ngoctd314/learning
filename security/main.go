package main

import (
	"context"
	"security/owasp/cmdinjection"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	ctx context.Context
)

func main() {
	cmdinjection.APIServer()
}
