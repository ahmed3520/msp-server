package events

import (
	"fmt"
	"net/http"

	"github.com/ahmed3520/msp-server/domain"
	"github.com/ahmed3520/msp-server/services"
	"github.com/ahmed3520/msp-server/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEvent(c *gin.Context) {
	var newEvent domain.Event

	fmt.Println(newEvent)
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		restErr := utils.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.CreateEvent(&newEvent)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func FindEvent(c *gin.Context) {
	eventId := c.Query("id")
	objID, err := primitive.ObjectIDFromHex(eventId)
	if err != nil {
		panic(err)
	}

	if eventId == "" {
		restErr := utils.BadRequest("no id found..")
		c.JSON(restErr.Status, restErr)
		return
	}
	event, restErr := services.FindEvent(objID)

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, event)
}

func DeleteEvent(c *gin.Context) {
	eventId := c.Query("id")
	objID, err := primitive.ObjectIDFromHex(eventId)
	if err != nil {
		panic(err)
	}
	if eventId == "" {
		restErr := utils.BadRequest("no id found..")
		c.JSON(restErr.Status, restErr)
		return
	}
	restErr := services.DeleteEvent(objID)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"isRemoved": true})
}
