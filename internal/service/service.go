package service

import (
	"pet-s.a/internal/entity"
	"pet-s.a/internal/repository"
)

type Service interface {
	InsertClient(client entity.NewClient) (string, error)
	DeleteClient(client entity.DeleteClient) (string, error)

	InsertPet(pet entity.NewPet) (string, error)

	InsertMedicine(medicine entity.NewMedicine) (string, error)
	InsertRecipe(recipe entity.NewRecipe) (string, error)
}

type MyService struct {
	Repository repository.Repository
}

func New(r repository.Repository) Service {
	return &MyService{
		Repository: r,
	}
}
