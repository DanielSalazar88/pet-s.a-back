package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"pet-s.a/internal/entity"
)

func (h *MyHandlers) insertPet(c echo.Context) error {
	params := entity.NewPet{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.InsertPet(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Mascota Insertada Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})

}
