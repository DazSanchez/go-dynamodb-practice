package server

import (
	"github.com/DazSanchez/go-dynamodb-practice/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	user := new(controllers.UserController)

	router.GET("/users", user.GetUsers)
	router.GET("/users/:id", user.GetUserById)
	router.POST("/users", user.CreateUser)

	return router
}
