package app

import (
	"github.com/ahmed3520/msp-server/domain"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	MapUrls()
	domain.ConnDB()
	router.Use(cors.Default())
	router.Run(":8080")
}
