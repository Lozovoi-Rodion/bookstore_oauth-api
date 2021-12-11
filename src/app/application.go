package app

import (
	"github.com/Lozovoi-Rodion/bookstore_oauth-api/src/http"
	"github.com/Lozovoi-Rodion/bookstore_oauth-api/src/repository/db"
	"github.com/Lozovoi-Rodion/bookstore_oauth-api/src/repository/rest"
	"github.com/Lozovoi-Rodion/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(rest.NewRepository(), db.NewRepository())
	atHandler := http.NewAccessTokenHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.PUT("/oauth/access_token/:access_token_id", atHandler.Update)
	router.Run(":8080")
}
