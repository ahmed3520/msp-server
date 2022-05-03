package app

import (
	"github.com/ahmed3520/msp-server/controllers/ping"
	"github.com/ahmed3520/msp-server/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/find", users.FindUser)
	router.GET("users/delete", users.DeleteUser)
	router.GET("/users/update", users.UpdateUser)
	router.POST("/users/create", users.CreateUser)
}
