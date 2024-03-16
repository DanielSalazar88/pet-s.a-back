package handlers

import "github.com/labstack/echo/v4"

func (h *MyHandlers) registerRoutes(e *echo.Echo) {
	clients := e.Group("/clients")
	clients.POST("/insert-client/", h.insertClient)
	clients.POST("/delete-client/", h.deleteClient)

	pets := e.Group("/pets")
	pets.POST("/insert-pet/", h.insertPet)

	medicines := e.Group("/medicines")
	medicines.POST("/insert-medicine/", h.insertMedicine)
	medicines.POST("/insert-recipe/", h.insertRecipe)

}
