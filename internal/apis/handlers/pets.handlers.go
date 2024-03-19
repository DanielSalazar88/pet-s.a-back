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

func (h *MyHandlers) deletePet(c echo.Context) error {
	params := entity.DeletePet{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.DeletePet(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Mascota Eliminada Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})

}

func (h *MyHandlers) updatePet(c echo.Context) error {
	params := entity.UpdatePet{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.UpdatePet(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Mascota Actualizada Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})

}

func (h *MyHandlers) getPets(c echo.Context) error {

	result, err := h.Service.GetPets()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Mascotas Obtenidas Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})
}
