package handler

import (
	"bookStoreOauthApi/src/domain/accessToken"
	"bookStoreOauthApi/src/domain/users"
	"bookStoreOauthApi/src/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	LoginUser(*gin.Context)
}

type userHandler struct {
	userService accessToken.UserService
}

func NewUserHandler(service accessToken.UserService) *userHandler {
	return &userHandler{
		userService: service,
	}
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var ul users.UserLoginRequest
	if err := c.ShouldBindJSON(&ul); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body!")
		c.JSON(int(restErr.Status), restErr.Error)
		return
	}
	result, err := h.userService.LoginUser(ul.Email, ul.Password)
	if err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
