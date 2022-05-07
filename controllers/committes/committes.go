package committes

import (
	"fmt"
	"net/http"

	"github.com/ahmed3520/msp-server/domain"
	"github.com/ahmed3520/msp-server/services"
	"github.com/ahmed3520/msp-server/utils"
	"github.com/gin-gonic/gin"
)

func CreateCommitte(c *gin.Context) {
	var newCommitte domain.Committe

	fmt.Println(newCommitte)
	if err := c.ShouldBindJSON(&newCommitte); err != nil {
		restErr := utils.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.CreateCommite(&newCommitte)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func FindCommitte(c *gin.Context) {
	eventName := c.Query("name")

	if eventName == "" {
		restErr := utils.BadRequest("no name found..")
		c.JSON(restErr.Status, restErr)
		return
	}
	event, restErr := services.FindCommitte(eventName)

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, event)
}
func GetAllCommitte(c *gin.Context) {
	committe, restErr := services.GetAllCommitte()
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, committe)
}
func UpdateCommitte(c *gin.Context) {
	committeName := c.Query("name")
	field := c.Query("field")
	value := c.Query("value")
	if committeName == "" {
		restErr := utils.BadRequest("no name..")
		c.JSON(restErr.Status, restErr)
		return
	}
	if field == "" {
		restErr := utils.BadRequest("no field..")
		c.JSON(restErr.Status, restErr)
		return
	}
	if value == "" {
		restErr := utils.BadRequest("no value..")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.UpdateCommitte(committeName, field, value)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
