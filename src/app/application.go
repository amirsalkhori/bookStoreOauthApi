package app

import (
	"bookStoreOauthApi/src/domain/accessToken"
	"bookStoreOauthApi/src/http/handler"
	"bookStoreOauthApi/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := accessToken.NewService(dbRepository)
	atHandler := handler.NewHandler(atService)

	router.GET("/oath/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
