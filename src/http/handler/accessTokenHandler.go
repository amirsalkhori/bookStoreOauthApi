package handler

import (
	"bookStoreOauthApi/src/domain/accessToken"
	"bookStoreOauthApi/src/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service accessToken.Service
}

func NewHandler(service accessToken.Service) *accessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenParam := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := h.service.GetById(accessTokenParam)
	if err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusAccepted, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at accessToken.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body!")
		c.JSON(int(restErr.Status), restErr.Error)
		return
	}

	if err := h.service.Create(at); err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusCreated, at)
}

func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	// accessTokenParam := strings.TrimSpace(c.Param("access_token_id"))
	var at accessToken.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body!")
		c.JSON(int(restErr.Status), restErr.Error)
		return
	}

	if err := h.service.UpdateExpirationTime(at); err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusCreated, at)
}
