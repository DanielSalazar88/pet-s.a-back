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

func (s *MyService) DeleteMedicine(medicine entity.DeleteMedicine) (string, error) {
	values := map[string]interface{}{
		"replace_id": medicine.ID,
	}

	result, err := s.Repository.DeleteMedicine(values)
	if err != nil {
		return "", err
	}

	return result, err
}

func (s *MyService) UpdateMedicine(medicine entity.UpdateMedicine) (string, error) {
	values := map[string]interface{}{
		"id_replace":          medicine.ID,
		"nombre_replace":      medicine.Nombre,
		"descripcion_replace": medicine.Descripcion,
		"dosis_replace":       medicine.Dosis,
	}

	result, err := s.Repository.UpdateMedicine(values)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *MyService) GetMedicines() (string, error) {
	result, err := s.Repository.GetMedicines()
	if err != nil {
		return "", err
	}

	return result, err
}
