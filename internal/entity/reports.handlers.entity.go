package entity

type GeneralClientReport struct {
	CedulaCliente string `json:"cedula_cliente" validate:"required"`
}

type RecipesPetReport struct {
	IDMascota int `json:"id_mascota" validate:"required"`
}
