package routes

import (
	"github.com/RafaelRochaS/above-api/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(server *gin.Engine) {
	v1 := server.Group("api/v1")

	handlers.HandleHealthcheck(v1)
	handlers.HandleAccounts(v1)
}
