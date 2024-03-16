package service

import "pet-s.a/internal/entity"

func (s *MyService) InsertMedicine(medicine entity.NewMedicine) (string, error) {
	values := map[string]interface{}{
		"nombre_replace":      medicine.Nombre,
		"descripcion_replace": medicine.Descripcion,
		"dosis_replace":       medicine.Dosis,
	}

	result, err := s.Repository.InsertMedicine(values)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *MyService) InsertRecipe(recipe entity.NewRecipe) (string, error) {
	values := map[string]interface{}{
		"id_mascota_replace":     recipe.IDMascota,
		"id_medicamento_replace": recipe.IDMedicamento,
	}

	result, err := s.Repository.InsertRecipe(values)
	if err != nil {
		return "", err
	}

	return result, nil
}
