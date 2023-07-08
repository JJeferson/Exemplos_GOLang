package person

import (
	"context"
	"net/http"

	"rest_api_crud_pessoas_melhorado/dto"
	"rest_api_crud_pessoas_melhorado/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonUpdateController struct {
	client *mongo.Client
}

func NewPersonUpdateController(client *mongo.Client) *PersonUpdateController {
	return &PersonUpdateController{client: client}
}

func (puc *PersonUpdateController) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var updatedPerson models.Person

	if err := c.ShouldBindJSON(&updatedPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := puc.client.Database("crud_1_go_lang").Collection("pessoas")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedPerson}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseDTO := dto.PersonResponseDTO{
		ID:   updatedPerson.ID,
		Nome: updatedPerson.Nome,
	}

	c.JSON(http.StatusOK, responseDTO)
}
