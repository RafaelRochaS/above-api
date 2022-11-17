package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAccounts(rg *gin.RouterGroup) {
	accounts := rg.Group("/accounts")

	accounts.POST("/", handleAccountCreation)
}

func handleAccountCreation(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "[WIP] account created")
}
