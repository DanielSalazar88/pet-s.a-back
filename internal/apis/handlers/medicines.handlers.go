package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"pet-s.a/internal/entity"
)

func (h *MyHandlers) insertMedicine(c echo.Context) error {
	params := entity.NewMedicine{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.InsertMedicine(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Medicina Insertada Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})

}

func (h *MyHandlers) deleteMedicine(c echo.Context) error {
	params := entity.DeleteMedicine{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.DeleteMedicine(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Medicina Eliminada Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})

}

func (h *MyHandlers) updateMedicine(c echo.Context) error {
	params := entity.UpdateMedicine{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.UpdateMedicine(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Medicina Actualizada Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})

}

func (h *MyHandlers) getMedicines(c echo.Context) error {

	result, err := h.Service.GetMedicines()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Medicamentos Obtenidos Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})
}
