// go get -u github.com/gin-gonic/gin
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var userList = []User{
	{ID: "1", Username: "user1", Email: "user1@example.com"},
	{ID: "2", Username: "user2", Email: "user2@example.com"},
	{ID: "3", Username: "user3", Email: "user3@example.com"},
}

func main() {
	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run(":8080")
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	for _, user := range userList {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func createUser(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userList = append(userList, newUser)

	c.JSON(http.StatusCreated, newUser)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range userList {
		if user.ID == id {
			var updatedUser User

			if err := c.ShouldBindJSON(&updatedUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			userList[i] = updatedUser

			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range userList {
		if user.ID == id {
			userList = append(userList[:i], userList[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
