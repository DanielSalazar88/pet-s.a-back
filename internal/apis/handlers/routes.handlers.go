package handlers

import "github.com/labstack/echo/v4"

func (h *MyHandlers) registerRoutes(e *echo.Echo) {
	clients := e.Group("/clients")
	clients.POST("/insert-client/", h.insertClient)
	clients.POST("/delete-client/", h.deleteClient)
	clients.POST("/update-client/", h.updateClient)
	clients.GET("/get-clients/", h.getClients)

	pets := e.Group("/pets")
	pets.POST("/insert-pet/", h.insertPet)
	pets.POST("/delete-pet/", h.deletePet)
	pets.POST("/update-pet/", h.updatePet)
	pets.GET("/get-pets/", h.getPets)

	medicines := e.Group("/medicines")
	medicines.POST("/insert-medicine/", h.insertMedicine)
	medicines.POST("/delete-medicine/", h.deleteMedicine)
	medicines.POST("/update-medicine/", h.updateMedicine)
	medicines.GET("/get-medicines/", h.getMedicines)

	recipes := e.Group("/recipes")
	recipes.POST("/insert-recipe/", h.insertRecipe)
	recipes.POST("/delete-recipe/", h.deleteRecipe)
	recipes.GET("/get-recipes/", h.getRecipes)
}
