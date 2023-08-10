package db

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func DefaultConfig() Postgres {
	return Postgres{
		Host:     "localhost",
		Port:     "3357",
		User:     "scotty",
		Password: "pizza",
		DBName:   "scottygpizza",
		SSLMode:  "disable",
	}
}

func Open(c Postgres) (*sql.DB, error) {
	sql, err := sql.Open("pgx", c.String())
	if err != nil {
		fmt.Errorf("Open: %v", err)
	}

	return sql, nil
}

func (c Postgres) String() string {
	return fmt.Sprintf("host=%s,port=%s,user=%s,password=%s,dbname=%s,sslmode=%s", c.Host, c.Port,
		c.User, c.Password, c.Port, c.SSLMode)
}
