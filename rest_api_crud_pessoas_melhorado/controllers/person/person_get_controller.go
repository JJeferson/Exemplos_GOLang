package person

import (
	"context"
	"net/http"

	"rest_api_crud_pessoas_melhorado/dto"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonGetController struct {
	client *mongo.Client
}

func NewPersonGetController(client *mongo.Client) *PersonGetController {
	return &PersonGetController{client: client}
}

func (pgc *PersonGetController) GetAllPeople(c *gin.Context) {
	collection := pgc.client.Database("crud_1_go_lang").Collection("pessoas")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(context.Background())

	var responseDTOs []dto.PersonResponseDTO
	for cur.Next(context.Background()) {
		var person dto.PersonResponseDTO
		err := cur.Decode(&person)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		responseDTOs = append(responseDTOs, person)
	}

	if err := cur.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responseDTOs)
}

func (pgc *PersonGetController) GetPersonByID(c *gin.Context) {
	id := c.Param("id")

	collection := pgc.client.Database("crud_1_go_lang").Collection("pessoas")
	var person dto.PersonResponseDTO
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (pgc *PersonGetController) GetPeopleByName(c *gin.Context) {
	name := c.Param("nome")

	collection := pgc.client.Database("crud_1_go_lang").Collection("pessoas")
	cur, err := collection.Find(context.Background(), bson.M{"nome": name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(context.Background())

	var responseDTOs []dto.PersonResponseDTO
	for cur.Next(context.Background()) {
		var person dto.PersonResponseDTO
		err := cur.Decode(&person)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		responseDTOs = append(responseDTOs, person)
	}

	if err := cur.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responseDTOs)
}
