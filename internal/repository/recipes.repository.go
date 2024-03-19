package repository

import (
	"fmt"
	"strings"
)

var (
	errInsertRecipe = "error insertando la receta"
	errDeleteRecipe = "error eliminando la receta"
)

func (r *MyRepository) InsertRecipe(infoRecipe map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/recipes/insertRecipe.sql")
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

func (r *MyRepository) DeleteRecipe(infoRecipe map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/recipes/deleteRecipe.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoRecipe {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	_, err = r.DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errDeleteRecipe, err)
	}

	return "OK", nil
}
