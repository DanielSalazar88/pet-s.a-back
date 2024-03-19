package service

import "pet-s.a/internal/entity"

func (s *MyService) GeneralClientReport(client entity.GeneralClientReport) (string, error) {
	values := map[string]string{
		"cedula_replace": client.CedulaCliente,
	}

	result, err := s.Repository.GeneralClientReport(values)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *MyService) RecipesPetReport(pet entity.RecipesPetReport) (string, error) {
	values := map[string]interface{}{
		"id_mascota_replace": pet.IDMascota,
	}

	result, err := s.Repository.RecipesPetReport(values)
	if err != nil {
		return "", err
	}

	return result, nil
}
