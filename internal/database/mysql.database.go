package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	errMysqlConn = errors.New("error conectando a mysql")
	errMysqlExec = errors.New("error ejecutando consulta a mysql")
	errMysqlScan = errors.New("Error al escanear el resultado")
)

func (db *MyDataBase) mysqlConn() (*sql.DB, error) {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		db.Settings.Mysql.Username,
		db.Settings.Mysql.Password,
		db.Settings.Mysql.Address,
		db.Settings.Mysql.Port,
		db.Settings.Mysql.Database,
	)

	mysql, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, fmt.Errorf("%s : %s", errMysqlConn, err)
	}

	return mysql, nil
}

func (db *MyDataBase) ExecMysqlQuery(query string) (string, error) {
	mysqlClient, err := db.mysqlConn()
	if err != nil {
		return "", err
	}

	err = mysqlClient.Ping()
	if err != nil {
		return "", fmt.Errorf("%s : %s", errMysqlConn, err)
	}

	rows, err := mysqlClient.Query(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errMysqlExec, err)
	}

	var result string
	for rows.Next() {
		if err := rows.Scan(&result); err != nil {
			return "", fmt.Errorf("%s : %s", errMysqlScan, err)
		}
	}

	return result, nil
}
