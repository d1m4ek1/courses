package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Databse interface {
	Init() (*sqlx.DB, error)
}

type Context struct {
	DB *sqlx.DB
}

const (
	user    string = "postgres"
	pass    string = "aezakmi202"
	dbname  string = "courses"
	sslmode string = "disable"
)

const connectString string = "user=" + user + " password=" + pass + " dbname=" + dbname + " sslmode=" + sslmode

func (c *Context) Init() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
