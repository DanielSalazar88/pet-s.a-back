package entity

type NewClient struct {
	Cedula    string `json:"cedula" validate:"required"`
	Nombres   string `json:"nombres" validate:"required"`
	Apellidos string `json:"apellidos" validate:"required"`
	Direccion string `json:"direccion" validate:"required"`
	Telefono  string `json:"telefono" validate:"required"`
	Correo    string `json:"correo" validate:"required"`
}

type DeleteClient struct {
	Cedula string `json:"cedula" validate:"required"`
}
