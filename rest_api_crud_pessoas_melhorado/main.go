package main

import (
	"rest_api_crud_pessoas_melhorado/controllers/person"
	"rest_api_crud_pessoas_melhorado/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Inicializar o cliente do MongoDB
	client := database.ConnectDB()

	// Inicializar os controllers
	personCreateController := person.NewPersonCreateController(client)
	personGetController := person.NewPersonGetController(client)
	personUpdateController := person.NewPersonUpdateController(client)
	personDeleteController := person.NewPersonDeleteController(client)

	// Definir as rotas e os handlers dos controllers
	router.POST("/pessoas", personCreateController.CreatePerson)
	router.GET("/pessoas", personGetController.GetAllPeople)
	router.GET("/pessoas/:id", personGetController.GetPersonByID)
	router.GET("/pessoas/nome/:nome", personGetController.GetPeopleByName)
	router.PUT("/pessoas/:id", personUpdateController.UpdatePerson)
	router.DELETE("/pessoas/:id", personDeleteController.DeletePerson)

	router.Run(":3031")
}
