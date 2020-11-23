package database

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var defaultConn *sql.DB

func Init(env string) error {
	if defaultConn != nil {
		_ = defaultConn.Close()
		defaultConn = nil
	}

	var dsn string
	switch env {
	case "test":
		dsn = "root:123456@tcp(0.0.0.0:33066)/test"
	default:
		return errors.New("unknown env")
	}

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	conn.SetMaxOpenConns(64)
	conn.SetConnMaxLifetime(30 * time.Second)
	defaultConn = conn
	return nil
}

func Get() *sql.DB {
	return defaultConn
}

func MustGet() *sql.DB {
	for defaultConn == nil {
		time.Sleep(100 * time.Millisecond)
	}
	return defaultConn
}

func CloseRow(row *sql.Rows) {
	if row != nil {
		_ = row.Close()
	}
}
