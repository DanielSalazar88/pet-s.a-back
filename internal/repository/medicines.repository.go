package repository

import (
	"fmt"
	"strings"
)

var (
	errInsertMedicine = "error insertando el medicamento"
	errInsertRecipe   = "error insertando la receta"
)

func (r *MyRepository) InsertMedicine(infoMedicine map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/medicines/insertMedicine.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoMedicine {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	_, err = r.DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errInsertMedicine, err)
	}

	return "OK", nil
}

func (r *MyRepository) InsertRecipe(infoRecipe map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/medicines/insertRecipe.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoRecipe {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	_, err = r.DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errInsertRecipe, err)
	}

	return "OK", nil
}
