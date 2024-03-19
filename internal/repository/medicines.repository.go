package repository

import (
	"fmt"
	"strings"
)

var (
	errInsertMedicine = "error insertando el medicamento"
	errDeleteMedicine = "error eliminando el medicamento"
	errUpdateMedicine = "error actualizando el medicamento"
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

func (r *MyRepository) DeleteMedicine(infoMedicine map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/medicines/deleteMedicines.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoMedicine {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	_, err = r.DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errDeleteMedicine, err)
	}

	return "OK", nil
}

func (r *MyRepository) UpdateMedicine(infoMedicine map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/medicines/updateMedicines.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoMedicine {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	_, err = r.DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errUpdateMedicine, err)
	}

	return "OK", nil
}
