package repository

import (
	"fmt"
	"strings"
)

var (
	errInsertPet = "error insertando la mascota"
)

func (r *MyRepository) InsertPet(infoPet map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/pets/insertPet.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoPet {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	_, err = r.DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errInsertPet, err)
	}

	return "OK", nil
}
