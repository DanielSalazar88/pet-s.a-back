package service

import "pet-s.a/internal/entity"

func (s *MyService) InsertPet(pet entity.NewPet) (string, error) {
	values := map[string]interface{}{
		"nombre_replace": pet.Nombre,
		"raza_replace":   pet.Raza,
		"edad_replace":   pet.Edad,
		"peso_replace":   pet.Peso,
		"cedula_replace": pet.CedulaCliente,
	}

	result, err := s.Repository.InsertPet(values)
	if err != nil {
		return "", err
	}

	return result, nil
}
