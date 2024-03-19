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

func (s *MyService) DeletePet(pet entity.DeletePet) (string, error) {
	values := map[string]interface{}{
		"replace_id": pet.ID,
	}

	result, err := s.Repository.DeletePet(values)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *MyService) UpdatePet(pet entity.UpdatePet) (string, error) {
	values := map[string]interface{}{
		"id_replace":     pet.ID,
		"nombre_replace": pet.Nombre,
		"raza_replace":   pet.Raza,
		"edad_replace":   pet.Edad,
		"peso_replace":   pet.Peso,
	}

	result, err := s.Repository.UpdatePet(values)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *MyService) GetPets() (string, error) {
	result, err := s.Repository.GetPets()
	if err != nil {
		return "", err
	}

	return result, err
}
