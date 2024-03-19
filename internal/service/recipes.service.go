package service

import "pet-s.a/internal/entity"

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

func (s *MyService) DeleteRecipe(recipe entity.DeleteRecipe) (string, error) {
	values := map[string]interface{}{
		"id_replace": recipe.ID,
	}

	result, err := s.Repository.DeleteRecipe(values)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *MyService) GetRecipes() (string, error) {
	result, err := s.Repository.GetRecipes()
	if err != nil {
		return "", err
	}

	return result, err
}
