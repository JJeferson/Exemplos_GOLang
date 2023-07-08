package person

import (
	"context"
	"net/http"

	"rest_api_crud_pessoas_melhorado/dto"
	"rest_api_crud_pessoas_melhorado/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonCreateController struct {
	client *mongo.Client
}

func NewPersonCreateController(client *mongo.Client) *PersonCreateController {
	return &PersonCreateController{client: client}
}

func (pcc *PersonCreateController) CreatePerson(c *gin.Context) {
	var newPerson models.Person

	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := pcc.client.Database("crud_1_go_lang").Collection("pessoas")
	_, err := collection.InsertOne(context.Background(), newPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseDTO := dto.PersonResponseDTO{
		ID:   newPerson.ID,
		Nome: newPerson.Nome,
	}

	c.JSON(http.StatusOK, responseDTO)
}
