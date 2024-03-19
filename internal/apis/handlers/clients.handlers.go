package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"pet-s.a/internal/entity"
)

func (h *MyHandlers) insertClient(c echo.Context) error {
	params := entity.NewClient{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.InsertClient(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Cliente Insertado Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})
}

func (h *MyHandlers) deleteClient(c echo.Context) error {
	params := entity.DeleteClient{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.DeleteClient(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Cliente Eliminado Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})
}

func (h *MyHandlers) updateClient(c echo.Context) error {
	params := entity.NewClient{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.UpdateClient(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Cliente Actualizado Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})
}

func (h *MyHandlers) getClients(c echo.Context) error {

	result, err := h.Service.GetClients()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Cliente Obtenidos Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})
}
