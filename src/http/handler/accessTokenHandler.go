package handler

import (
	"bookStoreOauthApi/src/domain/accessToken"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
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
		fmt.Println("errrrrrrr")
		c.JSON(int(err.Status), err)
	}
	c.JSON(http.StatusAccepted, accessToken)
}
