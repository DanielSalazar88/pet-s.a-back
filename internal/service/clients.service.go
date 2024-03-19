package service

import "pet-s.a/internal/entity"

func (s *MyService) InsertClient(client entity.NewClient) (string, error) {
	values := map[string]string{
		"cedula_replace":    client.Cedula,
		"nombres_replace":   client.Nombres,
		"apellidos_replace": client.Apellidos,
		"direccion_replace": client.Direccion,
		"telefono_replace":  client.Telefono,
		"correo_replace":    client.Correo,
	}

	result, err := s.Repository.InsertClient(values)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *MyService) DeleteClient(client entity.DeleteClient) (string, error) {
	values := map[string]interface{}{
		"cedula_replace": client.Cedula,
	}

	result, err := s.Repository.DeleteClient(values)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *MyService) UpdateClient(client entity.NewClient) (string, error) {
	values := map[string]string{
		"cedula_replace":    client.Cedula,
		"nombres_replace":   client.Nombres,
		"apellidos_replace": client.Apellidos,
		"direccion_replace": client.Direccion,
		"telefono_replace":  client.Telefono,
		"correo_replace":    client.Correo,
	}

	result, err := s.Repository.UpdateClient(values)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *MyService) GetClients() (string, error) {
	result, err := s.Repository.GetClients()
	if err != nil {
		return "", err
	}

	return result, err
}
