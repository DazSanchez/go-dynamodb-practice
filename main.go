package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	FirstSurname string `json:"firstSurname"`
	Email        string `json:"email"`
}

type CreateUser struct {
	Name         string `json:"name"`
	FirstSurname string `json:"firstSurname"`
	Email        string `json:"email"`
}

var users = []User{
	{ID: uuid.NewString(), Name: "Hugo", FirstSurname: "Sanchez", Email: "hsanchez@yopmail.com"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserById)
	router.POST("/users", createUser)

	router.Run("localhost:8000")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser CreateUser

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	var user = User{
		ID:           uuid.NewString(),
		Name:         newUser.Name,
		FirstSurname: newUser.FirstSurname,
		Email:        newUser.Email,
	}

	users = append(users, user)

	c.JSON(http.StatusCreated, user)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")

	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
