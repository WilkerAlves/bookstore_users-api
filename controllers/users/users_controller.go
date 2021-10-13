package users

import (
	"github.com/WilkerAlves/bookstore_users-api/domain/users"
	"github.com/WilkerAlves/bookstore_users-api/services"
	"github.com/WilkerAlves/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadResquestError("invalid json body")
		c.JSON(restErr.Status, restErr.Message)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr.Message)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
