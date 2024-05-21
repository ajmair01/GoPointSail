package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"pointsale.example/GoPointSail/pkg/model"
	"pointsale.example/GoPointSail/pkg/service/encryption"
	"pointsale.example/GoPointSail/pkg/service/persistence"
	"time"
)

func StartRouter() {
	router := gin.Default()

	cors.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type", "Origin"},
	}))

	router.GET("/users", getUsers)
	router.POST("/users", postUser)
	router.POST("/login", login)

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
		slog.Error("Error mapping request to user model for signup", err)
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		newUser.CreatedAt = time.Now()
		newUser.LastUpdated = time.Now()
		passwordHash, err := encryption.HashPassword(newUser.Password)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		newUser.Password = passwordHash
		err = persistence.AddUser(newUser)
		if err == nil {
			c.IndentedJSON(http.StatusCreated, newUser)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}
}

func login(c *gin.Context) {
	var loginUser model.User
	err := c.BindJSON(&loginUser)
	slog.Info("Login request for " + fmt.Sprintf("%#v", loginUser))
	passwordFromRequest := loginUser.Password
	if err != nil {
		slog.Error("Error mapping request to user model for login", err)
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		loginUser, err = persistence.GetUserByEmail(loginUser.Email)
		if err != nil {
			slog.Error("Error retrieving user by email", err)
			c.AbortWithError(http.StatusInternalServerError, err)
		} else {
			passwordMatches := encryption.CheckPasswordHash(passwordFromRequest, loginUser.Password)
			if passwordMatches {
				c.IndentedJSON(http.StatusOK, loginUser)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}
