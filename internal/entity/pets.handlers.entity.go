package entity

type NewPet struct {
	Nombre        string  `json:"nombre" validate:"required"`
	Raza          string  `json:"raza" validate:"required"`
	Edad          int     `json:"edad" validate:"required"`
	Peso          float32 `json:"peso" validate:"required"`
	CedulaCliente string  `json:"cedula_cliente" validate:"required"`
}

type DeletePet struct {
	ID int `json:"id_mascota" validate:"required"`
}

type UpdatePet struct {
	ID     int     `json:"id_mascota" validate:"required"`
	Nombre string  `json:"nombre" validate:"required"`
	Raza   string  `json:"raza" validate:"required"`
	Edad   int     `json:"edad" validate:"required"`
	Peso   float32 `json:"peso" validate:"required"`
}
