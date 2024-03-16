package database

import (
	"database/sql"

	"pet-s.a/settings"
)

type DataBase interface {
	ExecMysqlQuery(query string) (string, error)
	mysqlConn() (*sql.DB, error)
}

type MyDataBase struct {
	Settings *settings.MySettings
}

func New(s *settings.MySettings) DataBase {
	return &MyDataBase{
		Settings: s,
	}
}
