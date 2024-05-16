package router

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"pointsale.example/GoPointSail/pkg/model"
	"pointsale.example/GoPointSail/pkg/service/persistence"
	"time"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/users", getUsers)
	router.POST("/users", postUser)

	err := router.Run("localhost:8080")
	if err != nil {
		slog.Error("Error running gin router", err)
	}
}

func getUsers(c *gin.Context) {
	users, err := persistence.AllUsers()
	if err == nil {
		c.IndentedJSON(http.StatusOK, users)
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func postUser(c *gin.Context) {
	var newUser model.User
	err := c.BindJSON(&newUser)
	if err != nil {
		slog.Error("Error mapping request to user model", err)
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		newUser.CreatedAt = time.Now()
		newUser.LastUpdated = time.Now()
		err = persistence.AddUser(newUser)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, newUser)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}
}
