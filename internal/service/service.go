package service

import (
	"pet-s.a/internal/entity"
	"pet-s.a/internal/repository"
)

type Service interface {
	InsertClient(client entity.NewClient) (string, error)
	DeleteClient(client entity.DeleteClient) (string, error)
	UpdateClient(client entity.NewClient) (string, error)

	InsertPet(pet entity.NewPet) (string, error)
	DeletePet(pet entity.DeletePet) (string, error)
	UpdatePet(pet entity.UpdatePet) (string, error)

	InsertMedicine(medicine entity.NewMedicine) (string, error)
	DeleteMedicine(medicine entity.DeleteMedicine) (string, error)
	UpdateMedicine(medicine entity.UpdateMedicine) (string, error)

	InsertRecipe(recipe entity.NewRecipe) (string, error)
	DeleteRecipe(recipe entity.DeleteRecipe) (string, error)
}

type MyService struct {
	Repository repository.Repository
}

func New(r repository.Repository) Service {
	return &MyService{
		Repository: r,
	}
}
