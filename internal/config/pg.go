package config

import (
	"errors"
	"os"
)

const (
	pgConn = "POSTGRES_CONN"
)

type PGConfig interface {
	CONN() string
}

type pgConfig struct {
	conn string
}

func NewPGConfig() (PGConfig, error) {
	conn := os.Getenv(pgConn)
	if len(conn) == 0 {
		return nil, errors.New("pg connection string not found")
	}

	return &pgConfig{
		conn: conn,
	}, nil
}

func (cfg *pgConfig) CONN() string {
	return cfg.conn
}
