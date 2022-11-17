package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHealthcheck(rg *gin.RouterGroup) {
	healthcheck := rg.Group("/health")

	healthcheck.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
}
