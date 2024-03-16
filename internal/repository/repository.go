package repository

import (
	"embed"
	"errors"

	"pet-s.a/internal/database"
)

var (
	//go:embed queries/*
	content embed.FS

	errReadSqlFIle = errors.New("error leyendo el archivo sql")
)

type Repository interface {
	InsertClient(infoReplace map[string]string) (string, error)
	DeleteClient(infoClientDelete map[string]interface{}) (string, error)

	InsertPet(infoPet map[string]interface{}) (string, error)

	InsertMedicine(infoMedicine map[string]interface{}) (string, error)
	InsertRecipe(infoRecipe map[string]interface{}) (string, error)
}

type MyRepository struct {
	DB database.DataBase
}

func New(db database.DataBase) Repository {
	return &MyRepository{
		DB: db,
	}
}
