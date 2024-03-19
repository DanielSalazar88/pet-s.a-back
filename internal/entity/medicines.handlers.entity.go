package entity

type NewMedicine struct {
	Nombre      string `json:"nombre" validate:"required"`
	Descripcion string `json:"descripcion" validate:"required"`
	Dosis       string `json:"dosis" validate:"required"`
}

type DeleteMedicine struct {
	ID int `json:"id_medicine" validate:"required"`
}

type UpdateMedicine struct {
	ID          int    `json:"id_medicamento" validate:"required"`
	Nombre      string `json:"nombre" validate:"required"`
	Descripcion string `json:"descripcion" validate:"required"`
	Dosis       string `json:"dosis" validate:"required"`
}
