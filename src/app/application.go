package app

import (
	"bookStoreOauthApi/src/client/cassandra"
	"bookStoreOauthApi/src/domain/accessToken"
	"bookStoreOauthApi/src/http/handler"
	"bookStoreOauthApi/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}

	defer session.Close()
	dbRepository := db.NewRepository()
	atService := accessToken.NewService(dbRepository)
	atHandler := handler.NewHandler(atService)

	router.GET("/oath/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oath/access_token", atHandler.Create)
	router.POST("/oath/access_token/:access_token_id", atHandler.UpdateExpirationTime)
	router.Run(":8080")
}
