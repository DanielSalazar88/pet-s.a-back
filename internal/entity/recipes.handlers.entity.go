package entity

type NewRecipe struct {
	IDMedicamento int `json:"id_medicamento" validate:"required"`
	IDMascota     int `json:"id_mascota" validate:"required"`
}

type DeleteRecipe struct {
	ID int `json:"id_receta" validate:"required"`
}
