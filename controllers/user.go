package controllers

import (
	"net/http"

	"github.com/DazSanchez/go-dynamodb-practice/forms"
	"github.com/DazSanchez/go-dynamodb-practice/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) GetUsers(c *gin.Context) {
	users, err := userModel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving users",
			"error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}

	user, err := userModel.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving user", "error": err})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}

	c.JSON(http.StatusOK, user)
}

func (u UserController) CreateUser(c *gin.Context) {
	var newUser forms.CreateUserDTO

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	user, err := userModel.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving user", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
