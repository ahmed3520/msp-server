package events

import (
	"fmt"
	"net/http"

	"github.com/ahmed3520/msp-server/domain"
	"github.com/ahmed3520/msp-server/services"
	"github.com/ahmed3520/msp-server/utils"
	"github.com/gin-gonic/gin"
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
