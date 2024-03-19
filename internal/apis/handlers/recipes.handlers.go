package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"pet-s.a/internal/entity"
)

func (h *MyHandlers) insertRecipe(c echo.Context) error {
	params := entity.NewRecipe{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.InsertRecipe(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Receta Insertada Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})
}

func (h *MyHandlers) deleteRecipe(c echo.Context) error {
	params := entity.DeleteRecipe{}

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": fmt.Sprintf("Invalid Request, %s ", err)})
	}

	if err := h.DataValidator.Struct(params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	result, err := h.Service.DeleteRecipe(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Receta Eliminada Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})
}

func (h *MyHandlers) getRecipes(c echo.Context) error {

	result, err := h.Service.GetRecipes()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Message": err.Error()})
	}

	log.Println("#########################################################")
	log.Println("Recetas Obtenidas Correctamente!")
	log.Println("#########################################################")

	return c.JSON(http.StatusOK, map[string]string{"Message": result})
}
