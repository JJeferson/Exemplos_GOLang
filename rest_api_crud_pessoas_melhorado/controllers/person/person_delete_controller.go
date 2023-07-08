package person

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonDeleteController struct {
	client *mongo.Client
}

func NewPersonDeleteController(client *mongo.Client) *PersonDeleteController {
	return &PersonDeleteController{client: client}
}

func (pdc *PersonDeleteController) DeletePerson(c *gin.Context) {
	id := c.Param("id")

	collection := pdc.client.Database("crud_1_go_lang").Collection("pessoas")
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
}
