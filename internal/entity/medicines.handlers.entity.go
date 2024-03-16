package entity

type NewMedicine struct {
	Nombre      string `json:"nombre" validate:"required"`
	Descripcion string `json:"descripcion" validate:"required"`
	Dosis       string `json:"dosis" validate:"required"`
}

type NewRecipe struct {
	IDMedicamento int `json:"id_medicamento" validate:"required"`
	IDMascota     int `json:"id_mascota" validate:"required"`
}
