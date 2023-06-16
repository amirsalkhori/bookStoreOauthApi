package app

import (
	"bookStoreOauthApi/src/domain/accessToken"
	"bookStoreOauthApi/src/http/handler"
	"bookStoreOauthApi/src/repository/db"
	"bookStoreOauthApi/src/repository/rest"
	"bookStoreOauthApi/src/service"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := service.NewService(dbRepository)
	atHandler := handler.NewHandler(atService)

	userDbRepository := rest.NewUserRepository()
	userService := accessToken.NewUserService(userDbRepository)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/oath/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oath/access_token", atHandler.Create)
	router.POST("/oath/access_token/:access_token_id", atHandler.UpdateExpirationTime)
	router.POST("/oath/login/", userHandler.LoginUser)
	router.Run(":8080")
}
