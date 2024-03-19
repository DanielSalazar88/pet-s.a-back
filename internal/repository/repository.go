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
	UpdateClient(infoReplace map[string]string) (string, error)
	GetClients() (string, error)

	InsertPet(infoPet map[string]interface{}) (string, error)
	DeletePet(infoPet map[string]interface{}) (string, error)
	UpdatePet(infoPet map[string]interface{}) (string, error)
	GetPets() (string, error)

	InsertMedicine(infoMedicine map[string]interface{}) (string, error)
	DeleteMedicine(infoMedicine map[string]interface{}) (string, error)
	UpdateMedicine(infoMedicine map[string]interface{}) (string, error)
	GetMedicines() (string, error)

	InsertRecipe(infoRecipe map[string]interface{}) (string, error)
	DeleteRecipe(infoRecipe map[string]interface{}) (string, error)
	GetRecipes() (string, error)
}

type MyRepository struct {
	DB database.DataBase
}

func New(db database.DataBase) Repository {
	return &MyRepository{
		DB: db,
	}
}
